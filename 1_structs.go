package main

// https://golang.org/doc/faq#Is_Go_an_object-oriented_language

// don't mix https://tour.golang.org/methods/8

type Vertex struct {
	X, Y int
}

// func (v Vertex) Abs() {
// 	v.X = 1
// }

func (v *Vertex) Abs() { //change to abs and check interface implementation
	v.X = 1
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
	d  = new(Vertex)
)

type Abser interface {
	Abs()
}
