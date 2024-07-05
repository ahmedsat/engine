package engine

import (
	"github.com/go-gl/gl/v4.5-core/gl"
)

type VertexAttribute struct {
	Index, Size, Stride, Offset int32
}

func LoadVertices(vertices []float32, attributes ...VertexAttribute) (VAO uint32) {

	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	createVBO(vertices, attributes...)

	gl.BindVertexArray(0)

	return
}

func LoadVerticesWithIndices(vertices []float32, indices []uint32, attributes ...VertexAttribute) (VAO uint32) {

	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	createVBO(vertices, attributes...)

	createEBO(indices)

	gl.BindVertexArray(0)

	return

}

func createVBO(vertices []float32, attributes ...VertexAttribute) (VBO uint32) {
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)

	// ? GL_STREAM_DRAW: the data is set only once and used by the GPU at most a few times.
	// ? GL_STATIC_DRAW: the data is set only once and used many times.
	// ? GL_DYNAMIC_DRAW: the data is changed a lot and used many times.
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	for _, attribute := range attributes {
		gl.VertexAttribPointerWithOffset(uint32(attribute.Index), attribute.Size, gl.FLOAT, false, attribute.Stride*4, uintptr(attribute.Offset))
		gl.EnableVertexAttribArray(0)
	}

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	return VBO
}

func createEBO(indices []uint32) (EBO uint32) {
	gl.GenBuffers(1, &EBO)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO)

	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	// gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)s
	return
}
