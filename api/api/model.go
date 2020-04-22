package api

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func DefaultQuery() Query {
	return Query{
		Pagination: Pagination{
			Offset: 0,
			Limit:  0,
		},
	}
}

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Order struct {
	OrderBy    string `json:"orderBy"`
	Descending bool   `json:"descending"`
}

type Filter struct {
	Condition string        `json:"condition"`
	Args      []interface{} `json:"args"`
}

type Query struct {
	Pagination
	Order
	Filter
}
