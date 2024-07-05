package engine

import (
	"github.com/ahmedsat/engine/math/vectors"
	"github.com/go-gl/gl/v4.5-core/gl"
)

func UniformVec4float32(sh Shader, name string, vec vectors.Vec4f32) {
	vertexColorLocation := gl.GetUniformLocation(uint32(sh), gl.Str(name+"\x00"))
	gl.UseProgram(uint32(sh))
	gl.Uniform4f(vertexColorLocation, vec.X, vec.Y, vec.Z, vec.W)
}

func UniformFloat32(sh Shader, name string, value float32) {
	vertexColorLocation := gl.GetUniformLocation(uint32(sh), gl.Str(name+"\x00"))
	gl.UseProgram(uint32(sh))
	gl.Uniform1f(vertexColorLocation, value)

}
