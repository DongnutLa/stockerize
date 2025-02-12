package shared_domain

type PagingParams struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

type SearchQueryParams struct {
	PagingParams
	Search string `json:"search"`
}

type Metadata struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
	Count    int64 `json:"count"`
	HasNext  bool  `json:"hasNext"`
}

type PagingResponse[T any] struct {
	Metadata Metadata `json:"metadata"`
	Data     []T      `json:"data"`
}
