package client

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// expenseDetailJSON is a literal official ExpenseDetailResDto payload written
// independently of the Go types, so a type change cannot make the assertions
// below pass by construction.
const expenseDetailJSON = "{" +
	"\"result\":{\"code\":20000000,\"desc\":\"success\"},\"totalCount\":1,\"data\":{" +
	"\"expenseId\":1234567890123," +
	"\"cardApprovalNumber\":\"998877\",\"expenseDate\":\"20260108\",\"expenseTime\":\"103024\"," +
	"\"card\":{\"cardNumber\":\"1234****5678\",\"encryptedCardNumber\":\"ENC-ABC-123\"," +
	"\"cardUser\":{\"userName\":\"HONG\",\"email\":\"hong@example.com\",\"mobileNumber\":\"01012345678\",\"status\":\"NORMAL\"}," +
	"\"alias\":\"corp\",\"limitAmount\":5000000,\"usedAmount\":123456,\"remainAmount\":4876544," +
	"\"companyCode\":\"0305\",\"namedYn\":\"Y\",\"cardName\":\"BC\",\"cardType\":\"0305\"," +
	"\"userNm\":\"HONG GILDONG\",\"duplicationStatus\":\"NORMAL\",\"invalid\":false}," +
	"\"user\":{\"userName\":\"HONG\",\"email\":\"hong@example.com\",\"status\":\"NORMAL\",\"isActivated\":true}," +
	"\"useAmount\":35.0,\"currency\":\"USD\",\"krwAmount\":50732," +
	"\"approvalStatus\":\"PARTIAL_APPROVED\",\"approvedAmount\":50000,\"approvedAt\":\"20260108150121\",\"approvedBy\":\"manager\"," +
	"\"comments\":[{\"author\":{\"userName\":\"HONG\"},\"content\":\"ok\",\"createdAt\":\"2026-01-08T15:00:00\"}]," +
	"\"purpose\":{\"name\":\"meal\",\"category\":{\"categoryId\":10,\"name\":\"food\"},\"limitAmount\":100000," +
	"\"listOrder\":1,\"isActivated\":true,\"hasRequirement\":false,\"limitType\":\"PERSON\"}," +
	"\"participants\":[{\"userName\":\"HONG\",\"email\":\"hong@example.com\",\"status\":\"NORMAL\"}]," +
	"\"expenseExternalUsers\":[{\"name\":\"ext\",\"company\":\"acme\"}]," +
	"\"storeName\":\"UNITED\",\"storeAddress\":\"seoul\",\"storeRegistrationNumber\":\"1234567890\"," +
	"\"memo\":\"dinner\",\"evidenceList\":[{\"evidenceId\":1001,\"fileName\":\"r.jpg\",\"mimeType\":\"image/jpeg\",\"signedUrl\":\"https://x\"}]," +
	"\"companyCode\":\"0305\",\"commentCount\":1," +
	"\"purposeRequirementItem\":\"visit\",\"purposeRequirementItemType\":\"SELECT\",\"purposeRequirementValue\":\"meeting\"," +
	"\"expenseDeductionResDto\":{\"isExpenseDeductible\":true,\"isDeducted\":false}," +
	"\"isDomestic\":true,\"isPurposeRequirementInputValue\":true}}"

type recordedRequest struct {
	method   string
	path     string
	rawQuery string
	body     string
}

func newRecordingClient(t *testing.T, responses map[string]string) (*Client, *[]recordedRequest) {
	t.Helper()
	recorded := &[]recordedRequest{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := ""
		if r.Body != nil {
			b, _ := io.ReadAll(r.Body)
			body = string(b)
		}
		*recorded = append(*recorded, recordedRequest{r.Method, r.URL.Path, r.URL.RawQuery, body})
		w.Header().Set("Content-Type", "application/json")
		if resp, ok := responses[r.URL.Path]; ok {
			_, _ = w.Write([]byte(resp))
			return
		}
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("{\"result\":{\"code\":40400000,\"desc\":\"not found\"}}"))
	}))
	t.Cleanup(server.Close)

	c := NewClient("secret")
	c.BaseURL = server.URL
	return c, recorded
}

