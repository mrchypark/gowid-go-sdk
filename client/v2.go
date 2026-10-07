package client

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// maxPageSizeV2 is the documented maximum for GET /v2/expense-statements and
// GET /v2/cards. Exceeding it is a 400 error. The deprecated GET /v2/expenses
// documents no maximum, so it is not capped here.
const maxPageSizeV2 = 100

type ResponseV2[T any] struct {
	Result     Result `json:"result"`
	TotalCount int64  `json:"totalCount"`
	Data       T      `json:"data"`
}

type CategoryV2 struct {
	CategoryID int64  `json:"categoryId"`
	Name       string `json:"name"`
}

type UserV2 struct {
	UserName          string  `json:"userName"`
	Email             string  `json:"email"`
	MobileNumber      string  `json:"mobileNumber"`
	IsInvitedUser     bool    `json:"isInvitedUser"`
	IsContractor      bool    `json:"isContractor"`
	IsActivated       bool    `json:"isActivated"`
	Position          string  `json:"position"`
	ActivatedAt       string  `json:"activatedAt"`
	DeactivatedAt     *string `json:"deactivatedAt"`
	NotificationOnOff bool    `json:"notificationOnOff"`
	Status            string  `json:"status"`
}

// CardV2 mirrors the expense detail CardVo, which is unrelated to the
// /v2/cards CardDtoV2 below.
type CardV2 struct {
	CardNumber          string `json:"cardNumber"`
	EncryptedCardNumber string `json:"encryptedCardNumber"`
	CardUser            UserV2 `json:"cardUser"`
	Alias               string `json:"alias"`
	LimitAmount         int64  `json:"limitAmount"`
	UsedAmount          int64  `json:"usedAmount"`
	RemainAmount        int64  `json:"remainAmount"`
	CompanyCode         string `json:"companyCode"`
	NamedYN             string `json:"namedYn"`
	CardName            string `json:"cardName"`
	CardType            string `json:"cardType"`
	UserName            string `json:"userNm"`
	DuplicationStatus   string `json:"duplicationStatus"`
	Invalid             bool   `json:"invalid"`
}

// CardUserDtoV2 is the primary (대표) card holder.
type CardUserDtoV2 struct {
	UserID         int64  `json:"userId"`
	Name           string `json:"name"`
	DepartmentName string `json:"departmentName"`
}

// CardHolderUserDtoV2 is a joint holder that is a user.
type CardHolderUserDtoV2 struct {
	UserID         int64  `json:"userId"`
	Name           string `json:"name"`
	DepartmentName string `json:"departmentName"`
}

// CardHolderDepartmentDtoV2 is a joint holder that is a department.
type CardHolderDepartmentDtoV2 struct {
	DepartmentID int64  `json:"departmentId"`
	Name         string `json:"name"`
}

// CardDtoV2 is a company card with limits and the full holder graph.
type CardDtoV2 struct {
	CardID              int64                       `json:"cardId"`
	CompanyCode         string                      `json:"companyCode"`
	CardStatus          string                      `json:"cardStatus"`
	CardAlias           string                      `json:"cardAlias"`
	ManagementNumber    string                      `json:"managementNumber"`
	MaskedCardNumber    string                      `json:"maskedCardNumber"`
	EncryptedCardNumber string                      `json:"encryptedCardNumber"`
	LimitAmount         int64                       `json:"limitAmount"`
	UsedAmount          int64                       `json:"usedAmount"`
	RemainAmount        int64                       `json:"remainAmount"`
	CardUser            CardUserDtoV2               `json:"cardUser"`
	UserHolders         []CardHolderUserDtoV2       `json:"userHolders"`
	DepartmentHolders   []CardHolderDepartmentDtoV2 `json:"departmentHolders"`
}

type CardPageableDtoV2 struct {
	TotalPages    int64       `json:"totalPages"`
	TotalElements int64       `json:"totalElements"`
	Last          bool        `json:"last"`
	Content       []CardDtoV2 `json:"content"`
}

