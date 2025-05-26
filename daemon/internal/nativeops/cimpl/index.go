package cimpl

import (
	"victord/daemon/internal/index/factory"
	"victord/daemon/internal/nativeops"
	"victord/daemon/platform/victor"
)

type IndexConstructor struct{}

type NativeIndex struct {
	CIndex *victor.Index
}

func NewIndex() *NativeIndex {
	return &NativeIndex{}
}

func (io *IndexConstructor) AllocIndex(indexOption factory.GenericIndex) (nativeops.VectorOps, error) { //laura
	idx, err := victor.AllocIndex(int(indexOption.IndexType()), int(indexOption.Method()),
		indexOption.Dimension(), indexOption.Parameters())
	if err != nil {
		return nil, err
	}
	return &NativeIndex{CIndex: idx}, nil
}

func (i *NativeIndex) DestroyIndex() {
	if i.CIndex != nil {
		i.CIndex.DestroyIndex()
	}
}

func NewIndexConstructor() nativeops.IndexOps {
	return &IndexConstructor{}
}
