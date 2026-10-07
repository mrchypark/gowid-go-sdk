package client

type Result struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
}

type Response[T any] struct {
	Result     Result `json:"result"`
	TotalCount int64  `json:"totalCount"`
	Data       T      `json:"data"`
}

// PageOptions covers the shared page/size/sort query parameters.
type PageOptions struct {
	Page int
	Size int
	Sort string
}

type PageInfo struct {
	TotalElements    int  `json:"totalElements"`
	TotalPages       int  `json:"totalPages"`
	Last             bool `json:"last"`
	NumberOfElements int  `json:"numberOfElements"`
	Size             int  `json:"size"`
	Number           int  `json:"number"`
	First            bool `json:"first"`
	Empty            bool `json:"empty"`
}

type Sort struct {
	Empty    bool `json:"empty"`
	Sorted   bool `json:"sorted"`
	Unsorted bool `json:"unsorted"`
}

type Pageable struct {
	Sort       Sort `json:"sort"`
	Offset     int  `json:"offset"`
	PageNumber int  `json:"pageNumber"`
	PageSize   int  `json:"pageSize"`
	Paged      bool `json:"paged"`
	Unpaged    bool `json:"unpaged"`
}
