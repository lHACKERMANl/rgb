package shaderService

import (
	"fmt"
	"github.com/go-gl/gl/v4.5-core/gl"
	"strings"
)

type ShaderService struct {
	program uint32
}

func New() *ShaderService {
	return &ShaderService{}
}

func (s *ShaderService) Init() error {
	vertexShader, err := s.compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return fmt.Errorf("vertex shader compilation failed: %v", err)
	}
	defer gl.DeleteShader(vertexShader)

	fragmentShader, err := s.compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return fmt.Errorf("fragment shader compilation failed: %v", err)
	}
	defer gl.DeleteShader(fragmentShader)

	s.program = gl.CreateProgram()
	gl.AttachShader(s.program, vertexShader)
	gl.AttachShader(s.program, fragmentShader)
	gl.LinkProgram(s.program)

	var status int32
	gl.GetProgramiv(s.program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(s.program, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(s.program, logLength, nil, gl.Str(log))
		return fmt.Errorf("program linking failed: %s", log)
	}

	fmt.Println("Шейдеры успешно скомпилированы и связаны")
	return nil
}

func (s *ShaderService) compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	defer free()

	gl.ShaderSource(shader, 1, csources, nil)
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
		return 0, fmt.Errorf("shader compilation failed: %s", log)
	}

	return shader, nil
}

func (s *ShaderService) Use() {
	gl.UseProgram(s.program)
}

func (s *ShaderService) Cleanup() {
	if s.program != 0 {
		gl.DeleteProgram(s.program)
	}
}
