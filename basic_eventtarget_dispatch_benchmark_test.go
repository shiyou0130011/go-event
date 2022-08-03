package event_test

import (
	"testing"

	event "github.com/shiyou0130011/go-event"
)

func BenchmarkBasicEvent_dispatchEvent(b *testing.B) {
	target := &event.BasicEventTarget{}

	target.AddEventListener(
		"foo",
		func(e event.Event) bool {
			return true
		},
	)
	for n := 0; n < b.N; n++ {
		target.DispatchEvent(SampleEvent("foo"))
	}
}
func BenchmarkBasicEvent_dispatchNoListenerEvent(b *testing.B) {
	target := &event.BasicEventTarget{}

	for n := 0; n < b.N; n++ {
		target.DispatchEvent(SampleEvent("foo"))
	}
}

func BenchmarkBasicEvent_dispatchCanaelableEvent(b *testing.B) {
	target := &event.BasicEventTarget{}

	target.AddEventListener(
		"foo",
		func(e event.Event) bool {
			return true
		},
	)

	for n := 0; n < b.N; n++ {
		target.DispatchEvent(SampleCancelableEvent("foo"))
	}
}
