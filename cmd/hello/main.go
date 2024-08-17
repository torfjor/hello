package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	fmt.Println("hello, world")
	words := []string{"one", "two", "three", "four"}
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println(strings.Join(words, " "))
}
