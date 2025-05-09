package wrapper

import (
	"victord/daemon/platform/victor"
)

type VIndex struct {
	Index *victor.Index
}

func NewIndex() *VIndex {
	return &VIndex{}
}

func (i *VIndex) AllocIndex(indexType, method int, dims uint16) (*VIndex, error) {
	idx, err := victor.AllocIndex(indexType, method, dims)
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
