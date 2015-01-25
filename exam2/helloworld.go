package main

import (
	"fmt"
	"os"

	"gopkg.in/qml.v1"
)

func main() {
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	engine := qml.NewEngine()

	controls, err := engine.LoadFile("helloworld.qml")
	if err != nil {
		return err
	}

	window := controls.CreateWindow(nil)

	/*
		window.On("visibleChanged", func(visible bool) {
			if visible {
				fmt.Println("Width", window.Int("width"))
			}
		})
	*/
	window.On("visibleChanged", onVisibleChanged)

	window.Show()
	window.Wait()

	return nil
}

func onVisibleChanged(visible bool) {
	if visible {
		fmt.Println("visible changed")
	}
}
