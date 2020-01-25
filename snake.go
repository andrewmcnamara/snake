package main

type Direction int

const (
  Up Direction = iota
  Down
  Left
  Right
)

var directions = [...] string{
  "Up",
  "Down",
  "left",
  "Right",
}

// String returns the English name of the day ("Sunday", "Monday", ...).
func (d Direction) String() string {
  if Up <= d && d <= Right {
    return directions[d]
  }
  return ""
  //"Direction(" +d + ")"
}

type Position struct {
  X  int
  Y int
}

type Snake struct {
  Body [] Position
  Direction Direction

}

func ( s *Snake) RotateLeft() {
  if s.Direction != Right{
    s.Direction = Left
  }
}

func ( s *Snake) RotateRight() {
  if s.Direction != Left{
    s.Direction = Right
  }
}

func ( s *Snake) RotateDown() {
  if s.Direction != Up{
    s.Direction = Down
  }
}

func ( s *Snake) RotateUp() {
  if s.Direction != Down{
    s.Direction = Up
  }
}
