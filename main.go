package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell"
)

func drawSnake(screen tcell.Screen, snake Snake) {
	w, h := screen.Size()
	screen.Clear()

	if w == 0 || h == 0 {
		return
	}
	st := tcell.StyleDefault
	st.Background(tcell.ColorGold)

	for idx, pos := range snake.Body {
		var partRune rune
		if idx == 0 {
			partRune = 'O'
		} else {
			partRune = '#'
		}

		screen.SetCell(pos.X, pos.Y, st, partRune)
	}

	screen.Show()
}

func main() {
	snake := &Snake{
		Body:   []position{{X: 10, Y: 20}},
		Facing: Right,
	}
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)

	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e = s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	s.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorBlack).
		Background(tcell.ColorWhite))
	s.Clear()
	w, h := s.Size()
	fmt.Println("Hello", snake, w, h)

	quit := make(chan struct{})
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter:
					close(quit)
					return
				case tcell.KeyCtrlL:
					s.Sync()
				}
			case *tcell.EventResize:
				s.Sync()
			}
		}
	}()

loop:
	for {
		select {
		case <-quit:
			break loop
		case <-time.After(time.Millisecond * 50):
		}
		snake.Move()
		pos := snake.HeadPosition()
		if pos.X == (w - 1) {
			if pos.Y == 1 {
				snake.Rotate(Left)
			} else {
				snake.Rotate(Up)
			}
		} else if pos.X == 1 {
			if pos.Y == 1 && snake.Facing == Left {
				snake.Rotate(Down)
			} else if pos.Y == (h-1) && snake.Facing == Down {
				snake.Rotate(Right)
			}
		}
		drawSnake(s, *snake)
		// fmt.Println(start)

	}

}
