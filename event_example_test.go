package event_test

import (
	"log"

	event "github.com/shiyou0130011/go-event"
)

type SampleEvent string

func (e SampleEvent) Type() string {
	return string(e)
}

func Example() {
	target := &event.BasicEventTarget{}
	target.AddEventListener("foo", func(e event.Event) bool {
		log.Println("Dispatch foo event")
		return true
	})
	target.AddEventListener("foo", func(e event.Event) bool {
		log.Printf("Dispatch event at %v", e.Timestamp())
		return true
	})

	target.DispatchEvent(SampleEvent("foo"))

	// Output:
	//
}
