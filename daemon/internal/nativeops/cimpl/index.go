package wrapper

import (
	"victord/daemon/internal/nativeops"
	"victord/daemon/platform/victor"
)

type index struct{}

func NewIndex() nativeops.IndexOps {
	return &index{}
}

func (i *index) AllocIndex(indexType, method int, dims uint16) (nativeops.VectorOps, error) {
	idx, err := victor.AllocIndex(indexType, method, dims)
	if err != nil {
		return nil, err
	}
	return &cindex{Index: idx}, nil
}
