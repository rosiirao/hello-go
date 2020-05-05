package main

import (
	"fmt"
	"time"
)

// MyError contains error time and detail
type MyError struct {
	When time.Time
	What string
}

func init() {
	fmt.Println(" errors init ")
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func runError() (e error) {
	fmt.Printf("---- run error ----\n")
	defer func() {
		if er := recover(); er != nil {
			e = &MyError{
				time.Now(),
				"error from recover",
			}
		}
	}()
	// panic("run error")
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
