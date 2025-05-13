package factory

type IndexType int

const (
	FlatIndexType = IndexType(iota)
	HNSWIndexType
	NSWIndexType
)

type MethodType int

const (
	L2NORM = MethodType(iota)
	COSINE
)

type GenericIndex interface {
	IndexType() IndexType
	Method() MethodType
	Parameters() map[string]int
	Dimension() uint16
}
