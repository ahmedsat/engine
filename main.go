package main

// TODO: implement don't use default shader
// TODO: implement resizable screen
// TODO: 	pass screen resolution to shader via uniform
// TODO: 	enable custom resize callback

import (
	"fmt"
	"os"
	"runtime"

	"github.com/ahmedsat/engine/demos"
	"github.com/ahmedsat/engine/engine"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {

	for _, g := range demos.Demos {
		gi, err := engine.LoadGame(g, 800, 600)
		orExit(err)
		orExit(gi.Run())
		orExit(gi.Destroy())
	}

	// gi, err := engine.LoadGame(&demos.MultipleAttribute{}, 800, 600)
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
