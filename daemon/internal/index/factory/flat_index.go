package factory

type flatIndex struct {
	MethodT MethodType
	IType   IndexType
	Dim     uint16
}

func NewFlatIndex(itype IndexType, method MethodType, dimension uint16) GenericIndex {
	return &flatIndex{
		IType:   itype,
		MethodT: method,
		Dim:     dimension,
	}
}

func (f *flatIndex) IndexType() IndexType {
	return f.IType
}

func (f *flatIndex) Parameters() map[string]int {
	return map[string]int{}
}

func (f *flatIndex) Method() MethodType {
	return f.MethodT
}

func (f *flatIndex) Dimension() uint16 {
	return f.Dim
}
