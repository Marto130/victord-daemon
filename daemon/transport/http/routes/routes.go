package routes

const (
	CreateIndexPath  = "/api/index/{indexName}"
	DestroyIndexPath = "/api/index/{indexName}"
	InsertVectorPath = "/api/vector/{indexName}"
	DeleteVectorPath = "/api/vector/{indexName}/{vectorID}"
	SearchVectorPath = "/api/vector/{indexName}/search"
)
