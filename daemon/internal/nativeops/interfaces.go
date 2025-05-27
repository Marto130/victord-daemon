package nativeops

import (
	"victord/daemon/internal/index/factory"
	"victord/daemon/platform/types"
)

// VectorOps defines the interface for managing vector data within an index.
//
// This allows vector operations to be performed without exposing the underlying
// C bindings located in an external library.
type VectorOps interface {
	Delete(uint64) error
	Insert(uint64, []float32) error
	Search([]float32, int) (*types.MatchResult, error)
	DestroyIndex()
}

type IndexOps interface {
	AllocIndex(factory.GenericIndex) (VectorOps, error)
}
