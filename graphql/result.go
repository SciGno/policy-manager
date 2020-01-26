package graphql

// Results is the GraphQL response
type Results struct {
	// Connection `json:"connection"`

	TotalCount int    `json:"totalCount"`
	Edges      []Node `json:"edges"`
	PageInfo   `json:"pageInfo"`
	Metrics    `json:"metrics"`
}

// Connection struct
// type Connection struct {
// 	TotalCount int    `json:"totalCount"`
// 	Edges      []Node `json:"edges"`
// 	PageInfo   `json:"pageInfo"`
// 	Metrics    `json:"metrics"`
// }

// Node struct
type Node struct {
	Node   interface{} `json:"node"`
	Cursor int         `json:"cursor"`
}

// PageInfo struct
type PageInfo struct {
	EndCursor   int  `json:"endCursor"`
	HasNextPage bool `json:"hasNextPage"`
}

// Metrics struct
type Metrics struct {
	QueryTime   string `json:"queryTime"`
	RequestTime string `json:"requestTime"`
}
