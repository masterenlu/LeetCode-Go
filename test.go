package main

import "fmt"

func main() {
	path := []int{0, 1, 2, 3}
	another := make([]int, len(path))
	copy(another, path)

	fmt.Printf("p: %v, anP: %v\n", path, another)
	path[0] = 9
	fmt.Printf("p: %v, anP: %v\n", path, another)
}
