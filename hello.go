package main

import (
	"fmt"
	"strings"

	"github.com/rosiirao/hello-go/module"
)

func trace(s string) string {
	fmt.Printf("---- entering %v  ----\n", s)
	return s
}

func init() {
	fmt.Println(" hello init ")
}

func main() {
	defer fmt.Printf("\n---- leaving %v  ----\n", trace("main"))

	fmt.Println("hello, world")

	var a [3]int = [3]int{4, 5, 6}
	// var s []int
	s, s1 := a[2:], a[2:]
	// s[0] = 9
	// s1 := []int{30, 40}
	// s1[3] = 50
	// s = s1

	// fmt.Println(a[2])

	// fmt.Println(s == s)
	PrintType(len(s))
	PrintType(cap(s1))

	var m map[string]int
	// m["ok"] = 1

	fmt.Println(m["ok"])

	// sumChannel()
	// go say("world")
	// say("hello")
	// WordCount("x b")

	read("hello, world!")
	// read2()

	sumChannel()

	module.HelloModule()
	// runServer()
	if err := runError(); err != nil {
		fmt.Println(err)
	}
}

// WordCount count string
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range strings.Fields(s) {
		m[v]++
	}
	return m
}