func decodeBody(t *testing.T, body string) map[string]any {
	t.Helper()
	var out map[string]any
	if err := json.Unmarshal([]byte(body), &out); err != nil {
		t.Fatalf("request body is not JSON: %v (body=%s)", err, body)
	}
	return out
}

func ptr[T any](v T) *T { return &v }

// The five single-expense PUTs all return the full expense detail.
func TestV1SingleMutationsReturnExpenseDetail(t *testing.T) {
	c, recorded := newRecordingClient(t, map[string]string{
		"/v1/expenses/1234567890123":                 expenseDetailJSON,
		"/v1/expenses/1234567890123/purposes":        expenseDetailJSON,
		"/v1/expenses/1234567890123/memo":            expenseDetailJSON,
		"/v1/expenses/1234567890123/participants":    expenseDetailJSON,
		"/v1/expenses/1234567890123/approval-status": expenseDetailJSON,
	})

	amount := int64(50000)
	approvedAt := "20260108150121"
	memo := "dinner"
	calls := []struct {
		name string
		fn   func() (*Response[ExpenseDetail], error)
	}{
		{"full", func() (*Response[ExpenseDetail], error) {
			return c.UpdateExpense("1234567890123", UpdateExpenseRequest{PurposeId: ptr(int64(42)), Memo: &memo})
		}},
		{"purpose", func() (*Response[ExpenseDetail], error) {
			return c.UpdateExpensePurpose("1234567890123", UpdatePurposeRequest{PurposeId: 42})
		}},
		{"memo", func() (*Response[ExpenseDetail], error) {
			return c.UpdateExpenseMemo("1234567890123", "dinner")
		}},
		{"participants", func() (*Response[ExpenseDetail], error) {
			return c.UpdateExpenseParticipants("1234567890123", UpdateParticipantsRequest{
				ParticipantIds: []int64{100, 200},
				ExternalUsers:  []ExpenseExternalUser{{Name: "ext", Company: "acme"}},
			})
		}},
		{"approval", func() (*Response[ExpenseDetail], error) {
			return c.UpdateExpenseApprovalStatus("1234567890123", UpdateApprovalStatusRequest{
				ApprovalStatus: "APPROVED", ApprovedAmount: &amount, ApprovedAt: &approvedAt,
			})
		}},
	}

	for _, call := range calls {
		resp, err := call.fn()
		if err != nil {
			t.Fatalf("%s: %v", call.name, err)
		}
		if resp.TotalCount != 1 {
			t.Errorf("%s: totalCount = %d, want 1", call.name, resp.TotalCount)
		}
		if resp.Data.ExpenseId != 1234567890123 {
			t.Errorf("%s: expenseId = %d, want 1234567890123 (int64, not string)", call.name, resp.Data.ExpenseId)
		}
		if resp.Data.KrwAmount != 50732 {
			t.Errorf("%s: krwAmount = %d, want 50732 (int64)", call.name, resp.Data.KrwAmount)
		}
		if resp.Data.StoreRegistrationNumber == nil || *resp.Data.StoreRegistrationNumber != "1234567890" {
			t.Errorf("%s: nullable storeRegistrationNumber not decoded", call.name)
		}
		if !resp.Data.IsDomestic {
			t.Errorf("%s: isDomestic not decoded", call.name)
		}
		if resp.Data.Card.EncryptedCardNumber != "ENC-ABC-123" || resp.Data.Card.LimitAmount != 5000000 {
			t.Errorf("%s: card fields not decoded: %+v", call.name, resp.Data.Card)
		}
		if len(resp.Data.Comments) != 1 || resp.Data.Comments[0].Author.UserName != "HONG" {
			t.Errorf("%s: comments not decoded: %+v", call.name, resp.Data.Comments)
		}
		if resp.Data.Purpose == nil || resp.Data.Purpose.Category.CategoryId != 10 || resp.Data.Purpose.LimitAmount != 100000 {
			t.Errorf("%s: purpose detail not decoded: %+v", call.name, resp.Data.Purpose)
		}
		if len(resp.Data.Participants) != 1 || len(resp.Data.ExpenseExternalUsers) != 1 || len(resp.Data.EvidenceList) != 1 {
			t.Errorf("%s: list fields not decoded: %+v", call.name, resp.Data)
		}
		if !resp.Data.ExpenseDeductionResDto.IsExpenseDeductible {
			t.Errorf("%s: deduction not decoded", call.name)
		}
	}

	if len(*recorded) != len(calls) {
		t.Fatalf("recorded %d requests, want %d", len(*recorded), len(calls))
	}
	for i, got := range *recorded {
		if got.method != "PUT" {
			t.Errorf("call %d: method = %s, want PUT", i, got.method)
		}
		if !strings.HasPrefix(got.path, "/v1/expenses/1234567890123") {
			t.Errorf("call %d: path = %s, want numeric id path", i, got.path)
		}
	}

	full := decodeBody(t, (*recorded)[0].body)
	if full["expenseId"] != float64(1234567890123) {
		t.Errorf("full update body expenseId = %v, want 1234567890123", full["expenseId"])
	}
	if full["isIgnoredFiles"] != true {
		t.Errorf("full update body isIgnoredFiles = %v, want true", full["isIgnoredFiles"])
	}
	if purpose := decodeBody(t, (*recorded)[1].body); purpose["expenseId"] != float64(1234567890123) {
		t.Errorf("purpose update body expenseId = %v, want auto-filled 1234567890123", purpose["expenseId"])
	}
	if participants := decodeBody(t, (*recorded)[3].body); participants["participantIds"] == nil {
		t.Errorf("participants body must use the participantIds key, got %s", (*recorded)[3].body)
	}
	if approval := decodeBody(t, (*recorded)[4].body); approval["approvedAmount"] != float64(50000) || approval["approvedAt"] != approvedAt {
		t.Errorf("approval body = %s, want approvedAmount/approvedAt present", (*recorded)[4].body)
	}
}

