package main

// E int
type E struct {
	i int
}

func tE() *E {
	e := E{
		10,
	}
	return &e
}

func runPointer() int {
	return tE().i
}
