package client

import (
	"fmt"
	"net/http"
)

type Category struct {
	CategoryId int    `json:"categoryId"`
	Name       string `json:"name"`
}

type Purpose struct {
	PurposeId      int         `json:"purposeId"`
	Name           string      `json:"name"`
	Category       Category    `json:"category"`
	ListOrder      int         `json:"listOrder"`
	LimitType      string      `json:"limitType"`
	LimitAmount    int         `json:"limitAmount"`
	IsActivated    bool        `json:"isActivated"`
	HasRequirement bool        `json:"hasRequirement"`
	Requirement    interface{} `json:"requirement"` // Using interface{} as it can be null or complex object
	IsDeducted     bool        `json:"isDeducted"`
}

type GetPurposesResponse Response[[]Purpose]

type GetPurposesOptions struct {
	IsActivated *bool
	Limit       int
	Page        int
}

func (c *Client) GetPurposes(opts *GetPurposesOptions) (*GetPurposesResponse, error) {
	url := fmt.Sprintf("%s/v1/purposes", c.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if opts != nil {
		if opts.IsActivated != nil {
			q.Add("isActivated", fmt.Sprintf("%t", *opts.IsActivated))
		}
		if opts.Limit > 0 {
			q.Add("limit", fmt.Sprintf("%d", opts.Limit))
		}
		if opts.Page > 0 {
			q.Add("page", fmt.Sprintf("%d", opts.Page))
		}
	}
	req.URL.RawQuery = q.Encode()

	var resp GetPurposesResponse
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

type GetPurposeRequirementsResponse Response[[]string]

func (c *Client) GetPurposeRequirements(purposeId int) (*GetPurposeRequirementsResponse, error) {
	url := fmt.Sprintf("%s/v1/purposes/%d/requirements", c.BaseURL, purposeId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var resp GetPurposeRequirementsResponse
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
