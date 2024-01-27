package editor

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

type Mode uint8

const (
	MODE_NORMAL Mode = iota
	MODE_EDIT
)

type Editor struct {
	screen tcell.Screen
	buffer string // just a joke!
	mode   Mode
	file   string
}

// probably should split this one into init and run
func (ed *Editor) Run() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorReset)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	s.SetStyle(defStyle)
	s.Clear()

	ed.screen = s

	s.ShowCursor(1, 1)
	x, y := 1, 1

	for {
		s.Show()
		ed.DrawText(1, 1, 42, 7, boxStyle, "Hello")
		event := s.PollEvent()

		switch eventType := event.(type) {
		case *tcell.EventKey:
			if eventType.Key() == tcell.KeyCtrlC {
				return
			}
			if eventType.Rune() == 108  { // l
				x += 1
				s.ShowCursor(x, y)
			}
			if eventType.Rune() == 104  { // h 
				x -= 1
				s.ShowCursor(x, y)
			}
			if eventType.Rune() == 106  { // j 
				y += 1
				s.ShowCursor(x, y)
			}
			if eventType.Rune() == 107  { // k
				y -= 1
				s.ShowCursor(x, y)
			}
			// mod, key, ch, name := eventType.Modifiers(), eventType.Key(), eventType.Rune(), eventType.Name()
			// log.Fatalf("EventKey Modifiers: %d Key: %d Rune: %d, Name: %s", mod, key, ch, name)

		}
	}

}

func (ed *Editor) DrawText(x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		ed.screen.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func (ed *Editor) Quit() {
	// You have to catch panics in a defer, clean up, and
	// re-raise them - otherwise your application can
	// die without leaving any diagnostic trace.
	maybePanic := recover()
	ed.screen.Fini()
	if maybePanic != nil {
		panic(maybePanic)
	}
}