type PurposeRequirementV2 struct {
	ID               int64  `json:"id"`
	Type             string `json:"type"`
	Item             string `json:"item"`
	GuideDescription string `json:"guideDesc"`
	AvailableInput   bool   `json:"isAvailableInput"`
	Required         bool   `json:"isRequired"`
}

type PurposeRequirementAnswerV2 struct {
	PurposeRequirementID   int64    `json:"purposeRequirementId"`
	PurposeRequirementName string   `json:"purposeRequirementName"`
	Answers                []string `json:"answers"`
}

type PurposeSimpleV2 struct {
	PurposeID    int64                  `json:"purposeId"`
	Name         string                 `json:"name"`
	LimitType    string                 `json:"limitType"`
	LimitAmount  int64                  `json:"limitAmount"`
	Requirements []PurposeRequirementV2 `json:"requirements"`
}

type PurposeV2 struct {
	PurposeID    int64                  `json:"purposeId"`
	Name         string                 `json:"name"`
	Category     CategoryV2             `json:"category"`
	ListOrder    int                    `json:"listOrder"`
	LimitType    string                 `json:"limitType"`
	LimitAmount  int64                  `json:"limitAmount"`
	IsActivated  bool                   `json:"isActivated"`
	Requirements []PurposeRequirementV2 `json:"requirements"`
	Deducted     bool                   `json:"deducted"`
}

type PurposeDetailV2 struct {
	Name           string     `json:"name"`
	Category       CategoryV2 `json:"category"`
	LimitAmount    int64      `json:"limitAmount"`
	ListOrder      int        `json:"listOrder"`
	IsActivated    bool       `json:"isActivated"`
	HasRequirement bool       `json:"hasRequirement"`
	LimitType      string     `json:"limitType"`
}

type ExpenseParticipantV2 struct {
	UserID   int64  `json:"userId"`
	UserName string `json:"userName"`
}

type ExpenseExternalUserV2 struct {
	Name    string `json:"name"`
	Company string `json:"company"`
}

