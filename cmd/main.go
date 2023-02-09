package main

import (
	"fmt"

	"github.com/Lorenc326/rogue/core/event"
	"github.com/Lorenc326/rogue/core/session"
	"github.com/Lorenc326/rogue/graphic"

	term "github.com/nsf/termbox-go"
)

func main() {
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()

	sess := session.New(
		graphic.NewASCII(0, false, true),
		session.SessionParametrs{Seed: 100, Width: 60, Height: 40},
	)
	render(sess)

	for {
		e := term.PollEvent()
		switch e.Type {
		case term.EventKey:
			switch e.Key {
			case term.KeyEsc:
				return
			case term.KeyArrowUp:
				sess.React(event.Event{Action: "move", Direction: "up"})
			case term.KeyArrowRight:
				sess.React(event.Event{Action: "move", Direction: "right"})
			case term.KeyArrowDown:
				sess.React(event.Event{Action: "move", Direction: "down"})
			case term.KeyArrowLeft:
				sess.React(event.Event{Action: "move", Direction: "left"})
			}
		}
		render(sess)
	}
}

func render(sess *session.Session) {
	term.Sync()
	fmt.Println(sess.Draw())
}
