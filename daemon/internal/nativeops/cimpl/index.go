package cimpl

import (
	"victord/daemon/internal/index/factory"
	"victord/daemon/internal/nativeops"
	"victord/daemon/platform/victor"
)

type IndexOpsImpl struct{}

type VIndex struct {
	Index *victor.Index
}

func NewIndex() *VIndex {
	return &VIndex{}
}

func (io *IndexOpsImpl) AllocIndex(indexOption factory.GenericIndex) (nativeops.VectorOps, error) { //laura
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

func NewIndexOps() nativeops.IndexOps {
	return &IndexOpsImpl{}
}
