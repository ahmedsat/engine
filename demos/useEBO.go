package demos

import (
	"image/color"

	"github.com/ahmedsat/engine/engine"
)

func init() {
	Demos = append(Demos, func() (err error) {
		gi, err := engine.LoadGame(
			&HelloEBO{},
			engine.GameConfig{
				Width:                   800,
				Height:                  600,
				Title:                   "HelloEBO",
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

type HelloEBO struct {
	engine.BaseGame
	drawId uint32
}

func (h *HelloEBO) Title() string { return "HelloEBO" }

func (h *HelloEBO) Init() (err error) {

	vertices := []float32{
		0.5, 0.5, 0.0, // top right
		0.5, -0.5, 0.0, // bottom right
		-0.5, -0.5, 0.0, // bottom left
		-0.5, 0.5, 0.0, // top left
	}

	indices := []uint32{
		0, 1, 3, // first Triangle
		1, 2, 3, // second Triangle
	}

	h.drawId = engine.LoadVerticesWithIndices(
		vertices, indices,
		engine.VertexAttribute{Index: 0, Size: 3, Stride: 3, Offset: 0})
	engine.DrawLines()
	return
}

func (h *HelloEBO) Render() (err error) {
	engine.ClearBackground(color.NRGBA{
		R: 51, G: 77, B: 77, A: 255,
	})

	engine.DrawIndices(h.drawId, 0, 6)

	return
}
