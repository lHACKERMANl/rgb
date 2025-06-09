package glpackage

type ShapeAdapter interface {
	GetVertices() []float32
	GetIndices() []uint32
}

type Shape struct {
	Vertices []float32
	Indices  []uint32
}

func (s *Shape) GetVertices() []float32 { return s.Vertices }
func (s *Shape) GetIndices() []uint32   { return s.Indices }
