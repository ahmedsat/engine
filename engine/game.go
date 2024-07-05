package engine

import "github.com/go-gl/glfw/v3.3/glfw"

type Game interface {
	Init() error
	HandelInput(w *glfw.Window) error
	Render() error
	Title() string
}

type BaseGame struct {
}

func (bg *BaseGame) Init() error                      { return nil }
func (bg *BaseGame) HandelInput(w *glfw.Window) error { return nil }
func (bg *BaseGame) Render() error                    { return nil }
func (bg *BaseGame) Title() string                    { return "" }
