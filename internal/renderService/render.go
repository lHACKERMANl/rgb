package renderService

import (
	"github.com/lHACKERMANl/rgb/pkg/glPackage"
)

type RenderService struct {
	glRender *glPackage.GlRender
	shape    glPackage.ShapeAdapter
}

func New() *RenderService {
	return &RenderService{
		glRender: glPackage.NewGlRender(),
	}
}

func (r *RenderService) Init() error {
	return nil
}

func (r *RenderService) PrepareShape(shape glPackage.ShapeAdapter) error {
	r.shape = shape
	return r.glRender.Init(shape)
}

func (r *RenderService) Draw() {
	if r.shape != nil {
		r.glRender.Render(r.shape)
	}
}

func (r *RenderService) Cleanup() {
	if r.glRender != nil {
		r.glRender.Cleanup()
	}
}
