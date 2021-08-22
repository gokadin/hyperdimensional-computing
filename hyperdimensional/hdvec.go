package hyperdimensional

type HdVec interface {
	Rand() *HdVec
	Bind(a, b *HdVec) *HdVec
	Unbind(a, b *HdVec) *HdVec
	Bundle(vectors ...*HdVec) *HdVec
	Similarity(a, b *HdVec) float64
	Equals(a, b *HdVec) bool
	Size() int
	MarshalJson() ([]byte, error)
	UnmarshalJSON(b []byte) error
}
