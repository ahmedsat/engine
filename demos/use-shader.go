package demos

import (
	"image/color"
	"os"

	"github.com/ahmedsat/engine/engine"
)

func init() {
	Demos = append(Demos, func() (err error) {
		gi, err := engine.LoadGame(
			&HelloShader{},
			engine.GameConfig{
				Width:                   800,
				Height:                  600,
				Title:                   "HelloShader",
				StopUsingDefaultShaders: true,
			},
		)
		if err != nil {
			return
		}
		err = gi.Run()
		if err != nil {
			return
		}

		err = gi.Destroy()
		if err != nil {
			return
		}
		return
	})
}

type HelloShader struct {
	engine.BaseGame
	triangle uint32
	shader   engine.Shader
}

func (h *HelloShader) Title() string { return "HelloShader" }

func (h *HelloShader) Init() (err error) {

	vertices := []float32{
		// first triangle
		0.98, 0.0, 0.98, // top right
		0.98, 0.0, -0.98, // bottom right
		-0.98, 0.0, 0.98, // top left
		// second triangle
		0.98, 0.0, -0.98, // bottom right
		-0.98, 0.0, -0.98, // bottom left
		-0.98, 0.0, 0.98, // top left
	}

	vertBytes, err := os.ReadFile("shaders/z-not-y.vert")
	if err != nil {
		return
	}

	fragBytes, err := os.ReadFile("shaders/mango.frag")
	if err != nil {
		return
	}

	h.shader, err = engine.CreateShader(string(vertBytes), string(fragBytes))
	if err != nil {
		return
	}

	h.triangle = engine.LoadVertices(vertices, engine.VertexAttribute{Index: 0, Size: 3, Stride: 3, Offset: 0})

	h.shader.Use()
	return
}

func (h *HelloShader) Render() (err error) {
	engine.ClearBackground(color.NRGBA{
		R: 51, G: 77, B: 77, A: 255,
	})
	engine.DrawVertices(h.triangle, 0, 6)
	return
}
