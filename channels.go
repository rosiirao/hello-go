package main

import "fmt"

func sum(s []int, c chan int, b chan bool) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
	b <- true
}

func sumChannel() {
	fmt.Println(` --- entering sumChannel --- `)
	defer fmt.Println(` --- leaving sumChannel ---`)
	s := []int{7, 2, 8, -9, 4, 0}
	// By default, sends and receives block until the other side is ready.
	// This allows goroutines to synchronize without explicit locks or condition variables.
	// Send to buffered channel block when the buffer is full.
	// Receivers block when the buffer is empty
	c := make(chan int, 100)
	b := make([]chan bool, 2)
	b[0] = make(chan bool)
	b[1] = make(chan bool)
	go sum(s[:len(s)/2], c, b[0])
	go sum(s[len(s)/2:], c, b[1])
	// cap(c)
	var sum int
	for _, v := range b {
		if !<-v {
			panic("sum routine error!")
		}
	}
	close(c)
	for n := range c {
		sum = sum + n
	}
	fmt.Println(cap(c), sum)
}
