package composite

import (
	dto "github.com/lHACKERMANl/rgb/renderService/internal/dto/handler"
)

type RenderHandlerAdapter interface {
	Init(dto.RenderInitInputHandler) error
	Render()
	Cleanup()
}
