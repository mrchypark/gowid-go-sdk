package client

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func newV2Server(t *testing.T, payload string, captured **http.Request) *Client {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if captured != nil {
			clone := r.Clone(r.Context())
			if r.Body != nil {
				body, _ := io.ReadAll(r.Body)
				clone.Body = io.NopCloser(strings.NewReader(string(body)))
			}
			*captured = clone
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(payload))
	}))
	t.Cleanup(server.Close)
	c := NewClient("secret")
	c.BaseURL = server.URL
	return c
}

func TestGetCardsV2DecodesHolders(t *testing.T) {
	c := newV2Server(t, `{
  "result": {"code": 20000000, "desc": "success"},
  "totalCount": 2,
  "data": {
    "totalPages": 1,
    "totalElements": 2,
    "last": true,
    "content": [
      {
        "cardId": 1001,
        "companyCode": "0305",
        "cardStatus": "NORMAL",
        "cardAlias": "법인 운영 카드",
        "managementNumber": "MGMT-001",
        "maskedCardNumber": "552576******1234",
        "encryptedCardNumber": "ENC-abc",
        "limitAmount": 10000000,
        "usedAmount": 2500000,
        "remainAmount": 7500000,
        "cardUser": {"userId": 77, "name": "홍길동", "departmentName": "개발팀"},
        "userHolders": [
          {"userId": 88, "name": "김철수", "departmentName": "기획팀"}
        ],
        "departmentHolders": [
          {"departmentId": 9, "name": "영업1팀"}
        ]
      },
      {
        "cardId": 1002,
        "companyCode": "0311",
        "cardStatus": "ISSUED",
        "cardAlias": "",
        "managementNumber": "MGMT-002",
        "maskedCardNumber": "461980******9876",
        "encryptedCardNumber": "ENC-def",
        "limitAmount": 0,
        "usedAmount": 0,
        "remainAmount": 0,
        "cardUser": {"userId": 91, "name": "박영희", "departmentName": ""},
        "userHolders": [],
        "departmentHolders": []
      }
    ]
  }
}`, nil)

	res, err := c.GetCardsV2(&PageOptionsV2{Page: 0, Size: 20, Sort: "limitAmount,desc"})
	if err != nil {
		t.Fatal(err)
	}
	if res.TotalCount != 2 || res.Result.Code != 20000000 {
		t.Fatalf("envelope = %+v", res)
	}
	page := res.Data
	if !page.Last || page.TotalPages != 1 || page.TotalElements != 2 || len(page.Content) != 2 {
		t.Fatalf("page = %+v", page)
	}
	card := page.Content[0]
	if card.CardID != 1001 || card.MaskedCardNumber != "552576******1234" || card.CardStatus != "NORMAL" {
		t.Fatalf("card = %+v", card)
	}
	if card.LimitAmount != 10000000 || card.UsedAmount != 2500000 || card.RemainAmount != 7500000 {
		t.Fatalf("card amounts = %+v", card)
	}
	if card.CardUser.UserID != 77 || card.CardUser.Name != "홍길동" || card.CardUser.DepartmentName != "개발팀" {
		t.Fatalf("cardUser = %+v", card.CardUser)
	}
	if len(card.UserHolders) != 1 || card.UserHolders[0].UserID != 88 || card.UserHolders[0].DepartmentName != "기획팀" {
		t.Fatalf("userHolders = %+v", card.UserHolders)
	}
	if len(card.DepartmentHolders) != 1 || card.DepartmentHolders[0].DepartmentID != 9 || card.DepartmentHolders[0].Name != "영업1팀" {
		t.Fatalf("departmentHolders = %+v", card.DepartmentHolders)
	}
	second := page.Content[1]
	if second.CardID != 1002 || second.CardUser.UserID != 91 || len(second.UserHolders) != 0 || len(second.DepartmentHolders) != 0 {
		t.Fatalf("second card = %+v", second)
	}
}

func TestGetCardsV2Query(t *testing.T) {
	var got *http.Request
	c := newV2Server(t, `{"result":{"code":20000000,"desc":"success"},"data":{"content":[]}}`, &got)
	if _, err := c.GetCardsV2(&PageOptionsV2{Page: 0, Size: 100, Sort: "limitAmount,desc"}); err != nil {
		t.Fatal(err)
	}
	q := got.URL.Query()
	if q.Get("size") != "100" || q.Get("sort") != "limitAmount,desc" {
		t.Fatalf("query = %v", q)
	}
	if _, err := c.GetCardsV2(&PageOptionsV2{Size: 101}); err == nil {
		t.Fatal("expected error for size 101")
	}
}