func TestV1BulkApproveUsesExpenseSimple(t *testing.T) {
	bulkJSON := "{\"result\":{\"code\":20000000,\"desc\":\"success\"},\"totalCount\":2,\"data\":{" +
		"\"succeedStatements\":[{\"expenseId\":1,\"expenseDate\":\"20260108\",\"expenseTime\":\"103024\"," +
		"\"useAmount\":35.0,\"currency\":\"USD\",\"krwAmount\":50000,\"storeName\":\"UNITED\",\"approvalStatus\":\"APPROVED\"," +
		"\"cardAlias\":\"corp\",\"shortCardNumber\":\"5678\",\"recommendedPurposeList\":[" +
		"{\"purposeId\":7,\"name\":\"meal\",\"category\":{\"categoryId\":10,\"name\":\"food\"},\"listOrder\":1," +
		"\"limitType\":\"PERSON\",\"limitAmount\":100000,\"isActivated\":true,\"hasRequirement\":false,\"isDeducted\":true}]}]," +
		"\"failedStatements\":[]}}"
	c, recorded := newRecordingClient(t, map[string]string{"/v1/expenses/approval-status/approved": bulkJSON})

	amount := int64(50000)
	resp, err := c.ApproveExpenses(ApproveExpensesRequest{
		{ExpenseId: ptr(int64(1)), ApprovalStatus: "APPROVED", ApprovedAmount: &amount},
		{ExpenseId: ptr(int64(2)), ApprovalStatus: "REJECTED"},
	})
	if err != nil {
		t.Fatalf("ApproveExpenses: %v", err)
	}
	if len(resp.Data.SucceedStatements) != 1 {
		t.Fatalf("succeedStatements = %+v, want 1", resp.Data.SucceedStatements)
	}
	row := resp.Data.SucceedStatements[0]
	if row.ExpenseId != 1 || row.KrwAmount != 50000 || row.CardAlias == nil || *row.CardAlias != "corp" {
		t.Errorf("ExpenseSimple fields not decoded: %+v", row)
	}
	if len(row.RecommendedPurposeList) != 1 || row.RecommendedPurposeList[0].PurposeId != 7 {
		t.Errorf("recommendedPurposeList not decoded: %+v", row.RecommendedPurposeList)
	}
	if resp.Data.FailedStatements == nil {
		t.Errorf("failedStatements = nil, want an empty slice for an explicit []")
	}

	var sent []map[string]any
	if err := json.Unmarshal([]byte((*recorded)[0].body), &sent); err != nil {
		t.Fatalf("bulk body is not a JSON array: %v", err)
	}
	if len(sent) != 2 || sent[0]["approvalStatus"] != "APPROVED" {
		t.Errorf("bulk body = %s", (*recorded)[0].body)
	}
	if _, ok := sent[1]["approvedAmount"]; ok {
		t.Errorf("omitted approvedAmount must not be sent: %s", (*recorded)[0].body)
	}
}

