package main

import (
  "fmt"
  "testing"
)

var directionTests = []struct {
	facing         Direction
	rotate         Direction
	facingExpected Direction
}{
	{Down, Up, Down},
	{Down, Left, Left},
	{Down, Right, Right},

	{Up, Down, Up},
	{Up, Left, Left},
	{Up, Right, Right},

	{Left, Right, Left},
	{Left, Down, Down},
	{Left, Up, Up},

	{Right, Left, Right},
	{Right, Down, Down},
	{Right, Up, Up},
}


func TestRotations(t *testing.T) {
	for _, testContext := range directionTests {
		snake := &Snake{
			Body:   []position{{X: 10, Y: 20}},
			Facing: testContext.facing,
		}

		snake.Rotate(testContext.rotate)

		if snake.Facing != testContext.facingExpected {
			t.Errorf("Oops expected %s got %s when snake was facing %s and rotated %s", snake.Facing,
				testContext.facingExpected, testContext.facing, testContext.rotate)
		}
	}
}

var movementTests = []struct {
  initialPosition position
  facing Direction
  expectedPosition  position
}{
 {position{X: 10,Y:10}, Up, position{X:10,Y:9}},
  {position{X: 10,Y:10}, Down, position{X:10,Y:11}},
  {position{X: 10,Y:10}, Left, position{X:9,Y:10}},
  {position{X: 10,Y:10}, Right, position{X:11,Y:10}},
}

func TestMoveSnake(t *testing.T) {

  for _, testContext := range movementTests {
    snake := &Snake{
      Body:   []position{testContext.initialPosition},
      Facing: testContext.facing,
    }

    snake.Move()
    fmt.Println(snake.Body)
    if snake.HeadPosition() != testContext.expectedPosition {
      t.Errorf("Oops expected %s got %s when snake was facing %s and moved", snake.HeadPosition(),
        testContext.expectedPosition, testContext.facing )
    }

  }
}

func TestMoveSnakeBody(t *testing.T) {


    snake := &Snake{
      Body:   []position{{X: 10,Y:10}},
      Facing: Right,
    }
    for i:=0;i<=MaxLength;i++ {
      snake.Move()
    }
    expectedPosition := position {X: 16, Y:10}
  if snake.HeadPosition() != expectedPosition  {
    t.Errorf("Oops expected %s got %s when snake was moved", expectedPosition, snake.HeadPosition())
  }
  if len(snake.Body) != MaxLength  {
    t.Errorf("Oops expected snake bits wrong")
  }
    fmt.Println(snake.Body)

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
