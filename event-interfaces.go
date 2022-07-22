package event

import "time"

type Listener func(e Event) bool

type Event interface {
	Type() string
	Target() EventTarget
	Timestamp() time.Time
}

type EventTarget interface {
	AddEventListener(eventName string, listener Listener)
	RemoveEventListener(eventName string, listener Listener)
	DispatchEvent(e Event) bool
}
