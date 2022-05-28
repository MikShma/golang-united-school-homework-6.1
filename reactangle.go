package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (c Rectangle) CalcArea() float64 {
	return c.Height * c.Weight
}


func (c Rectangle) CalcPerimeter() float64 {
	return 2 * ( c.Height + c.Weight )
}
