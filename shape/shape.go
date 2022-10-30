package shape

type ShapeInterface interface {
	Area() int
}

type Square struct {
	Length int
}

type Rectangle struct {
	Height, Width int
}

//Calculate area of squre
func (s Square) Area() int {
	return (s.Length * s.Length)
}

//Calculate area of rectangle
func (r Rectangle) Area() int {
	return r.Height * r.Width
}

//Constructor of rectangle with specific wigth and Height
func NewRectangle(height, width int) Rectangle {
	var object Rectangle
	object.Height = height
	object.Width = width
	return object
}

//Constructor of square with lenght
func NewSquare(lenght int) Square {
	var object Square
	object.Length = lenght
	return object
}

func CalculateArea(obj ShapeInterface) int {
	return obj.Area()
}
