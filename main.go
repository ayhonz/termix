package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

func main() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)


	for {
		s.Show()

		event := s.PollEvent()

		switch eventType := event.(type) {
		case *tcell.EventKey:
			if eventType.Key() == tcell.KeyCtrlC {
				return
			}
		}
	}

}
