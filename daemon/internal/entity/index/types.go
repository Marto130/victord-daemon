package index

type CreateIndexResult struct {
	IndexName string `json:"index_name"`
	ID        string `json:"id"`
	Dims      uint16 `json:"dims"`
	IndexType int    `json:"index_type"`
	Method    int    `json:"method"`
}

type DestroyIndexResult struct {
	ID        string `json:"id"`
	IndexName string `json:"index_name"`
}