func TestV1PurposeRequirementsContent(t *testing.T) {
	c, recorded := newRecordingClient(t, map[string]string{
		"/v1/purposes/9/requirements": "{\"result\":{\"code\":20000000},\"totalCount\":2,\"data\":{\"content\":[\"a\",\"b\"]}}",
	})
	resp, err := c.GetPurposeRequirements(9)
	if err != nil {
		t.Fatalf("GetPurposeRequirements: %v", err)
	}
	if len(resp.Data.Content) != 2 || resp.Data.Content[0] != "a" {
		t.Errorf("content = %v, want the data object's list", resp.Data.Content)
	}
	if (*recorded)[0].path != "/v1/purposes/9/requirements" {
		t.Errorf("path = %s", (*recorded)[0].path)
	}
}

func TestV1NotSubmittedExpenses(t *testing.T) {
	const payload = `{"result":{"code":20000000},"totalCount":1,"data":{"totalPages":3,"totalElements":41,"last":false,"content":[` +
		`{"expenseId":5,"expenseDate":"20260108","expenseTime":"103024","useAmount":12.5,"currency":"KRW",` +
		`"krwAmount":17000,"storeName":"GS25","approvalStatus":"NOT_SUBMITTED","cardAlias":null,"shortCardNumber":"1234",` +
		`"recommendedPurposeList":[{"purposeId":7,"name":"meal","category":{"categoryId":10,"name":"food"},` +
		`"limitAmount":100000,"isActivated":true,"hasRequirement":false}]}]}}`
	c, recorded := newRecordingClient(t, map[string]string{"/v1/expenses/not-submitted": payload})

	resp, err := c.GetNotSubmittedExpenses(&PageOptions{Page: 1, Size: 10, Sort: "expenseDate,desc"})
	if err != nil {
		t.Fatalf("GetNotSubmittedExpenses: %v", err)
	}
	if resp.TotalCount != 1 || resp.Data.TotalElements != 41 || resp.Data.Last {
		t.Errorf("page envelope not decoded: %+v", resp)
	}
	row := resp.Data.Content[0]
	if row.ExpenseId != 5 || row.CardAlias != nil || len(row.RecommendedPurposeList) != 1 {
		t.Errorf("not-submitted row not decoded: %+v", row)
	}
	got := (*recorded)[0]
	if got.path != "/v1/expenses/not-submitted" {
		t.Errorf("path = %s", got.path)
	}
	for _, want := range []string{"page=1", "size=10", "sort=expenseDate%2Cdesc"} {
		if !strings.Contains(got.rawQuery, want) {
			t.Errorf("query %q missing %s", got.rawQuery, want)
		}
	}
}

