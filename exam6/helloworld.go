package main

import (
	"fmt"
	"os"

	"gopkg.in/qml.v1"
	"gopkg.in/qml.v1/gl/2.0"
)

type GoRect struct {
	qml.Object
}

func (r *GoRect) Paint(p *qml.Painter) {
	gl := GL.API(p)

	width := float32(r.Int("width"))
	height := float32(r.Int("height"))

	gl.Enable(GL.BLEND)
	gl.BlendFunc(GL.SRC_ALPHA, GL.ONE_MINUS_SRC_ALPHA)
	gl.Color4f(1.0, 0.0, 1.0, 0.8)
	gl.Begin(GL.QUADS)
	gl.Vertex2f(0, 0)
	gl.Vertex2f(width, 0)
	gl.Vertex2f(width, height)
	gl.Vertex2f(0, height)
	gl.End()
}

func main() {
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	engine := qml.NewEngine()

	qml.RegisterTypes("GoExtensions", 1, 0, []qml.TypeSpec{{
		Init: func(r *GoRect, obj qml.Object) { r.Object = obj },
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
