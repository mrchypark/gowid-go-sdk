package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Card struct {
	CardId       string `json:"cardId"`
	CardNumber   string `json:"cardNumber"`
	CardNickName string `json:"cardNickName"`
	OwnerName    string `json:"ownerName"`
}

type Store struct {
	Name            string `json:"name"`
	BusinessNumber  string `json:"businessNumber"`
	IndustryCode    string `json:"industryCode"`
	IndustryName    string `json:"industryName"`
	Address         string `json:"address"`
	MasterStoreName string `json:"masterStoreName"`
}

type SimpleUser struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type SimplePurpose struct {
	PurposeId int    `json:"purposeId"`
	Name      string `json:"name"`
}

type ExpensePurpose struct {
	PurposeId      int    `json:"purposeId"`
	Name           string `json:"name"`
	LimitType      string `json:"limitType"`
	LimitAmount    int    `json:"limitAmount"`
	IsActivated    bool   `json:"isActivated"`
	HasRequirement bool   `json:"hasRequirement"`
}

type ExpenseParticipant struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
}

type ExpenseExternalUser struct {
	Company string `json:"company"`
	Name    string `json:"name"`
}

type Expense struct {
	ExpenseId     string        `json:"expenseId"`
	Card          Card          `json:"card"`
	TransactionAt string        `json:"transactionAt"`
	Store         Store         `json:"store"`
	Amount        int           `json:"amount"`
	Currency      string        `json:"currency"`
	Status        string        `json:"status"`
	Memo          string        `json:"memo"`
	User          SimpleUser    `json:"user"`
	Purpose       SimplePurpose `json:"purpose"`
}

