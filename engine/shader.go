package engine

import (
	"fmt"
	"strings"

	"github.com/ahmedsat/engine/math/vectors"
	"github.com/go-gl/gl/v4.5-core/gl"
)

var defaultVert = `
#version 330 core
layout (location = 0) in vec2 aPos;

void main()
{
    gl_Position = vec4(aPos.x, aPos.y, 0, 1);
}
`
var defaultFrag = `
#version 330 core
out vec4 FragColor;

void main()
{
    FragColor = vec4(1.0f, 1.0f, 1.0f, 1.0f);
}
`

type Shader uint32

func (sh Shader) Use() { gl.UseProgram(uint32(sh)) }

func (sh Shader) Delete() { gl.DeleteShader(uint32(sh)) }

func (sh Shader) UniformVec4float32(name string, vec vectors.Vec4f32) {
	vertexColorLocation := gl.GetUniformLocation(uint32(sh), gl.Str(name+"\x00"))
	gl.UseProgram(uint32(sh))
	gl.Uniform4f(vertexColorLocation, vec.X, vec.Y, vec.Z, vec.W)
}

func (sh Shader) UniformFloat32(name string, value float32) {
	vertexColorLocation := gl.GetUniformLocation(uint32(sh), gl.Str(name+"\x00"))
	gl.UseProgram(uint32(sh))
	gl.Uniform1f(vertexColorLocation, value)
}

func CreateShader(vertexSource, fragmentSource string) (sh Shader, err error) {

	vertexShader, err := compileShader(vertexSource+"\x00", gl.VERTEX_SHADER)
	if err != nil {
		return
	}

	fragmentShader, err := compileShader(fragmentSource+"\x00", gl.FRAGMENT_SHADER)
	if err != nil {
		return
	}

	sh = Shader(gl.CreateProgram())

	gl.AttachShader(uint32(sh), vertexShader)
	gl.AttachShader(uint32(sh), fragmentShader)

	gl.LinkProgram(uint32(sh))

	err = checkShaderLinkStatus(uint32(sh))
	if err != nil {
		return
	}
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return
}

func checkShaderLinkStatus(shaderProgram uint32) (err error) {
	var status int32
	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shaderProgram, logLength, nil, gl.Str(log))

		err = fmt.Errorf("failed to link program: %v", log)
	}
	return
}
func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func GetDefaultShader() (vert, frag string) { return defaultVert, defaultFrag }
