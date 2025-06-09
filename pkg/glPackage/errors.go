package glpackage

import "errors"

var (
	errInvalidIndicesAndVertices = errors.New("indices and vertices must not be empty")
	errInvalidVertexCount        = errors.New("vertices must be in x,y pairs")
)
