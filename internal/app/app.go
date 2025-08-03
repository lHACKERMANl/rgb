package app

import (
	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/lHACKERMANl/rgb/internal/renderService"
	"github.com/lHACKERMANl/rgb/internal/shaderService"
	"github.com/lHACKERMANl/rgb/internal/windowService"
	"github.com/lHACKERMANl/rgb/pkg/glPackage"
)

type App struct {
	window    *windowService.WindowService
	shader    *shaderService.ShaderService
	render    *renderService.RenderService
	isRunning bool
}

func NewApp() *App {
	return &App{
		window: windowService.New(),
		shader: shaderService.New(),
		render: renderService.New(),
	}
}

func (a *App) Run() error {

	if err := a.window.Init(800, 600, "2D Render App"); err != nil {
		return err
	}
	defer a.window.Cleanup()

	if err := gl.Init(); err != nil {
		return err
	}

	if err := a.shader.Init(); err != nil {
		return err
	}
	defer a.shader.Cleanup()

	if err := a.render.Init(); err != nil {
		return err
	}
	defer a.render.Cleanup()

	triangle := &glPackage.Shape{
		Vertices: []float32{

			-0.5, -0.5,
			0.5, -0.5,
			0.0, 0.5,
		},
		Indices: []uint32{0, 1, 2},
	}

	if err := a.render.PrepareShape(triangle); err != nil {
		return err
	}

	a.isRunning = true

	for a.isRunning && !a.window.ShouldClose() {
		a.update()
		a.render_frame()
		a.window.SwapBuffers()
		glfw.PollEvents()
	}

	return nil
}

func (a *App) update() {

	if a.window.IsKeyPressed(glfw.KeyEscape) {
		a.isRunning = false
	}
}

func (a *App) render_frame() {

	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.ClearColor(0.2, 0.3, 0.3, 1.0)

	a.shader.Use()

	a.render.Draw()
}
