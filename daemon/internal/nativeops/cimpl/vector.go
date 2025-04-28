package wrapper

import (
	"victord/daemon/platform/types"
	"victord/daemon/platform/victor"
)

type cindex struct {
	Index *victor.Index
}

func (c *cindex) Delete(id uint64) error {
	return c.Index.Delete(id)
}

func (c *cindex) Insert(id uint64, vector []float32) error {
	return c.Index.Insert(id, vector)
}

func (c *cindex) Search(vector []float32, dim int) (*types.MatchResult, error) {
	return c.Index.Search(vector, dim)
}
