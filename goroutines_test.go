package main

import (
	"fmt"
	"testing"
)

var aStr string

func TestSay(t *testing.T) {
	a := make(chan int)
	b := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			a <- i
			aStr = "aStr = 1"
		}
		close(a)
	}()
	go func() {
		for i := 0; i < 10; i++ {
			b <- fmt.Sprintf("b %v", i)
			aStr = "aStr = 2"
		}
		close(b)
		fmt.Println(aStr)
	}()

	func() {
		var t1 int
		var s1 string
		defer func() {
			fmt.Printf("defer t1 %v\n", t1)
			fmt.Printf("defer s1 %v\n", s1)
		}()
		for {
			// randomly choose one of the channels of a or b
			select {
			case x, ok := <-a:
				if !ok {
					return
				}
				t1 = x
				fmt.Printf("x %v\n", x)
			case y, ok := <-b:
				if !ok {
					return
				}
				fmt.Printf("f %v\n", y)
			}
		}
	}()

}
