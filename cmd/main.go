package main

import (
	"fmt"
	"strings"
	"time"

	"rogue.game/core/session"

	term "github.com/nsf/termbox-go"
)

func main() {
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()

	sess := session.Session{}
	sess.Init()
	render(&sess)

	for {
		event := term.PollEvent()
		switch event.Type {
		case term.EventKey:
			switch event.Key {
			case term.KeyEsc:
				return
			case term.KeyArrowUp:
				sess.React(session.Event{Action: "move", Direction: "up"})
			case term.KeyArrowRight:
				sess.React(session.Event{Action: "move", Direction: "right"})
			case term.KeyArrowDown:
				sess.React(session.Event{Action: "move", Direction: "down"})
			case term.KeyArrowLeft:
				sess.React(session.Event{Action: "move", Direction: "left"})
			}
		}
		render(&sess)
		if sess.IsEnded {
			time.Sleep(3 * time.Second)
			break
		}
	}
}

func render(sess *session.Session) {
	term.Sync()
	fmt.Println(widerOutput(sess.RenderASCII()))
}

func widerOutput(render string) string {
	return " " + strings.Join(strings.Split(render, ""), " ")
}
