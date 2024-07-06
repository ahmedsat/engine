package main

// TODO: implement don't use default shader
// TODO: implement resizable screen
// TODO: 	pass screen resolution to shader via uniform
// TODO: 	enable custom resize callback

// FIXME: cant move shader out of loop in MultipleAttribute

import (
	"fmt"
	"os"
	"runtime"

	"github.com/ahmedsat/engine/demos"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {

	for _, g := range demos.Demos {
		orExit(g())
	}

	// gi, err := engine.LoadGame(&demos.MultipleAttribute{}, engine.GameConfig{
	// 	Width:                   800,
	// 	Height:                  600,
	// 	Title:                   "MultipleAttribute",
	// 	StopUsingDefaultShaders: true,
	// })
	// orExit(err)
	// orExit(gi.Run())
	// orExit(gi.Destroy())

}

func orExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
