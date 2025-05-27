package cimpl

import (
	"victord/daemon/platform/types"
)

func (c *NativeIndex) Delete(id uint64) error {
	return c.CIndex.Delete(id)
}

func (c *NativeIndex) Insert(id uint64, vector []float32) error {
	return c.CIndex.Insert(id, vector)
}

func (c *NativeIndex) Search(vector []float32, dim int) (*types.MatchResult, error) {
	return c.CIndex.Search(vector, dim)
}
