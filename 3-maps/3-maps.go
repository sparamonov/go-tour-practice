package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	reps := make(map[string]int)	
	
	for _, val := range strings.Fields(s) {
		count := reps[val]
		reps[val] = count + 1
	}
	
	return reps
}

func main() {
	wc.Test(WordCount)
}