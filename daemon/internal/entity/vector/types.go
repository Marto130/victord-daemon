package vector

type InsertVectorResult struct {
	ID     uint64    `json:"id"`
	Vector []float32 `json:"vector"`
}

type DeleteVectorResult struct {
	ID uint64 `json:"id"`
}

type SearchVectorResult struct {
	ID       int     `json:"id"`
	Distance float32 `json:"distance"`
}
