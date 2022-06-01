package golang_united_school_homework

import ( 
	"errors"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return errors.New("out of the shapesCapacity range")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i > len(b.shapes) - 1 {
		return nil, errors.New("shape by index doesn't exist or index went out of the range")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	k, err := b.GetByIndex(i)
	if err != nil {
		return k, err
	}
	b.shapes = append( b.shapes[:i], b.shapes[i+1:]... )
	return k, err
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	k, err := b.GetByIndex(i)
	if err != nil {
		return k, err
	}
	b.shapes[i] = shape
	return k, err
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var perim float64 = 0
	for _, s := range b.shapes {
		perim += s.CalcPerimeter()
	}
	return perim
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var area float64 = 0
	for _, s := range b.shapes {
		area += s.CalcArea()
	}
	return area
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var l interface{}
	err := errors.New("circles are not exist in the list")
	k := len(b.shapes)
	for i :=0; i< k; i++ {
		l = b.shapes[i]
		_, ok := l.(*Circle)
		if !ok {
			continue
		}
		_, err = b.ExtractByIndex(k)
		k--
		i--
	}
	return err
}