func TestGetExpenseV2DecodesNullRegistrationNumber(t *testing.T) {
	c := newV2Server(t, `{
  "result": {"code": 20000000, "desc": "success"},
  "data": {
    "expenseId": 9001,
    "cardApprovalNumber": "1234567890",
    "expenseType": "PURCHASE",
    "expenseDate": "20260108",
    "expenseTime": "103024",
    "card": {
      "cardNumber": "5525761234561234",
      "encryptedCardNumber": "ENC-abc",
      "cardUser": {
        "userName": "홍길동",
        "email": "hong@example.com",
        "mobileNumber": "01011112222",
        "isInvitedUser": true,
        "isContractor": false,
        "isActivated": true,
        "position": "PARTNER",
        "activatedAt": "20240101120000",
        "deactivatedAt": null,
        "notificationOnOff": true,
        "status": "NORMAL"
      },
      "alias": "법인 운영 카드",
      "limitAmount": 10000000,
      "usedAmount": 2500000,
      "remainAmount": 7500000,
      "companyCode": "0305",
      "namedYn": "Y",
      "cardName": "BC CARD",
      "cardType": "0305",
      "userNm": "HONG GIL DONG",
      "duplicationStatus": "NORMAL",
      "invalid": false
    },
    "user": {"userName": "홍길동", "status": "NORMAL", "isActivated": true},
    "useAmount": 35.5,
    "currency": "USD",
    "krwAmount": 50732,
    "approvalStatus": "PARTIAL_APPROVED",
    "approvedAmount": 50000,
    "approvedAt": "20260108150121",
    "approvedBy": "관리자",
    "comments": [
      {
        "author": {"userName": "김철수", "status": "NORMAL"},
        "content": "확인했습니다",
        "createdAt": "2026-01-09T10:00:00"
      }
    ],
    "purpose": {
      "name": "저녁식사",
      "category": {"categoryId": 3, "name": "접대비"},
      "limitAmount": 100000,
      "listOrder": 12,
      "isActivated": true,
      "hasRequirement": true,
      "limitType": "PERSON"
    },
    "participants": [
      {"userName": "홍길동", "email": "hong@example.com", "status": "NORMAL"},
      {"userName": "김철수", "email": "kim@example.com", "status": "NORMAL"}
    ],
    "expenseExternalUsers": [{"name": "John Doe", "company": "Acme"}],
    "storeName": "UNITED",
    "storeAddress": "Chicago",
    "storeRegistrationNumber": null,
    "memo": "팀 회식",
    "evidenceList": [
      {
        "evidenceId": 1001,
        "fileName": "receipt_20260108.jpg",
        "mimeType": "image/jpeg",
        "signedUrl": "https://storage.gowid.com/x"
      }
    ],
    "companyCode": "0305",
    "commentCount": 1,
    "expenseDeductionResDto": {"isExpenseDeductible": true, "isDeducted": false},
    "purposeRequirementAnswers": [
      {"purposeRequirementId": 44, "purposeRequirementName": "식당명", "answers": ["홍콩반점"]}
    ],
    "isDomestic": false
  }
}`, nil)

	res, err := c.GetExpenseV2(9001)
	if err != nil {
		t.Fatal(err)
	}
	d := res.Data
	if d.ExpenseID != 9001 || d.KRWAmount != 50732 || d.UseAmount != 35.5 || d.Currency != "USD" {
		t.Fatalf("detail = %+v", d)
	}
	if d.StoreRegistrationNumber != nil {
		t.Fatalf("storeRegistrationNumber = %v, want nil for null", *d.StoreRegistrationNumber)
	}
	if d.Card.LimitAmount != 10000000 || d.Card.NamedYN != "Y" || d.Card.UserName != "HONG GIL DONG" {
		t.Fatalf("card = %+v", d.Card)
	}
	if d.Card.CardUser.DeactivatedAt != nil || !d.Card.CardUser.IsInvitedUser {
		t.Fatalf("cardUser = %+v", d.Card.CardUser)
	}
	if len(d.Comments) != 1 || d.Comments[0].Author.UserName != "김철수" {
		t.Fatalf("comments = %+v", d.Comments)
	}
	if len(d.Participants) != 2 || d.Participants[1].Email != "kim@example.com" {
		t.Fatalf("participants = %+v", d.Participants)
	}
	if len(d.ExpenseExternalUsers) != 1 || d.ExpenseExternalUsers[0].Company != "Acme" {
		t.Fatalf("externalUsers = %+v", d.ExpenseExternalUsers)
	}
	if d.Purpose == nil || d.Purpose.Category.CategoryID != 3 || !d.Purpose.HasRequirement {
		t.Fatalf("purpose = %+v", d.Purpose)
	}
	if len(d.EvidenceList) != 1 || d.EvidenceList[0].EvidenceID != 1001 || d.EvidenceList[0].MIMEType != "image/jpeg" {
		t.Fatalf("evidenceList = %+v", d.EvidenceList)
	}
	if !d.ExpenseDeduction.ExpenseDeductible || d.ExpenseDeduction.Deducted {
		t.Fatalf("deduction = %+v", d.ExpenseDeduction)
	}
	if len(d.PurposeRequirementAnswers) != 1 || d.PurposeRequirementAnswers[0].Answers[0] != "홍콩반점" {
		t.Fatalf("answers = %+v", d.PurposeRequirementAnswers)
	}
	if d.CommentCount == nil || *d.CommentCount != 1 || d.Memo == nil || *d.Memo != "팀 회식" || d.IsDomestic {
		t.Fatalf("detail tail = %+v", d)
	}
}

