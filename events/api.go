package events

import "github.com/lowl11/lazy-framework/events/script_event"

var (
	Script *script_event.Event
)

func Init(useDatabase bool) {
	var err error

	if useDatabase {
		Script, err = script_event.New()
		if err != nil {
			panic(err)
		}
	}
}
