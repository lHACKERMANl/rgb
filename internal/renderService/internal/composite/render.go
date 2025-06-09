package composite

type RenderHandlerAdapter interface {
	Init() error
	Render()
	Cleanup()
}

type Render struct {
	Handler RenderHandlerAdapter
}

func NewRender() *Render {
	return &Render{
		handler.NewRender(),
	}
}
