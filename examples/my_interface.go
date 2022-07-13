package examples

type MyInterface interface {
	Add(x, y int) int
}

type SimpleAdder struct {
}

func (s *SimpleAdder) Add(x, y int) int {
	return x + y
}

type ComplexAdder struct {
}

func (c ComplexAdder) Add(x, y int) int {
	return x * y
}

type UseExample struct {
	X int
	Y int
}

func (u *UseExample) UseAdder(a MyInterface) int {
	return a.Add(u.X, u.Y)
}
