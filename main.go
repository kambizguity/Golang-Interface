package main

import (
	"InterfaceInGolang/educode"
	"fmt"
)

func main() {
	fmt.Println("*****")
	fmt.Println("!SECTION")

	var shapes []educode.ShapeInterface

	rectangle := educode.NewRectangle(10, 20)
	fmt.Printf("rectangle.Area(): %v\n", rectangle.Area())

	square := educode.NewSquare(10)
	fmt.Printf("square.Area(): %v\n", square.Area())

	shapes = append(shapes, square)
	shapes = append(shapes, rectangle)

	fmt.Printf("shapes: %v\n", shapes)
}
