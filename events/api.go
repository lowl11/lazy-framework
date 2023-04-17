package events

import "github.com/lowl11/lazy-framework/events/script_event"

var (
	Script *script_event.Event
)

func Init() {
	var err error

	Script, err = script_event.New()
	if err != nil {
		panic(err)
	}
}
