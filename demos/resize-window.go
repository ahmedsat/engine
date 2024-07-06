package demos

import (
	"image/color"
	"os"

	"github.com/ahmedsat/engine/engine"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	Demos = append(Demos, func() (err error) {
		game := &ResizeWindow{}
		gi, err := engine.LoadGame(
			game,
			engine.GameConfig{
				Width:                   800,
				Height:                  600,
				Title:                   "ResizeWindow",
				StopUsingDefaultShaders: true,
				Resizable:               true,
			},
		)
		if err != nil {
			return
		}
		game.Window = gi.Window
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

type ResizeWindow struct {
	engine.BaseGame
	triangle uint32
	sh       engine.Shader
	Window   *glfw.Window
}

func (h *ResizeWindow) Title() string { return "ResizeWindow" }

func (h *ResizeWindow) Init() (err error) {

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
	frag, err := os.ReadFile("shaders/resize-window.frag")
	if err != nil {
		return
	}

	h.sh, err = engine.CreateShader(vert, string(frag))
	if err != nil {
		return
	}

	h.triangle = engine.LoadVertices(vertices, engine.VertexAttribute{Index: 0, Size: 2, Stride: 2, Offset: 0})

	h.sh.Use()
	return
}

func (h *ResizeWindow) Render() (err error) {
	engine.ClearBackground(color.NRGBA{
		R: 51, G: 77, B: 77, A: 255,
	})
	h.sh.Uniform1f("uTime", float32(engine.GetTime()))

	engine.DrawVertices(h.triangle, 0, 6)
	h.sh.ScreenResolutionUniforms(h.Window)
	return
}
