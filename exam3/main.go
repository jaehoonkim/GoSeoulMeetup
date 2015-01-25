package main

import (
	"fmt"
	"os"

	"gopkg.in/qml.v1"
)

const contents string = `
import QtQuick 2.0
Item {
	Component.onCompleted: console.log("Name is", person.name)
}
`

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

	// ***** publishing Go values to QML
	context := engine.Context() // root context
	context.SetVar("person", &Person{Name: "naeun"})

	controls, err := engine.LoadString("contents.qml", contents)
	if err != nil {
		return err
	}

	window := controls.CreateWindow(nil)

	window.Show()
	window.Wait()

	return nil
}
