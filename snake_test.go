
package main

import (
  "fmt"
  "testing"
)

var directionTests = []struct {
  facing Direction
  rotate Direction
  facingExpected Direction
}{
  {Down, Up, Down},
  {Down, Left, Left},
  {Down, Right, Right},

  {Up, Down, Up},
  {Up, Left, Left},
  {Up, Right, Right},

  {Left, Right, Left},
  {Left, Down, Left},
  {Left, Up, Left},

  {Right, Left, Right},
  {Right, Down, Right},
  {Right, Up, Right},
}

func TestRotations(t *testing.T) {
  for _,testContext:=range directionTests{
     snake :=&Snake{
       Body: []Position {{X:10, Y:20}},
       Direction: testContext.facing,
     }

     fmt.Println("When snake is facing", testContext.facing,"and is rotated", testContext.rotate)

     if snake.Direction != testContext.facingExpected {
       t.Errorf("Ooops expected %s got %s",snake.Direction, testContext.facingExpected) }
  }
}
//func TestRotateLeft(t *testing.T){
//  snake :=&Snake{
//    Body: []Position {{X:10, Y:20}},
//    Direction: Up,
//  }
//  snake.RotateLeft()
//  if snake.Direction != Left {
//    t.Errorf("Ooops expected %d got %d",snake.Direction, Left) }
//
//}
