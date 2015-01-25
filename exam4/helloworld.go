package main

import (
	"fmt"
	"os"

	"gopkg.in/qml.v1"
)

type Person struct {
	Name string
}

func main() {
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	engine := qml.NewEngine()

	// ***** Go types
	qml.RegisterTypes("GoExtensions", 1, 0, []qml.TypeSpec{{
		Init: func(p *Person, obj qml.Object) { p.Name = "<none>" },
	}})
	controls, err := engine.LoadFile("helloworld.qml")
	if err != nil {
		return err
	}

	window := controls.CreateWindow(nil)

	window.Show()
	window.Wait()

	return nil
}
