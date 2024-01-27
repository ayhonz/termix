package editor

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

type Mode uint8

const (
	MODE_NORMAL Mode = iota
	MODE_EDIT
)

const (
	KEY_H = 104
	KEY_J = 106
	KEY_K = 107
	KEY_L = 108
)

type Editor struct {
	screen tcell.Screen
	buffer []rune
	mode   Mode
	file   string
	style  *Style
	cursor cursor
}

type Style struct {
	defaultStyle tcell.Style
	boxStyle     tcell.Style
}

type cursor struct {
	x int
	y int
}

func Init() *Editor {
	style := &Style{
		defaultStyle: tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset),
		boxStyle:     tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorReset),
	}

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	s.SetStyle(style.defaultStyle)
	s.Clear()

	cursor := cursor{
		x: 1,
		y: 1,
	}

	return &Editor{
		screen: s,
		buffer: []rune("Hello there my friend"),
		mode:   MODE_NORMAL,
		style:  style,
		cursor: cursor,
	}

}

func (ed *Editor) Run() {
	ed.DrawText(1, 1, 42, 7, ed.style.boxStyle, string(ed.buffer))

	for {
		ed.screen.Show()
		ed.screen.ShowCursor(ed.cursor.x, ed.cursor.y)
		ed.EventHandler()
	}

}

func (ed *Editor) EventHandler() {
	event := ed.screen.PollEvent()

	switch eventType := event.(type) {
	case *tcell.EventKey:
		if eventType.Key() == tcell.KeyCtrlC {
			ed.Quit()
		}
		if eventType.Rune() == KEY_L { // l
			ed.cursor.x += 1
			ed.screen.ShowCursor(ed.cursor.x, ed.cursor.y)
		}
		if eventType.Rune() == KEY_H { // h
			ed.cursor.x -= 1
			ed.screen.ShowCursor(ed.cursor.x, ed.cursor.y)
		}
		if eventType.Rune() == KEY_J { // j
			ed.cursor.y += 1
			ed.screen.ShowCursor(ed.cursor.x, ed.cursor.y)
		}
		if eventType.Rune() == KEY_K { // k
			ed.cursor.y -= 1
			ed.screen.ShowCursor(ed.cursor.x, ed.cursor.y)
		}
		// mod, key, ch, name := eventType.Modifiers(), eventType.Key(), eventType.Rune(), eventType.Name()
		// log.Fatalf("EventKey Modifiers: %d Key: %d Rune: %d, Name: %s", mod, key, ch, name)
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
	maybePanic := recover()
	if maybePanic != nil {
		panic(maybePanic)
	}
	ed.screen.Fini()
	os.Exit(0)
}
