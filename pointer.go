package main

import "fmt"

// E int
type E struct {
	i int
}

// F for f()
type F interface {
	f() int
	f2() float32
}

// G for g()
type G interface {
	g()
	g2()
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

func runPointer(x interface{}) { // type switch requires an interface to introspect
	fmt.Println("--- entering runPointer ---")
	defer fmt.Println("--- leaving runPointer ---")
	switch e := x.(type) { // if x is value, then can cast it to an interface: interface{}(x)
	case E:
		fmt.Printf("pointer to E : %v \n", e.f())
	case F:
		fmt.Printf("F : %v \n", (e).f())
	default:

	}
}
