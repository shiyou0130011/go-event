package event

import (
	"reflect"
	"time"
)

type BasicEvent struct {
	NonDispatchedEvent
	target    EventTarget
	timestamp time.Time
}

func (e BasicEvent) Target() EventTarget {
	return e.target
}
func (e BasicEvent) Timestamp() time.Time {
	return e.timestamp
}

type BasicEventTarget struct {
	listeners map[string][]Listener
}

func (t BasicEventTarget) AddEventListener(eventName string, listener Listener) {
	lAdded := reflect.ValueOf(listener)
	for _, l := range t.listeners[eventName] {
		l1 := reflect.ValueOf(l)

		if l1.Pointer() == lAdded.Pointer() {
			return
		}
	}

	t.listeners[eventName] = append(t.listeners[eventName], listener)
}

func (t BasicEventTarget) RemoveEventListener(eventName string, listener Listener) {
	if len(t.listeners[eventName]) == 0 {
		return
	}
	lRemoved := reflect.ValueOf(listener)
	for i, l := range t.listeners[eventName] {
		l1 := reflect.ValueOf(l)

		if l1.Pointer() == lRemoved.Pointer() {
			t.listeners[eventName] = append(t.listeners[eventName][:i], t.listeners[eventName][i+1:]...)

			return
		}
	}
}

func (t BasicEventTarget) DispatchEvent(e NonDispatchedEvent) bool {
	eventName := e.Type()

	event := BasicEvent{
		NonDispatchedEvent: e,
		timestamp:          time.Now(),
		target:             t,
	}

	for _, listener := range t.listeners[eventName] {
		listener(event)
	}
	return true
}
