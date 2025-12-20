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
