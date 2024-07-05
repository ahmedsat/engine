package engine

import (
	"image/color"
	"unsafe"

	"github.com/go-gl/gl/v4.5-core/gl"
)

func ClearBackground(c color.Color) {
	r, g, b, a := c.RGBA()

	gl.ClearColor(
		float32(r)/float32(0xffff),
		float32(g)/float32(0xffff),
		float32(b)/float32(0xffff),
		float32(a)/float32(0xffff))
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func DrawVertices(vao uint32, first int32, vertexCount int32) {
	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, first, vertexCount)
	gl.BindVertexArray(0)
}

func DrawIndices(vao uint32, first int32, vertexCount int32) {
	gl.BindVertexArray(vao)
	gl.DrawElements(gl.TRIANGLES, vertexCount, gl.UNSIGNED_INT, unsafe.Pointer(nil))
	gl.BindVertexArray(0)
}

func DrawLines() {
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
}

func DrawFill() {
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
}
