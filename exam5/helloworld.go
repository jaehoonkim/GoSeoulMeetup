package main

import (
	"fmt"
	"os"

	"gopkg.in/qml.v1"
)

type Person struct {
	name string
}

func (p *Person) SetName(name string) {
	fmt.Println("Person::SetName()")
	p.name = name
	fmt.Println("New name is", p.name)
}

func (p *Person) Name() string {
	fmt.Println("Person::Name()")
	return p.name
}

func main() {
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	engine := qml.NewEngine()

	context := engine.Context()
	context.SetVar("person", &Person{name: "naeun"})
	controls, err := engine.LoadFile("helloworld.qml")
	if err != nil {
		return err
	}

	window := controls.CreateWindow(nil)

	window.Show()
	window.Wait()

	return nil
}