type CommentV2 struct {
	Author    UserV2 `json:"author"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

type ExpenseEvidenceV2 struct {
	EvidenceID int64  `json:"evidenceId"`
	FileName   string `json:"fileName"`
	MIMEType   string `json:"mimeType"`
	SignedURL  string `json:"signedUrl"`
}

type ExpenseDeductionV2 struct {
	ExpenseDeductible bool `json:"isExpenseDeductible"`
	Deducted          bool `json:"isDeducted"`
}

type ExpenseV2 struct {
	ExpenseID                 int64                        `json:"expenseId"`
	CardApprovalNumber        string                       `json:"cardApprovalNumber"`
	ExpenseDate               string                       `json:"expenseDate"`
	ExpenseTime               string                       `json:"expenseTime"`
	ExpenseType               string                       `json:"expenseType"`
	UseAmount                 float64                      `json:"useAmount"`
	Currency                  string                       `json:"currency"`
	KRWAmount                 int64                        `json:"krwAmount"`
	ApprovedAmount            *int64                       `json:"approvedAmount"`
	ApprovedAt                *string                      `json:"approvedAt"`
	ApprovalStatus            string                       `json:"approvalStatus"`
	Purpose                   *PurposeSimpleV2             `json:"purpose"`
	CardAlias                 *string                      `json:"cardAlias"`
	CardUserName              string                       `json:"cardUserName"`
	ShortCardNumber           string                       `json:"shortCardNumber"`
	EncryptedCardNumber       string                       `json:"encryptedCardNumber"`
	StoreName                 string                       `json:"storeName"`
	StoreAddress              string                       `json:"storeAddress"`
	Memo                      *string                      `json:"memo"`
	CommentCount              int                          `json:"commentCount"`
	EvidenceCount             int                          `json:"evidenceCount"`
	ParticipantCount          int                          `json:"participantCount"`
	RepresentativeParticipant *string                      `json:"representativeParticipant"`
	Participants              []ExpenseParticipantV2       `json:"participants"`
	ExpenseExternalUsers      []ExpenseExternalUserV2      `json:"expenseExternalUsers"`
	PurposeRequirementAnswers []PurposeRequirementAnswerV2 `json:"purposeRequirementAnswers"`
}

type ExpensePageV2 struct {
	TotalPages    int64       `json:"totalPages"`
	TotalElements int64       `json:"totalElements"`
	Last          bool        `json:"last"`
	Content       []ExpenseV2 `json:"content"`
}

type ExpenseSimpleV2 struct {
	ExpenseID       int64   `json:"expenseId"`
	ExpenseDate     string  `json:"expenseDate"`
	ExpenseTime     string  `json:"expenseTime"`
	UseAmount       float64 `json:"useAmount"`
	Currency        string  `json:"currency"`
	KRWAmount       int64   `json:"krwAmount"`
	StoreName       string  `json:"storeName"`
	ApprovalStatus  string  `json:"approvalStatus"`
	CardAlias       *string `json:"cardAlias"`
	ShortCardNumber string  `json:"shortCardNumber"`
}

type ExpenseSimplePageV2 struct {
	TotalPages    int64             `json:"totalPages"`
	TotalElements int64             `json:"totalElements"`
	Last          bool              `json:"last"`
	Content       []ExpenseSimpleV2 `json:"content"`
}

type ExpenseDetailV2 struct {
	ExpenseID                 int64                        `json:"expenseId"`
	CardApprovalNumber        string                       `json:"cardApprovalNumber"`
	ExpenseType               string                       `json:"expenseType"`
	ExpenseDate               string                       `json:"expenseDate"`
	ExpenseTime               string                       `json:"expenseTime"`
	Card                      CardV2                       `json:"card"`
	User                      UserV2                       `json:"user"`
	UseAmount                 float64                      `json:"useAmount"`
	Currency                  string                       `json:"currency"`
	KRWAmount                 int64                        `json:"krwAmount"`
	ApprovalStatus            string                       `json:"approvalStatus"`
	ApprovedAmount            *int64                       `json:"approvedAmount"`
	ApprovedAt                *string                      `json:"approvedAt"`
	ApprovedBy                *string                      `json:"approvedBy"`
	Comments                  []CommentV2                  `json:"comments"`
	Purpose                   *PurposeDetailV2             `json:"purpose"`
	Participants              []UserV2                     `json:"participants"`
	ExpenseExternalUsers      []ExpenseExternalUserV2      `json:"expenseExternalUsers"`
	StoreName                 string                       `json:"storeName"`
	StoreAddress              string                       `json:"storeAddress"`
	StoreRegistrationNumber   *string                      `json:"storeRegistrationNumber"`
	Memo                      *string                      `json:"memo"`
	EvidenceList              []ExpenseEvidenceV2          `json:"evidenceList"`
	CompanyCode               *string                      `json:"companyCode"`
	CommentCount              *int                         `json:"commentCount"`
	ExpenseDeduction          ExpenseDeductionV2           `json:"expenseDeductionResDto"`
	PurposeRequirementAnswers []PurposeRequirementAnswerV2 `json:"purposeRequirementAnswers"`
	IsDomestic                bool                         `json:"isDomestic"`
}

// ExpenseSearchOptionsV2 filters the v2 expense list endpoints. StartDate and
// EndDate are yyyyMMdd and inclusive. The deprecated SearchExpensesV2 has no
// endDate parameter and rejects this field when it is set.
type ExpenseSearchOptionsV2 struct {
	Memo, PurposeName, UserName, StartDate, EndDate, ApprovalState, EncryptedCardNumber string
	Page, Size                                                                          int
	Sort                                                                                string
}

type PageOptionsV2 struct {
	Page, Size int
	Sort       string
}

// UpdateExpenseRequestV2 updates purpose, memo and participants in one call.
// Nil slices, maps and pointers are omitted; an explicit empty slice or map is
// sent as an empty array or object, and a non-nil pointer to the zero value
// emits that zero value on the wire. How the server treats an omitted field on
// this PUT is not specified, so do not read omission as "retains the old value".
type UpdateExpenseRequestV2 struct {
	PurposeID                 *int64                  `json:"purposeId,omitzero"`
	ParticipantIDs            []int64                 `json:"participantIdList,omitzero"`
	ExternalUsers             []ExpenseExternalUserV2 `json:"externalUserList,omitzero"`
	PurposeRequirementAnswers map[string][]string     `json:"purposeRequirementAnswerMap,omitzero"`
	Memo                      *string                 `json:"memo,omitzero"`
	FileIDs                   []int64                 `json:"fileIdList,omitzero"`
}

type UpdatePurposeRequestV2 struct {
	PurposeID                 int64               `json:"purposeId"`
	PurposeRequirementAnswers map[string][]string `json:"purposeRequirementAnswerMap,omitzero"`
}

type UpdatePurposesRequestV2 struct {
	ExpenseIDs                []int64             `json:"expenseIds"`
	PurposeID                 int64               `json:"purposeId"`
	PurposeRequirementAnswers map[string][]string `json:"purposeRequirementAnswerMap,omitzero"`
}

type UpdateParticipantsRequestV2 struct {
	ParticipantIDs []int64                 `json:"participantIds"`
	ExternalUsers  []ExpenseExternalUserV2 `json:"externalUsers"`
}

type ApproveExpenseRequestV2 struct {
	ExpenseID      int64   `json:"expenseId,omitempty"`
	ApprovalStatus string  `json:"approvalStatus"`
	ApprovedAmount *int64  `json:"approvedAmount,omitempty"`
	ApprovedAt     *string `json:"approvedAt,omitempty"`
}

type BulkApproveResultV2 struct {
	SucceedStatements []ExpenseSimpleV2 `json:"succeedStatements"`
	FailedStatements  []ExpenseSimpleV2 `json:"failedStatements"`
}

// updateExpenseBodyV2 is the wire form of UpdateExpenseRequestV2: expenseId
// always comes from the path and isIgnoredFiles is always true because the
// Open API does not support file updates.
type updateExpenseBodyV2 struct {
	UpdateExpenseRequestV2
	ExpenseID      int64 `json:"expenseId"`
	IsIgnoredFiles bool  `json:"isIgnoredFiles"`
}

func (c *Client) GetCardsV2(opts *PageOptionsV2) (*ResponseV2[CardPageableDtoV2], error) {
	q := url.Values{}
	if opts != nil {
		if opts.Size > maxPageSizeV2 {
			return nil, fmt.Errorf("size %d exceeds documented maximum %d", opts.Size, maxPageSizeV2)
		}
		addPageOptionsV2(q, opts.Page, opts.Size, opts.Sort)
	}
	return getV2[CardPageableDtoV2](c, "/v2/cards", q)
}

func (c *Client) GetExpenseV2(expenseID int64) (*ResponseV2[ExpenseDetailV2], error) {
	return getV2[ExpenseDetailV2](c, fmt.Sprintf("/v2/expenses/%d", expenseID), nil)
}

func (c *Client) GetExpenseStatementsV2(opts *ExpenseSearchOptionsV2) (*ResponseV2[ExpensePageV2], error) {
	if opts != nil && opts.Size > maxPageSizeV2 {
		return nil, fmt.Errorf("size %d exceeds documented maximum %d", opts.Size, maxPageSizeV2)
	}
	return getV2[ExpensePageV2](c, "/v2/expense-statements", expenseSearchQueryV2(opts))
}

// SearchExpensesV2 calls the deprecated /v2/expenses endpoint. That endpoint
// takes no endDate: the period is fixed to roughly one month after startDate
// and the computed end day is exclusive, so a supplied EndDate is rejected
// instead of being silently dropped. Prefer GetExpenseStatementsV2.
func (c *Client) SearchExpensesV2(opts *ExpenseSearchOptionsV2) (*ResponseV2[ExpensePageV2], error) {
	if opts != nil && opts.EndDate != "" {
		return nil, errors.New("endDate is not supported by the deprecated /v2/expenses endpoint; use GetExpenseStatementsV2")
	}
	return getV2[ExpensePageV2](c, "/v2/expenses", expenseSearchQueryV2(opts))
}

func (c *Client) GetNotSubmittedExpensesV2(opts *PageOptionsV2) (*ResponseV2[ExpenseSimplePageV2], error) {
	q := url.Values{}
	if opts != nil {
		addPageOptionsV2(q, opts.Page, opts.Size, opts.Sort)
	}
	return getV2[ExpenseSimplePageV2](c, "/v2/expenses/not-submitted", q)
}

func (c *Client) GetPurposesV2(isActivated *bool) (*ResponseV2[[]PurposeV2], error) {
	q := url.Values{}
	if isActivated != nil {
		q.Set("isActivated", strconv.FormatBool(*isActivated))
	}
	return getV2[[]PurposeV2](c, "/v2/purposes", q)
}

func (c *Client) GetPurposeRequirementsV2(purposeID, requirementID int64) (*ResponseV2[struct {
	Content []string `json:"content"`
}], error) {
	return getV2[struct {
		Content []string `json:"content"`
	}](c, fmt.Sprintf("/v2/purposes/%d/requirements/%d", purposeID, requirementID), nil)
}

func (c *Client) UpdateExpenseV2(expenseID int64, input UpdateExpenseRequestV2) (*ResponseV2[ExpenseDetailV2], error) {
	body := updateExpenseBodyV2{
		UpdateExpenseRequestV2: input,
		ExpenseID:              expenseID,
		IsIgnoredFiles:         true,
	}
	body.PurposeRequirementAnswers = answersOrEmpty(body.PurposeRequirementAnswers)
	return doV2[ExpenseDetailV2](c, http.MethodPut, fmt.Sprintf("/v2/expenses/%d", expenseID), body)
}

func (c *Client) UpdateExpensePurposeV2(expenseID int64, input UpdatePurposeRequestV2) (*ResponseV2[ExpenseDetailV2], error) {
	input.PurposeRequirementAnswers = answersOrEmpty(input.PurposeRequirementAnswers)
	return doV2[ExpenseDetailV2](c, http.MethodPut, fmt.Sprintf("/v2/expenses/%d/purposes", expenseID), input)
}

func (c *Client) UpdateExpensesPurposeV2(input UpdatePurposesRequestV2) (*ResponseV2[bool], error) {
	if len(input.ExpenseIDs) == 0 {
		return nil, errors.New("expenseIds must not be empty")
	}
	seen := make(map[int64]bool, len(input.ExpenseIDs))
	for _, id := range input.ExpenseIDs {
		if seen[id] {
			return nil, fmt.Errorf("expenseIds contains duplicate id %d", id)
		}
		seen[id] = true
	}
	input.PurposeRequirementAnswers = answersOrEmpty(input.PurposeRequirementAnswers)
	return doV2[bool](c, http.MethodPut, "/v2/expenses/purposes", input)
}

func (c *Client) UpdateExpenseParticipantsV2(expenseID int64, input UpdateParticipantsRequestV2) (*ResponseV2[ExpenseDetailV2], error) {
	if input.ParticipantIDs == nil {
		input.ParticipantIDs = []int64{}
	}
	if input.ExternalUsers == nil {
		input.ExternalUsers = []ExpenseExternalUserV2{}
	}
	return doV2[ExpenseDetailV2](c, http.MethodPut, fmt.Sprintf("/v2/expenses/%d/participants", expenseID), input)
}

func (c *Client) UpdateExpenseMemoV2(expenseID int64, memo string) (*ResponseV2[ExpenseDetailV2], error) {
	return doV2[ExpenseDetailV2](c, http.MethodPut, fmt.Sprintf("/v2/expenses/%d/memo", expenseID), UpdateMemoRequest{Memo: memo})
}

func (c *Client) UpdateExpenseApprovalStatusV2(expenseID int64, input ApproveExpenseRequestV2) (*ResponseV2[ExpenseDetailV2], error) {
	if err := validApprovalStatusV2(input.ApprovalStatus); err != nil {
		return nil, err
	}
	return doV2[ExpenseDetailV2](c, http.MethodPut, fmt.Sprintf("/v2/expenses/%d/approval-status", expenseID), input)
}

func (c *Client) ApproveExpensesV2(input []ApproveExpenseRequestV2) (*ResponseV2[BulkApproveResultV2], error) {
	for _, item := range input {
		if err := validApprovalStatusV2(item.ApprovalStatus); err != nil {
			return nil, err
		}
		// Bulk approval has no path id, so each item carries its own expenseId.
		if item.ExpenseID <= 0 {
			return nil, fmt.Errorf("bulk approval item requires a positive expenseId, got %d", item.ExpenseID)
		}
	}
	return doV2[BulkApproveResultV2](c, http.MethodPatch, "/v2/expenses/approval-status/approved", input)
}

func getV2[T any](c *Client, path string, query url.Values) (*ResponseV2[T], error) {
	req, err := c.newRequest(http.MethodGet, c.BaseURL+path, nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = query.Encode()
	var response ResponseV2[T]
	if err := c.Do(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func doV2[T any](c *Client, method, path string, body any) (*ResponseV2[T], error) {
	req, err := c.newRequest(method, c.BaseURL+path, body)
	if err != nil {
		return nil, err
	}
	var response ResponseV2[T]
	if err := c.Do(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func expenseSearchQueryV2(opts *ExpenseSearchOptionsV2) url.Values {
	q := url.Values{}
	if opts == nil {
		return q
	}
	for key, value := range map[string]string{
		"memo": opts.Memo, "purposeName": opts.PurposeName, "userName": opts.UserName,
		"startDate": opts.StartDate, "endDate": opts.EndDate, "approvalState": opts.ApprovalState,
		"encryptedCardNumber": opts.EncryptedCardNumber,
	} {
		if value != "" {
			q.Set(key, value)
		}
	}
	addPageOptionsV2(q, opts.Page, opts.Size, opts.Sort)
	return q
}

// answersOrEmpty keeps a non-nil map from serializing nil values, which would
// send invalid null answer arrays. A nil map stays nil so it is omitted. The
// caller's map is never mutated.
func answersOrEmpty(answers map[string][]string) map[string][]string {
	var copied map[string][]string
	for key, value := range answers {
		if value == nil {
			if copied == nil {
				copied = make(map[string][]string, len(answers))
				for k, v := range answers {
					copied[k] = v
				}
			}
			copied[key] = []string{}
		}
	}
	if copied != nil {
		return copied
	}
	return answers
}

// validApprovalStatusV2 restricts approval changes to the two documented
// statuses; PARTIAL_APPROVED is decided by the server from approvedAmount.
func validApprovalStatusV2(status string) error {
	if status != "APPROVED" && status != "REJECTED" {
		return fmt.Errorf("approvalStatus %q is not requestable; use APPROVED or REJECTED", status)
	}
	return nil
}

func addPageOptionsV2(q url.Values, page, size int, sort string) {
	if page > 0 {
		q.Set("page", strconv.Itoa(page))
	}
	if size > 0 {
		q.Set("size", strconv.Itoa(size))
	}
	if sort != "" {
		q.Set("sort", sort)
	}
}
