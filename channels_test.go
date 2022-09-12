package main

import "testing"

func TestChannel(t *testing.T) {
	c := make(chan int)

	c <- 1
	c <- 2
	<-c
	<-c

	close(c)
}
