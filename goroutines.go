package main

import (
	"fmt"
	"time"
)

func Say(s string) {
	for i := 5; i < 5; i++ {
		go func() {
			time.Sleep(100 * time.Millisecond)
		}()
		fmt.Println(s)
	}
}