func TestV1FullUpdateOmissionVersusEmptyArrays(t *testing.T) {
	c, recorded := newRecordingClient(t, map[string]string{"/v1/expenses/77": expenseDetailJSON})

	if _, err := c.UpdateExpense("77", UpdateExpenseRequest{}); err != nil {
		t.Fatalf("UpdateExpense(omitted): %v", err)
	}
	omitted := decodeBody(t, (*recorded)[0].body)
	for _, key := range []string{"participantIdList", "externalUserList", "memo", "purposeId", "fileIdList", "purposeRequirementValue"} {
		if _, ok := omitted[key]; ok {
			t.Errorf("omitted field %q must not appear in body: %s", key, (*recorded)[0].body)
		}
	}
	if _, ok := omitted["externalUsers"]; ok {
		t.Errorf("full update must use externalUserList, got %s", (*recorded)[0].body)
	}

	if _, err := c.UpdateExpense("77", UpdateExpenseRequest{
		ParticipantIdList: []int64{},
		ExternalUserList:  []ExpenseExternalUser{},
		FileIdList:        []int64{},
		Memo:              ptr(""),
	}); err != nil {
		t.Fatalf("UpdateExpense(empty): %v", err)
	}
	empty := decodeBody(t, (*recorded)[1].body)
	for _, key := range []string{"participantIdList", "externalUserList", "fileIdList"} {
		arr, ok := empty[key].([]any)
		if !ok || len(arr) != 0 {
			t.Errorf("%q = %v, want an explicit empty array", key, empty[key])
		}
	}
	if empty["memo"] != "" {
		t.Errorf("explicit empty memo = %v, want an empty string", empty["memo"])
	}
}

func TestV1InvalidIDAndBodyMismatchDoNotHitServer(t *testing.T) {
	c, recorded := newRecordingClient(t, map[string]string{"/v1/expenses/1": expenseDetailJSON})

	for _, bad := range []string{"", "abc", "0", "-1", "1.5", " 7"} {
		if _, err := c.GetExpense(bad); err == nil {
			t.Errorf("GetExpense(%q): want error", bad)
		}
		if _, err := c.UpdateExpenseMemo(bad, "x"); err == nil {
			t.Errorf("UpdateExpenseMemo(%q): want error", bad)
		}
		if _, err := c.AddComment(bad, "x"); err == nil {
			t.Errorf("AddComment(%q): want error", bad)
		}
	}
	if _, err := c.UpdateExpensePurpose("abc", UpdatePurposeRequest{PurposeId: 1}); err == nil {
		t.Error("UpdateExpensePurpose(bad id): want error")
	}
	if _, err := c.UpdateExpenseParticipants("abc", UpdateParticipantsRequest{}); err == nil {
		t.Error("UpdateExpenseParticipants(bad id): want error")
	}
	if _, err := c.UpdateExpenseApprovalStatus("abc", UpdateApprovalStatusRequest{ApprovalStatus: "APPROVED"}); err == nil {
		t.Error("UpdateExpenseApprovalStatus(bad id): want error")
	}
	if _, err := c.UpdateExpenseApprovalStatus("1", UpdateApprovalStatusRequest{ApprovalStatus: "SUBMITTED"}); err == nil {
		t.Error("approval status must be APPROVED or REJECTED")
	}

	mismatch := []func() error{
		func() error {
			_, err := c.UpdateExpense("1", UpdateExpenseRequest{ExpenseId: ptr(int64(2))})
			return err
		},
		func() error {
			_, err := c.UpdateExpensePurpose("1", UpdatePurposeRequest{ExpenseId: ptr(int64(9)), PurposeId: 1})
			return err
		},
		func() error {
			_, err := c.UpdateExpenseApprovalStatus("1", UpdateApprovalStatusRequest{ExpenseId: ptr(int64(9)), ApprovalStatus: "APPROVED"})
			return err
		},
	}
	for i, fn := range mismatch {
		if err := fn(); err == nil || !strings.Contains(err.Error(), "mismatch") {
			t.Errorf("mismatch %d: want a mismatch error, got %v", i, err)
		}
	}

	if len(*recorded) != 0 {
		t.Errorf("invalid input must not reach the server, recorded %d requests: %+v", len(*recorded), *recorded)
	}
}

