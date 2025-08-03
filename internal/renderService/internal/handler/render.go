package handler

import (
	validator "github.com/go-playground/validator/v10"
	"github.com/lHACKERMANl/rgb/renderService/internal/dto"
	dto "github.com/lHACKERMANl/rgb/renderService/internal/dto/usecase"
)

type RenderUsecaseAdpter interface {
	Init(dto.RenderInitInputUsecase) error
	Render()
	Cleanup()
}

type Render struct {
	usecase   *usecase.RenderUsecaseAdpter
	validator *validator.Validate
}

func NewRender(u *usecase.RenderUsecaseAdpter, v *validator.Validate) *Render {
	return &Render{usecase: u, validator: v}
}

func (h *Render) Init(input dto.RenderInitInputUsecase) error {
	if err := h.validator.Validate(input); err != nil {
		return ErrInvalidInput
	}

	return h.usecase.Run(dto.RenderInitInputUsecase)
}

func (h *Render) Render() {
}

func (h *Render) Cleanup() {
}
