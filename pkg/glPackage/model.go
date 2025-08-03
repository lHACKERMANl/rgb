package glPackage

type ShapeAdapter interface {
	GetVertices() []float32
	GetIndices() []uint32
}

type Shape struct {
	Vertices []float32 `json:"vertices"`
	Indices  []uint32  `json:"indices"`
}

func (s *Shape) GetVertices() []float32 { return s.Vertices }
func (s *Shape) GetIndices() []uint32   { return s.Indices }
