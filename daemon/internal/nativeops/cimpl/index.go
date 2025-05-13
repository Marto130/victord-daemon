package cimpl

import (
	"victord/daemon/internal/index/factory"
	"victord/daemon/platform/victor"
)

type VIndex struct {
	Index *victor.Index
}

func NewIndex() *VIndex {
	return &VIndex{}
}

func AllocIndex(indexOption factory.GenericIndex) (*VIndex, error) { //laura
	idx, err := victor.AllocIndex(int(indexOption.IndexType()), int(indexOption.Method()),
		indexOption.Dimension(), indexOption.Parameters())
	if err != nil {
		return nil, err
	}
	return &VIndex{Index: idx}, nil
}

func (i *VIndex) DestroyIndex() {
	if i.Index != nil {
		i.Index.DestroyIndex()
	}
}