func TestGetExpenseV2DecodesPresentRegistrationNumber(t *testing.T) {
	c := newV2Server(t, `{
  "result": {"code": 20000000, "desc": "success"},
  "data": {
    "expenseId": 9002,
    "storeName": "GS25",
    "storeRegistrationNumber": "1234567890",
    "isDomestic": true
  }
}`, nil)
	res, err := c.GetExpenseV2(9002)
	if err != nil {
		t.Fatal(err)
	}
	if res.Data.StoreRegistrationNumber == nil || *res.Data.StoreRegistrationNumber != "1234567890" {
		t.Fatalf("storeRegistrationNumber = %+v", res.Data.StoreRegistrationNumber)
	}
	if !res.Data.IsDomestic {
		t.Fatal("isDomestic = false")
	}
}

func TestUpdateExpenseV2OmitsNilAndSendsExplicitClears(t *testing.T) {
	var got *http.Request
	c := newV2Server(t, `{"result":{"code":20000000,"desc":"success"},"data":{"expenseId":7}}`, &got)

	if _, err := c.UpdateExpenseV2(7, UpdateExpenseRequestV2{}); err != nil {
		t.Fatal(err)
	}
	body := decodeBodyV2(t, got)
	if body["expenseId"] != float64(7) {
		t.Fatalf("expenseId = %v, want path id 7", body["expenseId"])
	}
	if body["isIgnoredFiles"] != true {
		t.Fatalf("isIgnoredFiles = %v, want true", body["isIgnoredFiles"])
	}
	for _, key := range []string{"purposeId", "memo", "participantIdList", "externalUserList", "purposeRequirementAnswerMap", "fileIdList"} {
		if _, ok := body[key]; ok {
			t.Errorf("key %q must be omitted when nil, body=%v", key, body)
		}
	}

	zeroPurpose := int64(0)
	emptyMemo := ""
	if _, err := c.UpdateExpenseV2(7, UpdateExpenseRequestV2{
		PurposeID:                 &zeroPurpose,
		Memo:                      &emptyMemo,
		ParticipantIDs:            []int64{},
		ExternalUsers:             []ExpenseExternalUserV2{},
		PurposeRequirementAnswers: map[string][]string{},
		FileIDs:                   []int64{},
	}); err != nil {
		t.Fatal(err)
	}
	body = decodeBodyV2(t, got)
	if body["purposeId"] != float64(0) {
		t.Errorf("purposeId = %v, want explicit 0", body["purposeId"])
	}
	if body["memo"] != "" {
		t.Errorf("memo = %v, want explicit empty string", body["memo"])
	}
	for _, key := range []string{"participantIdList", "externalUserList", "purposeRequirementAnswerMap", "fileIdList"} {
		value, ok := body[key]
		if !ok {
			t.Errorf("key %q must be sent as explicit empty value", key)
			continue
		}
		switch typed := value.(type) {
		case []any:
			if len(typed) != 0 {
				t.Errorf("%s = %v, want empty array", key, typed)
			}
		case map[string]any:
			if len(typed) != 0 {
				t.Errorf("%s = %v, want empty object", key, typed)
			}
		default:
			t.Errorf("%s = %#v, want empty array or object", key, value)
		}
	}
}

