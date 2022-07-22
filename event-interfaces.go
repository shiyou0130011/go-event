package event

import "time"

type Listener func(e Event) bool

// The event not been dispatched
type NonDispatchedEvent interface {
	Type() string
}

// A dispatched event
type Event interface {
	NonDispatchedEvent

	Target() EventTarget
	Timestamp() time.Time
}

type EventTarget interface {
	AddEventListener(eventName string, listener Listener)
	RemoveEventListener(eventName string, listener Listener)
	//
	// To dispatch the event e.
	//
	// After dispatching, the listeners will execute a new Event with the information of e
	DispatchEvent(e NonDispatchedEvent) bool
}
