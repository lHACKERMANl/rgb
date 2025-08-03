package glPackage

import "errors"

var (
	ErrInvalidIndicesAndVertices = errors.New("indices and vertices must not be empty")
	ErrInvalidVertexCount        = errors.New("vertices must be in x,y pairs")
)
