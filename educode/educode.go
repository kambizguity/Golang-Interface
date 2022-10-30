package educode

type ShapeInterface interface {
	Area() int
}

type Square struct {
	Length int
}

type Rectangle struct {
	Heigth, Width int
}

//Calculate area of squre
func (s Square) Area() int {
	return (s.Length * s.Length)
}

//Calculate area of rectangle
func (r Rectangle) Area() int {
	return r.Heigth * r.Width
}

//Constructor of rectangle with specific wigth and heigth
func NewRectangle(height, width int) Rectangle {
	var object Rectangle
	object.Heigth = height
	object.Width = width
	return object
}

//Constructor of square with lenght
func NewSquare(lenght int) Square {
	var object Square
	object.Length = lenght
	return object
}
