package demos

import (
	"image/color"
	"os"

	"github.com/ahmedsat/engine/engine"
)

func init() {

	Demos = append(Demos, func() (err error) {
		gi, err := engine.LoadGame(
			&MultipleAttribute{},
			engine.GameConfig{
				Width:                   800,
				Height:                  600,
				Title:                   "MultipleAttribute",
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

type MultipleAttribute struct {
	engine.BaseGame

	drawId uint32
	sh     engine.Shader
}

func (h *MultipleAttribute) Title() string { return "MultipleAttribute" }

func (h *MultipleAttribute) Init() (err error) {

	vertices := []float32{
		// positions     // colors
		0.5, -0.5 /* */, 1.0, 0.0, 0.0, // bottom right
		-0.5, -0.5 /**/, 0.0, 1.0, 0.0, // bottom left
		0.0, 0.5 /*  */, 0.0, 0.0, 1.0, // top
	}

	vert, err := os.ReadFile("shaders/multiple-attribute.vert")
	if err != nil {
		return
	}

	frag, err := os.ReadFile("shaders/multiple-attribute.frag")
	if err != nil {
		return
	}

	h.sh, err = engine.CreateShader(string(vert), string(frag))
	if err != nil {
		return
	}

	h.drawId = engine.LoadVertices(
		vertices,
		engine.VertexAttribute{Index: 0, Size: 2, Stride: 5, Offset: 0},
		engine.VertexAttribute{Index: 1, Size: 3, Stride: 5, Offset: 2})

	h.sh.UniformFloat32("uOffset", 0.5)
	h.sh.Use()
	return
}

func (h *MultipleAttribute) Render() (err error) {
	engine.ClearBackground(color.NRGBA{
		R: 51, G: 77, B: 77, A: 255,
	})

	engine.DrawVertices(h.drawId, 0, 3)
	return
}
