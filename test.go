package main

import "fmt"

var result []int

func main() {
	path := 1
	another := path
	fmt.Printf("p: %v, anP: %v\n", path, another)
	path = 2
	fmt.Printf("p: %v, anP: %v\n", path, another)
}
