package windowService

import (
	"fmt"
	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"runtime"
)

type WindowService struct {
	window *glfw.Window
	width  int
	height int
}

func New() *WindowService {
	return &WindowService{}
}

func (w *WindowService) Init(width, height int, title string) error {
	w.width = width
	w.height = height

	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		return fmt.Errorf("failed to initialize GLFW: %v", err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.False)

	var err error
	w.window, err = glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		glfw.Terminate()
		return fmt.Errorf("failed to create window: %v", err)
	}

	w.window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		w.window.Destroy()
		glfw.Terminate()
		return fmt.Errorf("failed to initialize OpenGL: %v", err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Printf("OpenGL version: %s\n", version)

	glfw.SwapInterval(1)

	gl.Viewport(0, 0, int32(width), int32(height))

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	return nil
}

func (w *WindowService) ShouldClose() bool {
	return w.window.ShouldClose()
}

func (w *WindowService) SwapBuffers() {
	w.window.SwapBuffers()
}

func (w *WindowService) IsKeyPressed(key glfw.Key) bool {
	return w.window.GetKey(key) == glfw.Press
}

func (w *WindowService) Cleanup() {
	if w.window != nil {
		w.window.Destroy()
	}
	glfw.Terminate()
}
