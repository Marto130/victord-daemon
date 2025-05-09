package wrapper

import (
	"victord/daemon/platform/types"
)

func (c *VIndex) Delete(id uint64) error {
	return c.Index.Delete(id)
}

func (c *VIndex) Insert(id uint64, vector []float32) error {
	return c.Index.Insert(id, vector)
}

func (c *VIndex) Search(vector []float32, dim int) (*types.MatchResult, error) {
	return c.Index.Search(vector, dim)
}
