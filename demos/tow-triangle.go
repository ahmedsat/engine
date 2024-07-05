package demos

import (
	"image/color"

	"github.com/ahmedsat/engine/engine"
)

func init() {
	Demos = append(Demos, &TowTriangles{})
}

type TowTriangles struct {
	engine.BaseGame
	t1, t2 uint32
}

func (t *TowTriangles) Title() string { return "TowTriangles" }

func (t *TowTriangles) Init() (err error) {

	v1 := []float32{
		-0.8, 0.8,
		0.8, 0.8,
		0.0, 0.0,
	}

	v2 := []float32{
		0.0, 0.0,
		0.8, -0.8,
		-0.8, -0.8,
	}

	t.t1 = engine.LoadVertices(v1, engine.VertexAttribute{Index: 0, Size: 2, Stride: 2, Offset: 0})
	t.t2 = engine.LoadVertices(v2, engine.VertexAttribute{Index: 0, Size: 2, Stride: 2, Offset: 0})

	return
}

func (t *TowTriangles) Render() (err error) {
	engine.ClearBackground(color.NRGBA{
		R: 51, G: 77, B: 77, A: 255,
	})
	engine.DrawVertices(t.t1, 0, 3)
	engine.DrawVertices(t.t2, 0, 3)
	return
}
