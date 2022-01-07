package main

import "testing"

func TestWordCount(t *testing.T) {
	sentence := "hello world it's king of the world"
	want := map[string]int{
		"hello": 1,
		"world": 2,
		"it's":  1,
		"king":  1,
		"of":    1,
		"the":   1,
	}

	got := WordCount(sentence)
	for word, count := range want {
		if w, ok := got[word]; !ok || count != w {
			t.Errorf("%v appears %d time(s), you got %d time(s)", word, count, w)
		}
	}
}
