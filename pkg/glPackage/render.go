package glPackage

import (
	"github.com/go-gl/gl/v4.5-core/gl"
	"unsafe"
)

type GlRender struct {
	vao uint32
	vbo uint32
	ebo uint32
}

func NewGlRender() *GlRender {
	return &GlRender{}
}

func (r *GlRender) Init(shape ShapeAdapter) error {
	vertices := shape.GetVertices()
	indices := shape.GetIndices()

	if len(vertices) == 0 || len(indices) == 0 {
		return ErrInvalidIndicesAndVertices
	}
	if len(vertices)%2 != 0 {
		return ErrInvalidVertexCount
	}

	gl.GenVertexArrays(1, &r.vao)
	gl.GenBuffers(1, &r.vbo)
	gl.GenBuffers(1, &r.ebo)

	gl.BindVertexArray(r.vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, r.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*int(unsafe.Sizeof(float32(0))), gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, r.ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*int(unsafe.Sizeof(uint32(0))), gl.Ptr(indices), gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(0, 2, gl.FLOAT, false, 2*int32(unsafe.Sizeof(float32(0))), 0)

	gl.BindVertexArray(0)

	return nil
}

func (r *GlRender) Render(shape ShapeAdapter) {
	gl.BindVertexArray(r.vao)
	gl.DrawElements(gl.TRIANGLES, int32(len(shape.GetIndices())), gl.UNSIGNED_INT, nil)
	gl.BindVertexArray(0)
}

func (r *GlRender) Cleanup() {
	if r.vao != 0 {
		gl.DeleteVertexArrays(1, &r.vao)
	}
	if r.vbo != 0 {
		gl.DeleteBuffers(1, &r.vbo)
	}
	if r.ebo != 0 {
		gl.DeleteBuffers(1, &r.ebo)
	}
}
