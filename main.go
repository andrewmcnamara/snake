package main

import (
	"fmt"
)

func main() {
	snake := &Snake{
		Body:   []position{{X: 10, Y: 20}},
		Facing: Left,
	}

	fmt.Println("Hello", snake)

}