type ExpenseSummary struct {
	ExpenseId                      int                   `json:"expenseId"`
	ExpenseDate                    string                `json:"expenseDate"`
	ExpenseTime                    string                `json:"expenseTime"`
	UseAmount                      float64               `json:"useAmount"`
	Currency                       string                `json:"currency"`
	KrwAmount                      int                   `json:"krwAmount"`
	ApprovedAmount                 *int                  `json:"approvedAmount"`
	ApprovedAt                     *string               `json:"approvedAt"`
	ApprovalStatus                 string                `json:"approvalStatus"`
	Purpose                        *ExpensePurpose       `json:"purpose"`
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

type ExpenseUserDetail struct {
	UserId            int     `json:"userId"`
	UserName          string  `json:"userName"`
	Email             string  `json:"email"`
	MobileNumber      string  `json:"mobileNumber"`
	IsInvitedUser     bool    `json:"isInvitedUser"`
	IsContractor      bool    `json:"isContractor"`
	IsActivated       bool    `json:"isActivated"`
	Position          string  `json:"position"`
	ImgUrl            string  `json:"imgUrl"`
	ActivatedAt       string  `json:"activatedAt"`
	DeactivatedAt     *string `json:"deactivatedAt"`
	NotificationOnOff bool    `json:"notificationOnOff"`
	Status            string  `json:"status"`
	OsType            *string `json:"osType"`
}

type ExpenseCardDetail struct {
	CardNumber         string            `json:"cardNumber"`
	CardUser           ExpenseUserDetail `json:"cardUser"`
	Alias              *string           `json:"alias"`
	LimitAmount        int               `json:"limitAmount"`
	UsedAmount         int               `json:"usedAmount"`
	RemainAmount       int               `json:"remainAmount"`
	CompanyCode        string            `json:"companyCode"`
	NamedYn            *string           `json:"namedYn"`
	CardName           string            `json:"cardName"`
	CardType           string            `json:"cardType"`
	UserNm             string            `json:"userNm"`
	UnmaskedCardNumber *string           `json:"unmaskedCardNumber"`
	DuplicationStatus  string            `json:"duplicationStatus"`
	FullCardNumber     string            `json:"fullCardNumber"`
	Invalid            bool              `json:"invalid"`
}

type ExpensePurposeDetail struct {
	Name           string   `json:"name"`
	Category       Category `json:"category"`
	LimitAmount    int      `json:"limitAmount"`
	ListOrder      int      `json:"listOrder"`
	IsActivated    bool     `json:"isActivated"`
	HasRequirement bool     `json:"hasRequirement"`
	LimitType      string   `json:"limitType"`
}

type ExpenseEvidence struct {
	EvidenceId int    `json:"evidenceId"`
	FileName   string `json:"fileName"`
	MimeType   string `json:"mimeType"`
	SignedUrl  string `json:"signedUrl"`
}

type ExpenseDeduction struct {
	IsExpenseDeductible bool `json:"isExpenseDeductible"`
	IsDeducted          bool `json:"isDeducted"`
}

type ExpenseDetail struct {
	ExpenseId                      int                   `json:"expenseId"`
	CardApprovalNumber             string                `json:"cardApprovalNumber"`
	ExpenseDate                    string                `json:"expenseDate"`
	ExpenseTime                    string                `json:"expenseTime"`
	Card                           ExpenseCardDetail     `json:"card"`
	User                           ExpenseUserDetail     `json:"user"`
	UseAmount                      float64               `json:"useAmount"`
	Currency                       string                `json:"currency"`
	KrwAmount                      int                   `json:"krwAmount"`
	ApprovalStatus                 string                `json:"approvalStatus"`
	ApprovedAmount                 *int                  `json:"approvedAmount"`
	ApprovedAt                     *string               `json:"approvedAt"`
	ApprovedBy                     *string               `json:"approvedBy"`
	Comments                       json.RawMessage       `json:"comments"`
	Purpose                        *ExpensePurposeDetail `json:"purpose"`
	Participants                   []ExpenseUserDetail   `json:"participants"`
	ExpenseExternalUsers           []ExpenseExternalUser `json:"expenseExternalUsers"`
	StoreName                      string                `json:"storeName"`
	StoreAddress                   string                `json:"storeAddress"`
	Memo                           *string               `json:"memo"`
	EvidenceList                   []ExpenseEvidence     `json:"evidenceList"`
	CompanyCode                    *string               `json:"companyCode"`
	CommentCount                   *int                  `json:"commentCount"`
	PurposeRequirementItem         *string               `json:"purposeRequirementItem"`
	PurposeRequirementItemType     *string               `json:"purposeRequirementItemType"`
	PurposeRequirementValue        *string               `json:"purposeRequirementValue"`
	IsPurposeRequirementInputValue bool                  `json:"isPurposeRequirementInputValue"`
	ExpenseDeductionResDto         ExpenseDeduction      `json:"expenseDeductionResDto"`
}

type GetExpenseResponse Response[ExpenseDetail]

type GetExpensesOptions struct {
	ApprovalState string
	Memo          string
	PurposeName   string
	UserName      string
	StartDate     string // yyyyMMdd or yyyy-MM-dd
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
		if opts.Page > 0 {
			q.Add("page", fmt.Sprintf("%d", opts.Page))
		}
		if opts.StartDate != "" {
			q.Add("startDate", opts.StartDate)
		}
		if opts.Size > 0 {
			q.Add("size", fmt.Sprintf("%d", opts.Size))
		}
	}
	req.URL.RawQuery = q.Encode()

	var resp GetExpensesResponse
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetExpense(expenseId string) (*GetExpenseResponse, error) {
	url := fmt.Sprintf("%s/v1/expenses/%s", c.BaseURL, expenseId)
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

// Request/Response Structs for Updates

type UpdatePurposeRequest struct {
	PurposeId                      int    `json:"purposeId"`
	PurposeRequirementItem         string `json:"purposeRequirementItem,omitempty"`
	PurposeRequirementItemType     string `json:"purposeRequirementItemType,omitempty"`
	PurposeRequirementValue        string `json:"purposeRequirementValue,omitempty"`
	IsPurposeRequirementInputValue bool   `json:"isPurposeRequirementInputValue"`
}

type UpdateExpensePurposeResponse = Response[Expense]

func (c *Client) UpdateExpensePurpose(expenseId string, req UpdatePurposeRequest) (*UpdateExpensePurposeResponse, error) {
	url := fmt.Sprintf("%s/v1/expenses/%s/purposes", c.BaseURL, expenseId)
	return c.doUpdateExpense(url, "PUT", req)
}

type UpdateExpensesPurposeItem struct {
	ExpenseId                      int    `json:"expenseId"`
	PurposeId                      int    `json:"purposeId"`
	PurposeRequirementItem         string `json:"purposeRequirementItem,omitempty"`
	PurposeRequirementItemType     string `json:"purposeRequirementItemType,omitempty"`
	PurposeRequirementValue        string `json:"purposeRequirementValue,omitempty"`
	IsPurposeRequirementInputValue bool   `json:"isPurposeRequirementInputValue"`
}

type UpdateExpensesPurposeRequest []UpdateExpensesPurposeItem

type UpdateExpensesPurposeResponse = Response[bool]

func (c *Client) UpdateExpensesPurpose(req UpdateExpensesPurposeRequest) (*UpdateExpensesPurposeResponse, error) {
	url := fmt.Sprintf("%s/v1/expenses/purposes", c.BaseURL)
	// Manual implementation for array request body logic if needed, but standard Do should work
	// However, doUpdateExpense is typed for Response[Expense]. We need generic.
	// Let's make a generic helper later or just implement inline.
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

type UpdateMemoRequest struct {
	Memo string `json:"memo"`
}

type UpdateExpenseMemoResponse = Response[Expense]

func (c *Client) UpdateExpenseMemo(expenseId string, memo string) (*UpdateExpenseMemoResponse, error) {
	url := fmt.Sprintf("%s/v1/expenses/%s/memo", c.BaseURL, expenseId)
	return c.doUpdateExpense(url, "PUT", UpdateMemoRequest{Memo: memo})
}

type UpdateParticipantsRequest struct {
	ExternalUsers  []map[string]string `json:"externalUsers"` // [{"company": "...", "name": "..."}]
	ParticipantIds []int               `json:"participantIds"`
}

type UpdateExpenseParticipantsResponse = Response[Expense]

func (c *Client) UpdateExpenseParticipants(expenseId string, req UpdateParticipantsRequest) (*UpdateExpenseParticipantsResponse, error) {
	url := fmt.Sprintf("%s/v1/expenses/%s/participants", c.BaseURL, expenseId)
	return c.doUpdateExpense(url, "PUT", req)
}

type UpdateExpenseRequest struct {
	ExternalUserList               []map[string]string `json:"externalUserList"`
	Memo                           string              `json:"memo"`
	Participants                   []int               `json:"participants"`
	PurposeId                      int                 `json:"purposeId"`
	PurposeRequirementItem         string              `json:"purposeRequirementItem,omitempty"`
	PurposeRequirementItemType     string              `json:"purposeRequirementItemType,omitempty"`
	PurposeRequirementValue        string              `json:"purposeRequirementValue,omitempty"`
	IsPurposeRequirementInputValue bool                `json:"isPurposeRequirementInputValue"`
}

type UpdateExpenseResponse = Response[Expense]

func (c *Client) UpdateExpense(expenseId string, req UpdateExpenseRequest) (*UpdateExpenseResponse, error) {
	url := fmt.Sprintf("%s/v1/expenses/%s", c.BaseURL, expenseId)
	return c.doUpdateExpense(url, "PUT", req)
}

type UpdateApprovalStatusRequest struct {
	ApprovalStatus string `json:"approvalStatus"` // APPROVED
	ApprovedAmount int    `json:"approvedAmount"`
	ApprovedAt     string `json:"approvedAt"`
}

type UpdateExpenseApprovalStatusResponse = Response[Expense]

func (c *Client) UpdateExpenseApprovalStatus(expenseId string, req UpdateApprovalStatusRequest) (*UpdateExpenseApprovalStatusResponse, error) {
	url := fmt.Sprintf("%s/v1/expenses/%s/approval-status", c.BaseURL, expenseId)
	return c.doUpdateExpense(url, "PUT", req)
}

type ApproveExpensesRequest []struct {
	ExpenseId      int    `json:"expenseId"`
	ApprovalStatus string `json:"approvalStatus"`
	ApprovedAmount int    `json:"approvedAmount"`
	ApprovedAt     string `json:"approvedAt"`
}

type ApproveExpensesResponse Response[struct {
	SucceedStatements []Expense `json:"succeedStatements"`
	FailedStatements  []Expense `json:"failedStatements"` // Simplified for now
}]

func (c *Client) ApproveExpenses(req ApproveExpensesRequest) (*ApproveExpensesResponse, error) {
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

type AddCommentRequest struct {
	Comment string `json:"comment"`
	// ExpenseId is removed in spec
}

type AddCommentResponse Response[struct {
	CommentId *int   `json:"commentId"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}]

func (c *Client) AddComment(expenseId string, comment string) (*AddCommentResponse, error) {
	url := fmt.Sprintf("%s/v1/expenses/%s/comments", c.BaseURL, expenseId)
	httpReq, err := c.newRequest("POST", url, AddCommentRequest{Comment: comment})
	if err != nil {
		return nil, err
	}
	var resp AddCommentResponse
	if err := c.Do(httpReq, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Helper to avoid repetition for simple Expense return types
func (c *Client) doUpdateExpense(url, method string, body interface{}) (*Response[Expense], error) {
	req, err := c.newRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	var resp Response[Expense]
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
