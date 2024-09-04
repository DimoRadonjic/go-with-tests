package gomain

func Perimeter(rectangle Rectangle) float64  {
	width := rectangle.Width
	height := rectangle.Height
	return 2 * (width + height)
}


func Area(rectangle Rectangle) float64 {
	width := rectangle.Width
	height := rectangle.Height
	return width * height
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}