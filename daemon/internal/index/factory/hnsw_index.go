package factory

const (
	efSearchKey    = "ef_search"
	efConstructKey = "ef_construct"
	m0Key          = "M0"
)

type hnswIndex struct {
	MethodT     MethodType
	IType       IndexType
	Dim         uint16
	EfSearch    int
	EfConstruct int
	M0          int
}

func NewHnswIndex(itype IndexType, method MethodType, dimension uint16, options map[string]int) GenericIndex {
	return &hnswIndex{
		IType:       itype,
		MethodT:     method,
		Dim:         dimension,
		EfSearch:    options[efSearchKey],
		EfConstruct: options[efConstructKey],
		M0:          options[m0Key],
	}
}

func (h *hnswIndex) IndexType() IndexType {
	return h.IType
}

func (h *hnswIndex) Parameters() map[string]int {
	return map[string]int{
		efSearchKey:    h.EfSearch,
		efConstructKey: h.EfConstruct,
		m0Key:          h.M0,
	}
}

func (h *hnswIndex) Method() MethodType {
	return h.MethodT
}

func (h *hnswIndex) Dimension() uint16 {
	return h.Dim
}