func TestUpdateExpenseV2NormalizesNilAnswers(t *testing.T) {
	var got *http.Request
	c := newV2Server(t, `{"result":{"code":20000000,"desc":"success"},"data":{"expenseId":7}}`, &got)
	caller := map[string][]string{"44": nil, "45": {"A"}}
	if _, err := c.UpdateExpenseV2(7, UpdateExpenseRequestV2{PurposeRequirementAnswers: caller}); err != nil {
		t.Fatal(err)
	}
	answers, ok := decodeBodyV2(t, got)["purposeRequirementAnswerMap"].(map[string]any)
	if !ok {
		t.Fatal("purposeRequirementAnswerMap missing")
	}
	value, ok := answers["44"].([]any)
	if !ok || len(value) != 0 {
		t.Fatalf("answers[44] = %#v, want empty array (never null)", answers["44"])
	}
	if caller["44"] != nil {
		t.Fatalf("caller map was mutated: answers[44] = %#v", caller["44"])
	}
}

func TestUpdateExpensePurposeV2OmitsNilAnswers(t *testing.T) {
	var got *http.Request
	c := newV2Server(t, `{"result":{"code":20000000,"desc":"success"},"data":{"expenseId":7}}`, &got)
	if _, err := c.UpdateExpensePurposeV2(7, UpdatePurposeRequestV2{PurposeID: 2}); err != nil {
		t.Fatal(err)
	}
	body := decodeBodyV2(t, got)
	if body["purposeId"] != float64(2) {
		t.Fatalf("purposeId = %v", body["purposeId"])
	}
	if _, ok := body["purposeRequirementAnswerMap"]; ok {
		t.Errorf("nil answer map must be omitted, body=%v", body)
	}
	if _, err := c.UpdateExpensePurposeV2(7, UpdatePurposeRequestV2{PurposeID: 2, PurposeRequirementAnswers: map[string][]string{}}); err != nil {
		t.Fatal(err)
	}
	body = decodeBodyV2(t, got)
	answers, ok := body["purposeRequirementAnswerMap"].(map[string]any)
	if !ok || len(answers) != 0 {
		t.Fatalf("empty answer map = %#v, want explicit empty object", body["purposeRequirementAnswerMap"])
	}
}

func TestUpdateExpensesPurposeV2Constraints(t *testing.T) {
	var got *http.Request
	c := newV2Server(t, `{"result":{"code":20000000,"desc":"success"},"data":true}`, &got)
	if _, err := c.UpdateExpensesPurposeV2(UpdatePurposesRequestV2{PurposeID: 2}); err == nil {
		t.Fatal("expected error for empty expenseIds")
	}
	if _, err := c.UpdateExpensesPurposeV2(UpdatePurposesRequestV2{ExpenseIDs: []int64{1, 2, 1}, PurposeID: 2}); err == nil {
		t.Fatal("expected error for duplicate expenseIds")
	}
	if _, err := c.UpdateExpensesPurposeV2(UpdatePurposesRequestV2{ExpenseIDs: []int64{1, 2}, PurposeID: 2}); err != nil {
		t.Fatal(err)
	}
	body := decodeBodyV2(t, got)
	ids, ok := body["expenseIds"].([]any)
	if !ok || len(ids) != 2 {
		t.Fatalf("expenseIds = %#v", body["expenseIds"])
	}
}

func TestUpdateExpenseParticipantsV2NormalizesNilArrays(t *testing.T) {
	var got *http.Request
	c := newV2Server(t, `{"result":{"code":20000000,"desc":"success"},"data":{"expenseId":7}}`, &got)
	if _, err := c.UpdateExpenseParticipantsV2(7, UpdateParticipantsRequestV2{}); err != nil {
		t.Fatal(err)
	}
	body := decodeBodyV2(t, got)
	for _, key := range []string{"participantIds", "externalUsers"} {
		value, ok := body[key].([]any)
		if !ok || len(value) != 0 {
			t.Errorf("%s = %#v, want explicit empty array (field is required)", key, body[key])
		}
	}
}

