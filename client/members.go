package client

import (
	"fmt"
	"net/http"
)

type MemberRole struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MemberDepartment struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Member struct {
	UserID            int64            `json:"userId"`
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

func (c *Client) GetMembers() (*GetMembersResponse, error) {
	url := fmt.Sprintf("%s/v1/members", c.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var resp GetMembersResponse
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
