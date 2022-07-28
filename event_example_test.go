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

// A event which is cancelable
type SampleCancelableEvent string

func (e SampleCancelableEvent) Type() string {
	return string(e)
}
func (e SampleCancelableEvent) Cancelable() bool {
	return true
}
func ExampleCancelAbleNonDispatchedEvent() {
	target := &event.BasicEventTarget{}
	target.AddEventListener("foo", func(e event.Event) bool {
		log.Println("Dispatch foo event and return false")
		return false
	})
	target.AddEventListener("foo", func(e event.Event) bool {
		// The 2nd event listener.
		// Since the 1st listener return false, this listener will not be executed.
		log.Printf("Dispatch foo event at 2nd times")
		return true
	})

	target.DispatchEvent(SampleCancelableEvent("foo"))

	// Output:
	//
}
