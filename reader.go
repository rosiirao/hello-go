package main

import (
	"fmt"
	"io"
	"strings"
)

// MyReader a custom reader
type MyReader struct {
	reader io.Reader
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (a *MyReader) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := p[i] + 1; char != 0 {
			buf[i] = char
		}
	}

	copy(p, buf)
	return n, nil
}

// Read string by a 8 byte-length reader
func read(s string) {
	r := strings.NewReader(s)
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

func read2() {
	b := make([]byte, 8)
	r2 := MyReader{}
	l := 0
	const MaxLoop = 2
	for {
		fmt.Printf("---- read2 current loop %v  -----\n", l)
		_, err := r2.Read(b)
		l++
		if err == io.EOF || l >= MaxLoop {
			break
		}
	}
}
