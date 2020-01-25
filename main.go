package main

import (
  "fmt"
  )

func main() {
  snake :=&Snake{
    Body: []Position {{X:10, Y:20}},
    Direction: Left,
  }


  fmt.Println("Hello",snake)

}
