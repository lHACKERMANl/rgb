package handler

type RenderUsecaseAdpter interface {
	Init() error
	Render()
	Cleanup()
}

type Render struct{}

func NewRender() *Render {
	return &Render{}
}

func (r *Render) Init() error {
	return nil
}

func (r *Render) Render() {
}

func (r *Render) Cleanup() {
}
