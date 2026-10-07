package client

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// --------------------------------------------------------------------------
// Shared response DTOs
// --------------------------------------------------------------------------

// SimpleUser mirrors UserVo, the user shape embedded in expense detail.
type SimpleUser struct {
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

// SimplePurpose mirrors PurposeSimpleDto.
type SimplePurpose struct {
	PurposeId      int64  `json:"purposeId"`
	Name           string `json:"name"`
	LimitType      string `json:"limitType"`
	LimitAmount    int64  `json:"limitAmount"`
	IsActivated    bool   `json:"isActivated"`
	HasRequirement bool   `json:"hasRequirement"`
}

// ExpensePurpose mirrors PurposeVo: the purpose attached to an expense detail.
type ExpensePurpose struct {
	Name           string   `json:"name"`
	Category       Category `json:"category"`
	LimitAmount    int64    `json:"limitAmount"`
	ListOrder      int      `json:"listOrder"`
	IsActivated    bool     `json:"isActivated"`
	HasRequirement bool     `json:"hasRequirement"`
	LimitType      string   `json:"limitType"`
}

type ExpenseParticipant struct {
	UserId   int64  `json:"userId"`
	UserName string `json:"userName"`
}

type ExpenseExternalUser struct {
	Name    string `json:"name"`
	Company string `json:"company"`
}

type ExpenseComment struct {
	Author    SimpleUser `json:"author"`
	Content   string     `json:"content"`
	CreatedAt string     `json:"createdAt"`
}

type ExpenseEvidence struct {
	EvidenceId int64  `json:"evidenceId"`
	FileName   string `json:"fileName"`
	MimeType   string `json:"mimeType"`
	SignedUrl  string `json:"signedUrl"`
}

type ExpenseDeduction struct {
	IsExpenseDeductible bool `json:"isExpenseDeductible"`
	IsDeducted          bool `json:"isDeducted"`
}

type ExpenseCardDetail struct {
	CardNumber          string     `json:"cardNumber"`
	EncryptedCardNumber string     `json:"encryptedCardNumber"`
	CardUser            SimpleUser `json:"cardUser"`
	Alias               string     `json:"alias"`
	LimitAmount         int64      `json:"limitAmount"`
	UsedAmount          int64      `json:"usedAmount"`
	RemainAmount        int64      `json:"remainAmount"`
	CompanyCode         string     `json:"companyCode"`
	NamedYn             string     `json:"namedYn"`
	CardName            string     `json:"cardName"`
	CardType            string     `json:"cardType"`
	UserNm              string     `json:"userNm"`
	DuplicationStatus   string     `json:"duplicationStatus"`
	Invalid             bool       `json:"invalid"`
}

// ExpenseSummary mirrors ExpenseDto, the row shape of GET /v1/expenses.
type ExpenseSummary struct {
	ExpenseId                      int64                 `json:"expenseId"`
	ExpenseDate                    string                `json:"expenseDate"`
	ExpenseTime                    string                `json:"expenseTime"`
	UseAmount                      float64               `json:"useAmount"`
	Currency                       string                `json:"currency"`
	KrwAmount                      int64                 `json:"krwAmount"`
	ApprovedAmount                 *int64                `json:"approvedAmount"`
	ApprovedAt                     *string               `json:"approvedAt"`
	ApprovalStatus                 string                `json:"approvalStatus"`
	Purpose                        *SimplePurpose        `json:"purpose"`
	CardAlias                      *string               `json:"cardAlias"`
	CardUserName                   string                `json:"cardUserName"`
	ShortCardNumber                string                `json:"shortCardNumber"`
	StoreName                      string                `json:"storeName"`
	StoreAddress                   string                `json:"storeAddress"`
	Memo                           *string               `json:"memo"`
	CommentCount                   int                   `json:"commentCount"`
	EvidenceCount                  int                   `json:"evidenceCount"`
	ParticipantCount               int                   `json:"participantCount"`
	RepresentativeParticipant      *string               `json:"representativeParticipant"`
	Participants                   []ExpenseParticipant  `json:"participants"`
	ExpenseExternalUsers           []ExpenseExternalUser `json:"expenseExternalUsers"`
	PurposeRequirementItem         *string               `json:"purposeRequirementItem"`
	PurposeRequirementItemType     *string               `json:"purposeRequirementItemType"`
	PurposeRequirementValue        *string               `json:"purposeRequirementValue"`
	IsPurposeRequirementInputValue bool                  `json:"isPurposeRequirementInputValue"`
}

type ExpensePage struct {
	TotalPages    int              `json:"totalPages"`
	TotalElements int              `json:"totalElements"`
	Last          bool             `json:"last"`
	Content       []ExpenseSummary `json:"content"`
}

type GetExpensesResponse Response[ExpensePage]

// ExpenseSimple mirrors ExpenseSimpleDto, the row shape of the bulk approve
// response and of GET /v1/expenses/not-submitted.
type ExpenseSimple struct {
	ExpenseId              int64     `json:"expenseId"`
	ExpenseDate            string    `json:"expenseDate"`
	ExpenseTime            string    `json:"expenseTime"`
	UseAmount              float64   `json:"useAmount"`
	Currency               string    `json:"currency"`
	KrwAmount              int64     `json:"krwAmount"`
	StoreName              string    `json:"storeName"`
	ApprovalStatus         string    `json:"approvalStatus"`
	CardAlias              *string   `json:"cardAlias"`
	ShortCardNumber        string    `json:"shortCardNumber"`
	RecommendedPurposeList []Purpose `json:"recommendedPurposeList"`
}

type ExpenseSimplePage struct {
	TotalPages    int             `json:"totalPages"`
	TotalElements int             `json:"totalElements"`
	Last          bool            `json:"last"`
	Content       []ExpenseSimple `json:"content"`
}

type GetNotSubmittedExpensesResponse Response[ExpenseSimplePage]

// ExpenseDetail mirrors ExpenseDetailResDto.
type ExpenseDetail struct {
	ExpenseId                      int64                 `json:"expenseId"`
	CardApprovalNumber             string                `json:"cardApprovalNumber"`
	ExpenseDate                    string                `json:"expenseDate"`
	ExpenseTime                    string                `json:"expenseTime"`
	Card                           ExpenseCardDetail     `json:"card"`
	User                           SimpleUser            `json:"user"`
	UseAmount                      float64               `json:"useAmount"`
	Currency                       string                `json:"currency"`
	KrwAmount                      int64                 `json:"krwAmount"`
	ApprovalStatus                 string                `json:"approvalStatus"`
	ApprovedAmount                 *int64                `json:"approvedAmount"`
	ApprovedAt                     *string               `json:"approvedAt"`
	ApprovedBy                     *string               `json:"approvedBy"`
	Comments                       []ExpenseComment      `json:"comments"`
	Purpose                        *ExpensePurpose       `json:"purpose"`
	Participants                   []SimpleUser          `json:"participants"`
	ExpenseExternalUsers           []ExpenseExternalUser `json:"expenseExternalUsers"`
	StoreName                      string                `json:"storeName"`
	StoreAddress                   string                `json:"storeAddress"`
	StoreRegistrationNumber        *string               `json:"storeRegistrationNumber"`
	Memo                           *string               `json:"memo"`
	EvidenceList                   []ExpenseEvidence     `json:"evidenceList"`
	CompanyCode                    *string               `json:"companyCode"`
	CommentCount                   *int                  `json:"commentCount"`
	PurposeRequirementItem         *string               `json:"purposeRequirementItem"`
	PurposeRequirementItemType     *string               `json:"purposeRequirementItemType"`
	PurposeRequirementValue        *string               `json:"purposeRequirementValue"`
	ExpenseDeductionResDto         ExpenseDeduction      `json:"expenseDeductionResDto"`
	IsDomestic                     bool                  `json:"isDomestic"`
	IsPurposeRequirementInputValue bool                  `json:"isPurposeRequirementInputValue"`
}

type GetExpenseResponse Response[ExpenseDetail]

// --------------------------------------------------------------------------
// Expense ID validation (shared by every v1 expense path call)
// --------------------------------------------------------------------------

// parseExpenseID validates the string expense id accepted by the v1 methods
// and returns it as the int64 the official API documents. Rejecting non-numeric
// and non-positive ids here keeps a bad id from ever reaching the network.
func parseExpenseID(expenseId string) (int64, error) {
	id, err := strconv.ParseInt(expenseId, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid expenseId %q: must be a positive int64: %w", expenseId, err)
	}
	if id <= 0 {
		return 0, fmt.Errorf("invalid expenseId %q: must be a positive int64", expenseId)
	}
	return id, nil
}

func addPageQuery(q url.Values, opts *PageOptions) {
	if opts == nil {
		return
	}
	if opts.Page >= 0 {
		q.Add("page", strconv.Itoa(opts.Page))
	}
	if opts.Size > 0 {
		q.Add("size", strconv.Itoa(opts.Size))
	}
	if opts.Sort != "" {
		q.Add("sort", opts.Sort)
	}
}

// --------------------------------------------------------------------------
// Reads
// --------------------------------------------------------------------------

// GetExpensesOptions carries the ExpenseSearchCriteria fields as flat query
// parameters, plus the page/size/sort paging parameters.
type GetExpensesOptions struct {
	ApprovalState string
	Memo          string
	PurposeName   string
	UserName      string
	StartDate     string // yyyyMMdd
	Sort          string // e.g. expenseDate,desc
	Size          int
	Page          int
}

func (c *Client) GetExpenses(opts *GetExpensesOptions) (*GetExpensesResponse, error) {
	url := fmt.Sprintf("%s/v1/expenses", c.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if opts != nil {
		if opts.ApprovalState != "" {
			q.Add("approvalState", opts.ApprovalState)
		}
		if opts.Memo != "" {
			q.Add("memo", opts.Memo)
		}
		if opts.PurposeName != "" {
			q.Add("purposeName", opts.PurposeName)
		}
		if opts.UserName != "" {
			q.Add("userName", opts.UserName)
		}
		if opts.StartDate != "" {
			q.Add("startDate", opts.StartDate)
		}
		addPageQuery(q, &PageOptions{Page: opts.Page, Size: opts.Size, Sort: opts.Sort})
	}
	req.URL.RawQuery = q.Encode()

	var resp GetExpensesResponse
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetNotSubmittedExpenses lists expenses that still need a purpose, with the
// server's recommended purpose list per row.
func (c *Client) GetNotSubmittedExpenses(opts *PageOptions) (*GetNotSubmittedExpensesResponse, error) {
	url := fmt.Sprintf("%s/v1/expenses/not-submitted", c.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	addPageQuery(q, opts)
	req.URL.RawQuery = q.Encode()

	var resp GetNotSubmittedExpensesResponse
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetExpense(expenseId string) (*GetExpenseResponse, error) {
	id, err := parseExpenseID(expenseId)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/v1/expenses/%d", c.BaseURL, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var resp GetExpenseResponse
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// --------------------------------------------------------------------------
// Purpose updates
// --------------------------------------------------------------------------

// UpdatePurposeRequest mirrors ExpensePurposeUpdateRequestDto. Every property
// except PurposeId is optional; pointers distinguish "not supplied" from an
// explicit zero value.
type UpdatePurposeRequest struct {
	ExpenseId                      *int64  `json:"expenseId,omitempty"`
	PurposeId                      int64   `json:"purposeId"`
	PurposeRequirementItem         *string `json:"purposeRequirementItem,omitempty"`
	PurposeRequirementItemType     *string `json:"purposeRequirementItemType,omitempty"`
	PurposeRequirementValue        *string `json:"purposeRequirementValue,omitempty"`
	IsPurposeRequirementInputValue *bool   `json:"isPurposeRequirementInputValue,omitempty"`
	// PurposeRequirementInputValue is the bare boolean the published schema also
	// lists beside IsPurposeRequirementInputValue. Both names are represented so
	// either can be sent, but the docs never state which the server reads or
	// whether they differ, so the meaning is unresolved.
	PurposeRequirementInputValue *bool `json:"purposeRequirementInputValue,omitempty"`
}

type UpdateExpensePurposeResponse = Response[ExpenseDetail]

func (c *Client) UpdateExpensePurpose(expenseId string, req UpdatePurposeRequest) (*UpdateExpensePurposeResponse, error) {
	id, err := parseExpenseID(expenseId)
	if err != nil {
		return nil, err
	}
	if req.ExpenseId != nil && *req.ExpenseId != id {
		return nil, fmt.Errorf("expenseId mismatch: path %d, body %d", id, *req.ExpenseId)
	}
	req.ExpenseId = &id
	url := fmt.Sprintf("%s/v1/expenses/%d/purposes", c.BaseURL, id)
	return c.doUpdateExpense(url, "PUT", req)
}

// UpdateExpensesPurposeRequest is the bulk form of ExpensePurposeUpdateRequestDto.
// The shared schema marks expenseId optional, but the endpoint is documented as
// a per-item purpose update ("V1은 항목별로 서로 다른 용도를 지정하는 방식"),
// so every item must carry its expenseId: without it the server has no expense
// to apply the purpose to. UpdateExpensesPurpose rejects items that omit it.
type UpdateExpensesPurposeRequest []UpdatePurposeRequest

type UpdateExpensesPurposeResponse = Response[bool]

func (c *Client) UpdateExpensesPurpose(req UpdateExpensesPurposeRequest) (*UpdateExpensesPurposeResponse, error) {
	for i := range req {
		if req[i].ExpenseId == nil {
			return nil, fmt.Errorf("bulk purpose item %d: expenseId is required", i)
		}
		if *req[i].ExpenseId <= 0 {
			return nil, fmt.Errorf("invalid expenseId at index %d: must be a positive int64", i)
		}
	}
	url := fmt.Sprintf("%s/v1/expenses/purposes", c.BaseURL)
	httpReq, err := c.newRequest("PUT", url, req)
	if err != nil {
		return nil, err
	}
	var resp UpdateExpensesPurposeResponse
	if err := c.Do(httpReq, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// --------------------------------------------------------------------------
// Memo
// --------------------------------------------------------------------------

type UpdateMemoRequest struct {
	Memo string `json:"memo"`
}

type UpdateExpenseMemoResponse = Response[ExpenseDetail]

func (c *Client) UpdateExpenseMemo(expenseId string, memo string) (*UpdateExpenseMemoResponse, error) {
	id, err := parseExpenseID(expenseId)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/v1/expenses/%d/memo", c.BaseURL, id)
	return c.doUpdateExpense(url, "PUT", UpdateMemoRequest{Memo: memo})
}

// --------------------------------------------------------------------------
// Participants
// --------------------------------------------------------------------------

// UpdateParticipantsRequest mirrors ExpenseParticipantsUpdateRequestDto. Both
// arrays are required, so nil is normalized to an empty list before sending.
type UpdateParticipantsRequest struct {
	ExternalUsers  []ExpenseExternalUser `json:"externalUsers"`
	ParticipantIds []int64               `json:"participantIds"`
}

type UpdateExpenseParticipantsResponse = Response[ExpenseDetail]

func (c *Client) UpdateExpenseParticipants(expenseId string, req UpdateParticipantsRequest) (*UpdateExpenseParticipantsResponse, error) {
	id, err := parseExpenseID(expenseId)
	if err != nil {
		return nil, err
	}
	if req.ExternalUsers == nil {
		req.ExternalUsers = []ExpenseExternalUser{}
	}
	if req.ParticipantIds == nil {
		req.ParticipantIds = []int64{}
	}
	url := fmt.Sprintf("%s/v1/expenses/%d/participants", c.BaseURL, id)
	return c.doUpdateExpense(url, "PUT", req)
}

// --------------------------------------------------------------------------
// Full update
// --------------------------------------------------------------------------

// UpdateExpenseRequest mirrors ExpenseUpdateReqDto. All properties are
// optional on the wire: scalar pointers plus omitzero slices distinguish an
// omitted field from an explicit empty value.
//
// Pointers and slices only control what goes on the wire: nil omits the field,
// a non-nil pointer or empty slice sends the value. The official spec does not
// say whether an omitted property leaves the stored value untouched or clears
// it, so do not rely on this endpoint to clear anything. For a targeted change
// use the dedicated endpoint (memo, purposes, participants, approval-status).
type UpdateExpenseRequest struct {
	ExpenseId                      *int64                `json:"expenseId,omitempty"`
	PurposeId                      *int64                `json:"purposeId,omitempty"`
	PurposeRequirementItem         *string               `json:"purposeRequirementItem,omitempty"`
	PurposeRequirementItemType     *string               `json:"purposeRequirementItemType,omitempty"`
	PurposeRequirementValue        *string               `json:"purposeRequirementValue,omitempty"`
	ParticipantIdList              []int64               `json:"participantIdList,omitzero"`
	ExternalUserList               []ExpenseExternalUser `json:"externalUserList,omitzero"`
	Memo                           *string               `json:"memo,omitempty"`
	FileIdList                     []int64               `json:"fileIdList,omitzero"`
	IsIgnoredFiles                 *bool                 `json:"isIgnoredFiles,omitempty"`
	IsPurposeRequirementInputValue *bool                 `json:"isPurposeRequirementInputValue,omitempty"`
}

type UpdateExpenseResponse = Response[ExpenseDetail]

func (c *Client) UpdateExpense(expenseId string, req UpdateExpenseRequest) (*UpdateExpenseResponse, error) {
	id, err := parseExpenseID(expenseId)
	if err != nil {
		return nil, err
	}
	if req.ExpenseId != nil && *req.ExpenseId != id {
		return nil, fmt.Errorf("expenseId mismatch: path %d, body %d", id, *req.ExpenseId)
	}
	req.ExpenseId = &id
	// The Open API does not support replacing evidence files, so the official
	// contract requires isIgnoredFiles=true on every full update.
	ignored := true
	req.IsIgnoredFiles = &ignored
	url := fmt.Sprintf("%s/v1/expenses/%d", c.BaseURL, id)
	return c.doUpdateExpense(url, "PUT", req)
}

// --------------------------------------------------------------------------
// Approval
// --------------------------------------------------------------------------

// UpdateApprovalStatusRequest mirrors ExpenseApproveReqDto. ApprovedAmount and
// ApprovedAt are optional: nil omits them, a pointer sends the value (including
// an explicit zero).
type UpdateApprovalStatusRequest struct {
	ExpenseId      *int64  `json:"expenseId,omitempty"`
	ApprovalStatus string  `json:"approvalStatus"`
	ApprovedAmount *int64  `json:"approvedAmount,omitempty"`
	ApprovedAt     *string `json:"approvedAt,omitempty"`
}

type UpdateExpenseApprovalStatusResponse = Response[ExpenseDetail]

func (c *Client) UpdateExpenseApprovalStatus(expenseId string, req UpdateApprovalStatusRequest) (*UpdateExpenseApprovalStatusResponse, error) {
	id, err := parseExpenseID(expenseId)
	if err != nil {
		return nil, err
	}
	if req.ApprovalStatus != "APPROVED" && req.ApprovalStatus != "REJECTED" {
		return nil, fmt.Errorf("invalid approvalStatus %q: must be APPROVED or REJECTED", req.ApprovalStatus)
	}
	if req.ExpenseId != nil && *req.ExpenseId != id {
		return nil, fmt.Errorf("expenseId mismatch: path %d, body %d", id, *req.ExpenseId)
	}
	req.ExpenseId = &id
	url := fmt.Sprintf("%s/v1/expenses/%d/approval-status", c.BaseURL, id)
	return c.doUpdateExpense(url, "PUT", req)
}

// ApproveExpensesRequest is the bulk form of ExpenseApproveReqDto.
type ApproveExpensesRequest []UpdateApprovalStatusRequest

// ApproveExpensesResult mirrors StatementBulkApproveResDto.
type ApproveExpensesResult struct {
	SucceedStatements []ExpenseSimple `json:"succeedStatements"`
	// FailedStatements is marked deprecated in the official spec.
	FailedStatements []ExpenseSimple `json:"failedStatements"`
}

type ApproveExpensesResponse Response[ApproveExpensesResult]

func (c *Client) ApproveExpenses(req ApproveExpensesRequest) (*ApproveExpensesResponse, error) {
	for i := range req {
		// There is no path id here, so every item must name its own expense.
		if req[i].ExpenseId == nil {
			return nil, fmt.Errorf("approval item %d: expenseId is required", i)
		}
		if *req[i].ExpenseId <= 0 {
			return nil, fmt.Errorf("invalid expenseId at index %d: must be a positive int64", i)
		}
		// The operation description limits requests to APPROVED/REJECTED; the
		// wider schema enum covers states the server rejects for this call.
		if req[i].ApprovalStatus != "APPROVED" && req[i].ApprovalStatus != "REJECTED" {
			return nil, fmt.Errorf("approval item %d: invalid approvalStatus %q: must be APPROVED or REJECTED", i, req[i].ApprovalStatus)
		}
	}
	url := fmt.Sprintf("%s/v1/expenses/approval-status/approved", c.BaseURL)
	httpReq, err := c.newRequest("PATCH", url, req)
	if err != nil {
		return nil, err
	}
	var resp ApproveExpensesResponse
	if err := c.Do(httpReq, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// --------------------------------------------------------------------------
// Comments
// --------------------------------------------------------------------------

type AddCommentRequest struct {
	ExpenseId *int64 `json:"expenseId,omitempty"`
	Comment   string `json:"comment"`
}

// AddCommentResponse mirrors CommentResDto.
type AddCommentResponse Response[struct {
	Author     string `json:"author"`
	Department string `json:"department"`
	Content    string `json:"content"`
	CreatedAt  string `json:"createdAt"`
}]

func (c *Client) AddComment(expenseId string, comment string) (*AddCommentResponse, error) {
	id, err := parseExpenseID(expenseId)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/v1/expenses/%d/comments", c.BaseURL, id)
	httpReq, err := c.newRequest("POST", url, AddCommentRequest{ExpenseId: &id, Comment: comment})
	if err != nil {
		return nil, err
	}
	var resp AddCommentResponse
	if err := c.Do(httpReq, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// doUpdateExpense issues a single-expense mutation, all of which return the
// full ExpenseDetailResDto.
func (c *Client) doUpdateExpense(url, method string, body interface{}) (*Response[ExpenseDetail], error) {
	req, err := c.newRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	var resp Response[ExpenseDetail]
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
