package editor

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

type Mode uint8

const VERSION = "0.0.1"

const (
	MODE_NORMAL Mode = iota
	MODE_EDIT
)

const (
	KEY_H = 104
	KEY_J = 106
	KEY_K = 107
	KEY_L = 108
	KEY_I = 105
)

type Editor struct {
	screen tcell.Screen
	buffer []rune
	numRow int
	mode   Mode
	file   os.File
	style  *Style
	cursor cursor
	height int
	width  int
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

	var buffer string

	width, height := s.Size()

	if len(os.Args) > 1 {

		fileName := os.Args[1]

		file, err := os.Open(fileName)
		if err != nil {
			log.Fatalf("failed to open: %s with error %v", fileName, err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			buffer = scanner.Text()

		}
	}

	return &Editor{
		screen: s,
		numRow: 1,
		buffer: []rune(buffer),
		mode:   MODE_NORMAL,
		style:  style,
		cursor: cursor,
		height: height,
		width:  width,
	}
}

func (ed *Editor) Run() {
	// ed.SetMode(MODE_NORMAL)
	ed.DrawRows()

	for {
		ed.screen.Show()
		ed.screen.ShowCursor(ed.cursor.x, ed.cursor.y)
		ed.EventHandler()
	}

}

func (ed *Editor) MoveCursor(x, y int) {
	if ed.width <= x || x < 0 {
		return
	} else if ed.height <= y || y < 0 {
		return
	}

	ed.cursor.x = x
	ed.cursor.y = y
	ed.screen.ShowCursor(ed.cursor.x, ed.cursor.y)
}

func (ed *Editor) DrawRows() {
	for y := 0; y < ed.height; y++ {
		if y >= ed.numRow {
			if ed.numRow == 0 && y == ed.height/3 {
				welcome := fmt.Sprintf("Terminx Editor -- version %s", VERSION)
				padding := (ed.width - len(welcome)) / 2
				for _, s := range welcome {
					ed.screen.SetContent(padding, y, s, nil, ed.style.boxStyle)
					padding += 1
				}
			}
			ed.screen.SetContent(0, y, '~', nil, ed.style.boxStyle)
		} else {
			buffLen := len(ed.buffer)
			if buffLen > ed.width {
				buffLen = ed.width
			}
			for s := 0; s < buffLen; s++ {
				ed.screen.SetContent(s, y, ed.buffer[s], nil, ed.style.boxStyle)
			}
		}
	}
}

func (ed *Editor) SetMode(mode Mode) { // is there a union type?
	ed.mode = mode
	var modeText string
	if mode == MODE_EDIT {
		modeText = "EDIT  "
	} else {
		modeText = "NORMAL"
	}

	x := 1
	y := 38 //40 is the limit on my screen for now :)

	for _, r := range []rune(modeText) {
		ed.screen.SetContent(x, y, r, nil, ed.style.boxStyle)
		x++
	}
}

func (ed *Editor) EventHandler() {
	event := ed.screen.PollEvent()

	switch eventType := event.(type) {
	case *tcell.EventKey:
		if ed.mode == MODE_NORMAL {
			if eventType.Rune() == KEY_I {
				ed.SetMode(MODE_EDIT)
				return
			}
			if eventType.Rune() == KEY_L { // l
				ed.MoveCursor(ed.cursor.x+1, ed.cursor.y)
				return
			}
			if eventType.Rune() == KEY_H { // h
				ed.MoveCursor(ed.cursor.x-1, ed.cursor.y)
				return
			}
			if eventType.Rune() == KEY_J { // j
				ed.MoveCursor(ed.cursor.x, ed.cursor.y+1)
				return
			}
			if eventType.Rune() == KEY_K { // k
				ed.MoveCursor(ed.cursor.x, ed.cursor.y-1)
				return
			}
			if eventType.Key() == tcell.KeyCtrlC {
				ed.Quit()
				return
			}
		}

		if ed.mode == MODE_EDIT {
			if eventType.Key() == tcell.KeyCtrlC {
				ed.Quit()
			}
			if eventType.Key() == tcell.KeyESC {
				ed.SetMode(MODE_NORMAL)
				return
			}
			ed.buffer = append(ed.buffer, eventType.Rune())
			ed.DrawText(1, 1, 42, 7, ed.style.boxStyle, string(ed.buffer))
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
