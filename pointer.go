package main

import "fmt"

// E int
type E struct {
	i int
}

// F for f()
type F interface {
	f() int
}

// G for g()
type G interface {
	g()
}

func (e *E) f() int {
	return e.i
}

func pTE() *E {
	e := E{
		10,
	}
	return &e
}

func tE() E {
	return E{
		20,
	}
}

func runPointer(x interface{}) {
	fmt.Println("--- entering runPointer ---")
	defer fmt.Println("--- leaving runPointer ---")
	switch e := x.(type) {
	case E:
		fmt.Printf("pointer to E : %v \n", e.f())
	case F:
		fmt.Printf("F : %v \n", (e).f())
	default:

	}
}
