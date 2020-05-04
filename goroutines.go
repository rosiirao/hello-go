package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 5; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
