package main

import (
	"fmt"
	"testing"
)

func TestPrintType(t *testing.T) {
	fmt.Println("==== test PrintType ====")

	PrintType(len(make([]interface{}, 10)))
	PrintType(cap(make([]interface{}, 10)))
	PrintType(func(i interface{}) {})
}