func TestApprovalStatusValidation(t *testing.T) {
	c := newV2Server(t, `{"result":{"code":20000000,"desc":"success"},"data":{"expenseId":7}}`, nil)
	for _, status := range []string{"", "SUBMITTED", "NOT_SUBMITTED", "PARTIAL_APPROVED"} {
		if _, err := c.UpdateExpenseApprovalStatusV2(7, ApproveExpenseRequestV2{ApprovalStatus: status}); err == nil {
			t.Errorf("UpdateExpenseApprovalStatusV2 accepted status %q", status)
		}
		if _, err := c.ApproveExpensesV2([]ApproveExpenseRequestV2{{ExpenseID: 7, ApprovalStatus: status}}); err == nil {
			t.Errorf("ApproveExpensesV2 accepted status %q", status)
		}
	}
	for _, status := range []string{"APPROVED", "REJECTED"} {
		if _, err := c.UpdateExpenseApprovalStatusV2(7, ApproveExpenseRequestV2{ApprovalStatus: status}); err != nil {
			t.Errorf("UpdateExpenseApprovalStatusV2 rejected %q: %v", status, err)
		}
	}
}

func TestApproveExpensesV2RequiresExpenseID(t *testing.T) {
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"result":{"code":20000000,"desc":"success"},"data":{"succeedStatements":[],"failedStatements":[]}}`))
	}))
	defer server.Close()
	c := NewClient("secret")
	c.BaseURL = server.URL

	// No path id exists on the bulk endpoint, so an item without expenseId
	// cannot identify its expense and must never reach the network.
	if _, err := c.ApproveExpensesV2([]ApproveExpenseRequestV2{{ApprovalStatus: "APPROVED"}}); err == nil {
		t.Fatal("expected error for missing expenseId")
	}
	if _, err := c.ApproveExpensesV2([]ApproveExpenseRequestV2{{ExpenseID: 1, ApprovalStatus: "APPROVED"}, {ApprovalStatus: "REJECTED"}}); err == nil {
		t.Fatal("expected error when any item lacks expenseId")
	}
	if _, err := c.ApproveExpensesV2([]ApproveExpenseRequestV2{{ExpenseID: -1, ApprovalStatus: "APPROVED"}}); err == nil {
		t.Fatal("expected error for negative expenseId")
	}
	if calls != 0 {
		t.Fatalf("%d requests sent, want 0", calls)
	}

	if _, err := c.ApproveExpensesV2([]ApproveExpenseRequestV2{{ExpenseID: 1, ApprovalStatus: "APPROVED"}}); err != nil {
		t.Fatal(err)
	}
	if calls != 1 {
		t.Fatalf("%d requests sent, want 1 for valid input", calls)
	}
}

func TestSearchExpensesV2RejectsEndDate(t *testing.T) {
	c := newV2Server(t, `{"result":{"code":20000000,"desc":"success"},"data":{"content":[]}}`, nil)
	if _, err := c.SearchExpensesV2(&ExpenseSearchOptionsV2{StartDate: "20260101", EndDate: "20260131"}); err == nil {
		t.Fatal("expected error for endDate on deprecated endpoint")
	}
	if _, err := c.SearchExpensesV2(&ExpenseSearchOptionsV2{StartDate: "20260101"}); err != nil {
		t.Fatal(err)
	}
}

func TestSearchOptionsSizeLimit(t *testing.T) {
	c := newV2Server(t, `{"result":{"code":20000000,"desc":"success"},"data":{"content":[]}}`, nil)
	if _, err := c.GetExpenseStatementsV2(&ExpenseSearchOptionsV2{Size: 101}); err == nil {
		t.Fatal("expected error for size 101")
	}
}

func TestSearchExpensesV2HasNoSizeCap(t *testing.T) {
	var got *http.Request
	c := newV2Server(t, `{"result":{"code":20000000,"desc":"success"},"data":{"content":[]}}`, &got)
	if _, err := c.SearchExpensesV2(&ExpenseSearchOptionsV2{StartDate: "20260101", Size: 150}); err != nil {
		t.Fatal(err)
	}
	if size := got.URL.Query().Get("size"); size != "150" {
		t.Fatalf("size = %q, want 150 (deprecated endpoint documents no maximum)", size)
	}
}

func decodeBodyV2(t *testing.T, r *http.Request) map[string]any {
	t.Helper()
	if r == nil || r.Body == nil {
		t.Fatal("no captured request")
	}
	var body map[string]any
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	return body
}
