package main
import "fmt"
type Direction int

const MaxLength = 5
const (
	Up Direction = iota
	Down
	Left
	Right
)

var directions = [...]string{
	"Up",
	"Down",
	"Left",
	"Right",
}

// String returns the English name of the Direction
func (d Direction) String() string {
	if Up <= d && d <= Right {
		return directions[d]
	}
	return ""
}


type position struct {
	X int
	Y int
}

// String returns the English name of the Direction
func (p position) String() string {
  return fmt.Sprintf("{X : %d , Y: %d}",p.X, p.Y)
}
func (p position) move(dx, dy int) position {

    p.X= p.X + dx
    p.Y = p.Y + dy

    return p
}

type Snake struct {
	Body   []position
	Facing Direction
}

// Rotate
func (s *Snake) Rotate(direction Direction) {
	if direction == Left && s.Facing != Right {
		s.Facing = Left
	} else if direction == Right && s.Facing != Left {
		s.Facing = Right
	} else if direction == Down && s.Facing != Up {
		s.Facing = Down
	} else if direction == Up && s.Facing != Down {
		s.Facing = Up
	}
}

func (s *Snake) Move()  {
  var dx,dy int
  //naive approach for now
  if s.Facing == Left {
      dx= -1
  }
  if s.Facing == Right {
    dx= 1
  }
  if s.Facing == Up {
    dy= -1
  }

  if s.Facing == Down {
    dy= 1
  }

  var newHead = s.HeadPosition().move(dx, dy)

  s.Body= append([]position{newHead}, s.Body...)

  if len(s.Body) > MaxLength {
    s.Body = s.Body[:len(s.Body)-1]
  }
}


func (s *Snake) HeadPosition() position {
  return s.Body[0]
}