func TestV1PurposesQueryAndDecode(t *testing.T) {
	c, recorded := newRecordingClient(t, map[string]string{
		"/v1/purposes": "{\"result\":{\"code\":20000000},\"totalCount\":1,\"data\":[" +
			"{\"purposeId\":1,\"name\":\"meal\",\"category\":{\"categoryId\":10,\"name\":\"food\"},\"listOrder\":1," +
			"\"limitType\":\"PERSON\",\"limitAmount\":100000,\"isActivated\":true,\"hasRequirement\":false," +
			"\"requirement\":{\"id\":3,\"type\":\"TEXT\",\"item\":\"visit\",\"guideDesc\":\"enter name\",\"isAvailableInput\":true}," +
			"\"isDeducted\":true}]}",
	})
	active := true
	resp, err := c.GetPurposes(&GetPurposesOptions{IsActivated: &active})
	if err != nil {
		t.Fatalf("GetPurposes: %v", err)
	}
	if got := (*recorded)[0].rawQuery; got != "isActivated=true" {
		t.Errorf("query = %q, want only isActivated=true (no page/limit)", got)
	}
	p := resp.Data[0]
	if p.PurposeId != 1 || p.LimitAmount != 100000 || p.Requirement == nil || p.Requirement.Id != 3 || !p.IsDeducted {
		t.Errorf("purpose not decoded: %+v", p)
	}
}

func TestV1MembersDecode(t *testing.T) {
	c, _ := newRecordingClient(t, map[string]string{
		"/v1/members": "{\"result\":{\"code\":20000000},\"totalCount\":1,\"data\":[" +
			"{\"userId\":9007199254740993,\"userName\":\"HONG\",\"email\":\"hong@example.com\",\"isContractor\":true," +
			"\"status\":\"NORMAL\",\"department\":{\"id\":3,\"name\":\"finance\"},\"position\":\"manager\"," +
			"\"role\":{\"type\":\"ROLE_MASTER\",\"name\":\"master\",\"description\":\"all\"},\"notificationOnOff\":true}]}",
	})
	resp, err := c.GetMembers()
	if err != nil {
		t.Fatalf("GetMembers: %v", err)
	}
	m := resp.Data[0]
	if m.UserID != 9007199254740993 || m.Department.ID != 3 || m.Role.Type != "ROLE_MASTER" {
		t.Errorf("member not decoded: %+v", m)
	}
}

