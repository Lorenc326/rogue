package main

import (
	"fmt"
	"time"

	"rogue.game/core/event"
	"rogue.game/core/session"
	"rogue.game/graphic"

	term "github.com/nsf/termbox-go"
)

func main() {
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()

	sess := session.New(graphic.NewASCII(0, false, true))
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
		if sess.IsEnded {
			time.Sleep(3 * time.Second)
			break
		}
	}
}

func render(sess *session.Session) {
	term.Sync()
	fmt.Println(sess.Draw())
}
