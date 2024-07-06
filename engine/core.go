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

	StopDefaultShaders bool
	defaultShaders     Shader
	gameConfig         GameConfig
}

type GameConfig struct {
	Width                   int
	Height                  int
	Title                   string
	StopUsingDefaultShaders bool
}

func LoadGame(g Game, gameConfig GameConfig) (gi gameInstance, err error) {

	gi.Game = g
	gi.gameConfig = gameConfig

	window, err := glfw.CreateWindow(gameConfig.Width, gameConfig.Height, gameConfig.Title, nil, nil)
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

	gl.Viewport(0, 0, int32(gameConfig.Width), int32(gameConfig.Height))
	window.SetFramebufferSizeCallback(func(w *glfw.Window, width, height int) {
		gi.gameConfig.Width = width
		gi.gameConfig.Height = height
		gl.Viewport(0, 0, int32(gameConfig.Width), int32(gameConfig.Height))
	})
	if !gameConfig.StopUsingDefaultShaders {
		var shader Shader
		shader, err = CreateShader(defaultVert, defaultFrag)
		if err != nil {
			err = errors.Join(errors.New("can not load default shader"), err)
			return
		}
		gi.defaultShaders = shader
	}

	err = gi.Init()

	return
}

func (gi gameInstance) Destroy() (err error) {
	if !gi.gameConfig.StopUsingDefaultShaders {
		gi.defaultShaders.Delete()
	}
	if gi.Window != nil {
		gi.Window.Destroy()
	}
	return
}

func (gi *gameInstance) Run() (err error) {
	if !gi.gameConfig.StopUsingDefaultShaders {
		gi.defaultShaders.Use()
	}
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