// Remaining v1 endpoints: single GET, expense list GET, bulk purpose PUT and
// comment POST.
func TestV1RemainingEndpoints(t *testing.T) {
	const (
		expenseListJSON = "{\"result\":{\"code\":20000000},\"totalCount\":2,\"data\":{\"totalPages\":1,\"totalElements\":2,\"last\":true,\"content\":[" +
			"{\"expenseId\":11,\"expenseDate\":\"20260108\",\"expenseTime\":\"103024\",\"useAmount\":35.0,\"currency\":\"USD\",\"krwAmount\":50732," +
			"\"approvalStatus\":\"SUBMITTED\",\"purpose\":{\"purposeId\":3,\"name\":\"meal\",\"limitType\":\"PERSON\",\"limitAmount\":100000," +
			"\"isActivated\":true,\"hasRequirement\":false},\"cardAlias\":\"corp\",\"cardUserName\":\"HONG\",\"shortCardNumber\":\"5678\"," +
			"\"storeName\":\"UNITED\",\"storeAddress\":\"seoul\",\"memo\":\"dinner\",\"commentCount\":2,\"evidenceCount\":1," +
			"\"participantCount\":2,\"representativeParticipant\":\"HONG\",\"participants\":[{\"userId\":11,\"userName\":\"HONG\"}]," +
			"\"expenseExternalUsers\":[{\"name\":\"ext\",\"company\":\"acme\"}],\"purposeRequirementItem\":\"visit\"," +
			"\"purposeRequirementItemType\":\"SELECT\",\"purposeRequirementValue\":\"meeting\",\"isPurposeRequirementInputValue\":true}]}}"

		commentJSON = "{\"result\":{\"code\":20000000},\"totalCount\":1,\"data\":{\"author\":\"HONG\",\"department\":\"finance\"," +
			"\"content\":\"receipt ok\",\"createdAt\":\"2026-01-08T15:00:00\"}}"

		bulkPurposeJSON = "{\"result\":{\"code\":20000000,\"desc\":\"success\"},\"totalCount\":2,\"data\":true}"
	)

	c, recorded := newRecordingClient(t, map[string]string{
		"/v1/expenses/11":          expenseDetailJSON,
		"/v1/expenses":             expenseListJSON,
		"/v1/expenses/purposes":    bulkPurposeJSON,
		"/v1/expenses/11/comments": commentJSON,
	})

	// GET /v1/expenses/{expenseId}
	single, err := c.GetExpense("11")
	if err != nil {
		t.Fatalf("GetExpense: %v", err)
	}
	if single.Data.ExpenseId != 1234567890123 {
		t.Errorf("GetExpense data = %+v", single.Data)
	}

	// GET /v1/expenses with the sort query
	list, err := c.GetExpenses(&GetExpensesOptions{
		ApprovalState: "SUBMITTED",
		Memo:          "dinner",
		PurposeName:   "meal",
		UserName:      "HONG",
		StartDate:     "20260101",
		Sort:          "expenseDate,desc",
		Size:          20,
	})
	if err != nil {
		t.Fatalf("GetExpenses: %v", err)
	}
	if list.Data.TotalElements != 2 || !list.Data.Last || len(list.Data.Content) != 1 {
		t.Errorf("expense list not decoded: %+v", list.Data)
	}
	row := list.Data.Content[0]
	if row.ExpenseId != 11 || row.KrwAmount != 50732 || row.CommentCount != 2 || row.ParticipantCount != 2 {
		t.Errorf("expense list row not decoded: %+v", row)
	}
	if row.Purpose == nil || row.Purpose.PurposeId != 3 || row.CardAlias == nil || *row.CardAlias != "corp" {
		t.Errorf("expense list row nested fields not decoded: %+v", row)
	}
	if len(row.Participants) != 1 || row.Participants[0].UserId != 11 || len(row.ExpenseExternalUsers) != 1 {
		t.Errorf("expense list row lists not decoded: %+v", row)
	}

	// PUT /v1/expenses/purposes
	bulk, err := c.UpdateExpensesPurpose(UpdateExpensesPurposeRequest{
		{ExpenseId: ptr(int64(11)), PurposeId: 3},
		{ExpenseId: ptr(int64(12)), PurposeId: 4},
	})
	if err != nil {
		t.Fatalf("UpdateExpensesPurpose: %v", err)
	}
	if !bulk.Data {
		t.Errorf("bulk purpose data = %v, want true", bulk.Data)
	}

	// POST /v1/expenses/{expenseId}/comments
	comment, err := c.AddComment("11", "receipt ok")
	if err != nil {
		t.Fatalf("AddComment: %v", err)
	}
	if comment.Data.Author != "HONG" || comment.Data.Department != "finance" || comment.Data.Content != "receipt ok" {
		t.Errorf("comment not decoded: %+v", comment.Data)
	}

	want := []struct {
		method string
		path   string
		query  string
	}{
		{"GET", "/v1/expenses/11", ""},
		{"GET", "/v1/expenses", "approvalState=SUBMITTED&memo=dinner&page=0&purposeName=meal&size=20&sort=expenseDate%2Cdesc&startDate=20260101&userName=HONG"},
		{"PUT", "/v1/expenses/purposes", ""},
		{"POST", "/v1/expenses/11/comments", ""},
	}
	if len(*recorded) != len(want) {
		t.Fatalf("recorded %d requests, want %d: %+v", len(*recorded), len(want), *recorded)
	}
	for i, w := range want {
		got := (*recorded)[i]
		if got.method != w.method || got.path != w.path || got.rawQuery != w.query {
			t.Errorf("request %d = %s %s?%s, want %s %s?%s", i, got.method, got.path, got.rawQuery, w.method, w.path, w.query)
		}
	}

	var bulkSent []map[string]any
	if err := json.Unmarshal([]byte((*recorded)[2].body), &bulkSent); err != nil {
		t.Fatalf("bulk purpose body is not a JSON array: %v", err)
	}
	if len(bulkSent) != 2 {
		t.Fatalf("bulk purpose body = %s, want 2 items", (*recorded)[2].body)
	}
	for i, item := range bulkSent {
		if item["expenseId"] == nil || item["purposeId"] == nil {
			t.Errorf("bulk purpose item %d must carry expenseId and purposeId: %v", i, item)
		}
	}

	commentBody := decodeBody(t, (*recorded)[3].body)
	if _, ok := commentBody["comment"]; !ok {
		t.Errorf("comment body missing comment: %s", (*recorded)[3].body)
	}
	if _, ok := commentBody["expenseId"]; !ok {
		t.Errorf("comment body missing expenseId: %s", (*recorded)[3].body)
	}

	if _, err := c.UpdateExpensesPurpose(UpdateExpensesPurposeRequest{{PurposeId: 3}}); err == nil {
		t.Error("bulk purpose item without expenseId must be rejected before the request")
	}
}

