package client

import (
	"fmt"
	"net/http"
	"strconv"
)

type Category struct {
	CategoryId int64  `json:"categoryId"`
	Name       string `json:"name"`
}

// PurposeRequirement is the V1 requirement item. Per the official spec,
// purposeId and options are writeOnly: they are accepted on input but the
// server does not return them on GET /v1/purposes.
type PurposeRequirement struct {
	Id               int64    `json:"id"`
	PurposeId        int64    `json:"purposeId,omitempty"`
	Type             string   `json:"type"`
	Item             string   `json:"item"`
	Options          []string `json:"options,omitempty"`
	GuideDesc        string   `json:"guideDesc"`
	IsAvailableInput bool     `json:"isAvailableInput"`
}

type Purpose struct {
	PurposeId      int64               `json:"purposeId"`
	Name           string              `json:"name"`
	Category       Category            `json:"category"`
	ListOrder      int                 `json:"listOrder"`
	LimitType      string              `json:"limitType"`
	LimitAmount    int64               `json:"limitAmount"`
	IsActivated    bool                `json:"isActivated"`
	HasRequirement bool                `json:"hasRequirement"`
	Requirement    *PurposeRequirement `json:"requirement"`
	IsDeducted     bool                `json:"isDeducted"`
}

type GetPurposesResponse Response[[]Purpose]

type GetPurposesOptions struct {
	IsActivated *bool
}

func (c *Client) GetPurposes(opts *GetPurposesOptions) (*GetPurposesResponse, error) {
	url := fmt.Sprintf("%s/v1/purposes", c.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if opts != nil && opts.IsActivated != nil {
		q.Add("isActivated", strconv.FormatBool(*opts.IsActivated))
	}
	req.URL.RawQuery = q.Encode()

	var resp GetPurposesResponse
	if err := c.Do(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

type PurposeRequirementContent struct {
	Content []string `json:"content"`
}

type GetPurposeRequirementsResponse Response[PurposeRequirementContent]

func (c *Client) GetPurposeRequirements(purposeId int64) (*GetPurposeRequirementsResponse, error) {
	if purposeId <= 0 {
		return nil, fmt.Errorf("invalid purposeId %d: must be a positive int64", purposeId)
	}
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
