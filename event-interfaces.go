package event

import "time"

type Listener func(e Event) bool

// The event not been dispatched
type NonDispatchedEvent interface {
	Type() string
}

// An event which can set whether it is cancelable.
//
// When an EventTarget dispatches a CancelableEvent (with Cancelable() = true) and one of the listeners return false,
// other listeners will not be executed .
type CancelableEvent interface {
	NonDispatchedEvent
	Cancelable() bool
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