func TestV1ApproveExpensesRejectsMissingIDAndStatusWithoutNetwork(t *testing.T) {
	c, recorded := newRecordingClient(t, map[string]string{})

	bad := []ApproveExpensesRequest{
		{{ApprovalStatus: "APPROVED"}},
		{{ExpenseId: ptr(int64(0)), ApprovalStatus: "APPROVED"}},
		{{ExpenseId: ptr(int64(-1)), ApprovalStatus: "REJECTED"}},
		{{ExpenseId: ptr(int64(1)), ApprovalStatus: "SUBMITTED"}},
		{{ExpenseId: ptr(int64(1)), ApprovalStatus: "PARTIAL_APPROVED"}},
		{{ExpenseId: ptr(int64(1)), ApprovalStatus: ""}},
		{{ExpenseId: ptr(int64(1)), ApprovalStatus: "APPROVED"}, {ApprovalStatus: "REJECTED"}},
	}
	for i, req := range bad {
		if _, err := c.ApproveExpenses(req); err == nil {
			t.Errorf("case %d: want a validation error", i)
		}
	}

	if len(*recorded) != 0 {
		t.Errorf("invalid approval input must not reach the server, recorded %d: %+v", len(*recorded), *recorded)
	}
}

func TestV1PurposeRequirementInputValueBothNames(t *testing.T) {
	c, recorded := newRecordingClient(t, map[string]string{
		"/v1/expenses/5/purposes": expenseDetailJSON,
	})

	// Neither name set: neither key is emitted.
	if _, err := c.UpdateExpensePurpose("5", UpdatePurposeRequest{PurposeId: 3}); err != nil {
		t.Fatalf("UpdateExpensePurpose(default): %v", err)
	}
	body := decodeBody(t, (*recorded)[0].body)
	for _, key := range []string{"purposeRequirementInputValue", "isPurposeRequirementInputValue"} {
		if _, ok := body[key]; ok {
			t.Errorf("%q must be omitted when unset: %s", key, (*recorded)[0].body)
		}
	}

	// Explicit false on the bare name must survive as false, not be dropped.
	if _, err := c.UpdateExpensePurpose("5", UpdatePurposeRequest{
		PurposeId:                    3,
		PurposeRequirementInputValue: ptr(false),
	}); err != nil {
		t.Fatalf("UpdateExpensePurpose(bare false): %v", err)
	}
	body = decodeBody(t, (*recorded)[1].body)
	v, ok := body["purposeRequirementInputValue"]
	if !ok || v != false {
		t.Errorf("purposeRequirementInputValue = %v (present=%v), want explicit false", v, ok)
	}
	if _, ok := body["isPurposeRequirementInputValue"]; ok {
		t.Errorf("isPurposeRequirementInputValue must stay omitted: %s", (*recorded)[1].body)
	}

	// The two names are independent: setting one does not emit the other.
	if _, err := c.UpdateExpensePurpose("5", UpdatePurposeRequest{
		PurposeId:                      3,
		IsPurposeRequirementInputValue: ptr(true),
	}); err != nil {
		t.Fatalf("UpdateExpensePurpose(prefixed): %v", err)
	}
	body = decodeBody(t, (*recorded)[2].body)
	if body["isPurposeRequirementInputValue"] != true {
		t.Errorf("isPurposeRequirementInputValue = %v, want true", body["isPurposeRequirementInputValue"])
	}
	if _, ok := body["purposeRequirementInputValue"]; ok {
		t.Errorf("purposeRequirementInputValue must stay omitted: %s", (*recorded)[2].body)
	}
}
