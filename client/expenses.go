package client

import (
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

type GetExpensesResponse Response[[]Expense]

type GetExpensesOptions struct {
	Limit     int
	Page      int
	StartDate string // YYYY-MM-DD
	EndDate   string // YYYY-MM-DD
}

func (c *Client) GetExpenses(opts *GetExpensesOptions) (*GetExpensesResponse, error) {
	url := fmt.Sprintf("%s/v1/expenses", c.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if opts != nil {
		if opts.Limit > 0 {
			q.Add("limit", fmt.Sprintf("%d", opts.Limit))
		}
		if opts.Page > 0 {
			q.Add("page", fmt.Sprintf("%d", opts.Page))
		}
		if opts.StartDate != "" {
			q.Add("startDate", opts.StartDate)
		}
		if opts.EndDate != "" {
			q.Add("endDate", opts.EndDate)
		}
	}
	req.URL.RawQuery = q.Encode()

	var resp GetExpensesResponse
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

type UpdateExpensesPurposeRequest []struct {
	ExpenseId                      int    `json:"expenseId"`
	PurposeId                      int    `json:"purposeId"`
	PurposeRequirementItem         string `json:"purposeRequirementItem,omitempty"`
	PurposeRequirementItemType     string `json:"purposeRequirementItemType,omitempty"`
	PurposeRequirementValue        string `json:"purposeRequirementValue,omitempty"`
	IsPurposeRequirementInputValue bool   `json:"isPurposeRequirementInputValue"`
}

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
	Comment   string `json:"comment"`
	// ExpenseId is removed in spec
}

type AddCommentResponse Response[struct {
	CommentId int    `json:"commentId"`
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
