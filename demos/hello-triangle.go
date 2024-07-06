package demos

import (
	"image/color"

	"github.com/ahmedsat/engine/engine"
)

func init() {
	Demos = append(Demos, func() (err error) {
		gi, err := engine.LoadGame(
			&HelloTriangle{},
			engine.GameConfig{
				Width:                   800,
				Height:                  600,
				Title:                   "HelloTriangle",
				StopUsingDefaultShaders: false,
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

type HelloTriangle struct {
	engine.BaseGame
	triangle uint32
}

func (h *HelloTriangle) Init() (err error) {

	vertices := []float32{
		-0.9, -0.9,
		0.9, -0.9,
		0.0, 0.9,
	}

	h.triangle = engine.LoadVertices(vertices, engine.VertexAttribute{Index: 0, Size: 2, Stride: 2, Offset: 0})

	return
}

func (h *HelloTriangle) Render() (err error) {
	engine.ClearBackground(color.NRGBA{
		R: 51, G: 77, B: 77, A: 255,
	})
	engine.DrawVertices(h.triangle, 0, 3)
	return
}
