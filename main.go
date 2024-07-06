package main

//// TODO: add fps limit

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

	// game := &demos.ResizeWindow{}
	// gi, err := engine.LoadGame(
	// 	game,
	// 	engine.GameConfig{
	// 		Width:                   800,
	// 		Height:                  600,
	// 		Title:                   "ResizeWindow",
	// 		StopUsingDefaultShaders: true,
	// 		Resizable:               true,
	// 		ResizeCallback: func() {
	// 			println("resized")
	// 		},
	// 	},
	// )
	// if err != nil {
	// 	return
	// }
	// game.Window = gi.Window
	// err = gi.Run()
	// if err != nil {
	// 	return
	// }

	// err = gi.Destroy()
	// if err != nil {
	// 	return
	// }

}

func orExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
