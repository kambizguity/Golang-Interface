package main

import (
	"InterfaceInGolang/shape"
	"fmt"
)

func main() {
	//SECTION - I, Interface with shape object
	var shapes []shape.ShapeInterface

	rectangle := shape.NewRectangle(10, 20)
	fmt.Printf("rectangle.Area(): %v\n", rectangle.Area())

	square := shape.NewSquare(10)
	fmt.Printf("square.Area(): %v\n", square.Area())

	shapes = append(shapes, square)
	shapes = append(shapes, rectangle)

	fmt.Printf("shapes: %v\n", shapes)

	//SECTION - II, Use shape object as ShapeInterface as parameter of function.
	fmt.Printf("shape.CalculateArea(square): %v\n", shape.CalculateArea(square))
	fmt.Printf("shape.CalculateArea(rectangle): %v\n", shape.CalculateArea(rectangle))

	//SECTION - III, Use interface to get different type of parameters and identify the type Assertion find the type of parameters, Type Assertion
	//Iteration on array
	for _, obj := range shapes {
		switch object := obj.(type) {
		case shape.Rectangle:
			fmt.Printf("The object is rectangle with width:%d\n, height:%d", object.Width, object.Height)
		case shape.Square:
			fmt.Printf("The object is square with lenght:%d\n", object.Length)
		}
	}
}
