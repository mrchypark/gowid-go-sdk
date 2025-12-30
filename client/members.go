package client

import (
	"fmt"
	"net/http"
)

type MemberRole struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"` // Can be null
}

type MemberDepartment struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Member struct {
	UserID            int              `json:"userId"`
	UserName          string           `json:"userName"`
	Email             string           `json:"email"`
	IsContractor      bool             `json:"isContractor"`
	Status            string           `json:"status"`
	Department        MemberDepartment `json:"department"`
	Position          string           `json:"position"`
	Role              MemberRole       `json:"role"`
	NotificationOnOff bool             `json:"notificationOnOff"`
}

type GetMembersResponse Response[[]Member]

type GetMembersOptions struct {
}

func (c *Client) GetMembers(opts *GetMembersOptions) (*GetMembersResponse, error) {
	url := fmt.Sprintf("%s/v1/members", c.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	var resp GetMembersResponse
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
