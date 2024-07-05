package demos

import (
	"image/color"
	"os"

	"github.com/ahmedsat/engine/engine"
)

func init() {
	Demos = append(Demos, &HelloUniform{})
}

type HelloUniform struct {
	engine.BaseGame
	triangle uint32
	sh       engine.Shader
}

func (h *HelloUniform) Title() string { return "HelloUniform" }

func (h *HelloUniform) Init() (err error) {

	vertices := []float32{
		// first triangle
		0.98, 0.98, // top right
		0.98, -0.98, // bottom right
		-0.98, 0.98, // top left
		// second triangle
		0.98, -0.98, // bottom right
		-0.98, -0.98, // bottom left
		-0.98, 0.98, // top left
	}

	vert, _ := engine.GetDefaultShader()
	frag, err := os.ReadFile("shaders/time-uniform.frag")
	if err != nil {
		return
	}

	h.sh, err = engine.CreateShader(vert, string(frag))
	if err != nil {
		return
	}

	h.triangle = engine.LoadVertices(vertices, engine.VertexAttribute{Index: 0, Size: 2, Stride: 2, Offset: 0})

	return
}

func (h *HelloUniform) Render() (err error) {
	engine.ClearBackground(color.NRGBA{
		R: 51, G: 77, B: 77, A: 255,
	})
	h.sh.Use()
	h.sh.UniformFloat32("uTime", float32(engine.GetTime()))

	engine.DrawVertices(h.triangle, 0, 6)
	return
}
