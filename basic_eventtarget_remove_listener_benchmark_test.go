package event_test

import (
	"testing"

	event "github.com/shiyou0130011/go-event"
)

func BenchmarkRemoveListener(b *testing.B) {
	target := &event.BasicEventTarget{}
	listener := func(e event.Event) bool { return true }

	for i := 0; i < b.N; i++ {
		target.AddEventListener("foo", listener)
		target.RemoveEventListener("foo", listener)
	}
}
