package engine

import (
	"errors"

	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/go-gl/gl/v4.5-core/gl"
)

func init() {
	err := glfw.Init()
	if err != nil {
		panic(errors.New("can not initialize GLFW"))
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 5)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
}

type gameInstance struct {
	*glfw.Window
	Game

	defaultShaders Shader
}

func LoadGame(g Game, width int, height int) (gi gameInstance, err error) {

	gi.Game = g

	window, err := glfw.CreateWindow(width, height, g.Title(), nil, nil)
	if err != nil {
		err = errors.Join(errors.New("can not create new window"), err)
		return
	}
	gi.Window = window

	gi.MakeContextCurrent()

	if err = gl.Init(); err != nil {
		err = errors.Join(errors.New("can not initialize open gl"), err)
		return
	}

	gl.Viewport(0, 0, int32(width), int32(height))
	window.SetFramebufferSizeCallback(func(w *glfw.Window, width, height int) {
		gl.Viewport(0, 0, int32(width), int32(height))
	})

	shader, err := CreateShader(defaultVert, defaultFrag)
	if err != nil {
		err = errors.Join(errors.New("can not load default shader"), err)
		return
	}
	gi.defaultShaders = shader

	err = gi.Init()

	return
}

func (gi gameInstance) Destroy() (err error) {

	gi.defaultShaders.Delete()

	if gi.Window != nil {
		gi.Window.Destroy()
	}
	return
}

func (gi *gameInstance) Run() (err error) {

	gi.defaultShaders.Use()

	for !gi.ShouldClose() {

		if gi.GetKey(glfw.KeyEscape) == glfw.Press {
			gi.SetShouldClose(true)
		}

		if err = gi.Render(); err != nil {
			return
		}

		gi.SwapBuffers()
		glfw.PollEvents()
	}

	return
}
