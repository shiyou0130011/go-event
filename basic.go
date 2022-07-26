package event

import (
	"reflect"
	"sync"
	"time"
)

type BasicEvent struct {
	Original  NonDispatchedEvent
	target    EventTarget
	timestamp time.Time
}

func (e BasicEvent) Type() string {
	return e.Original.Type()
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

func (t *BasicEventTarget) AddEventListener(eventName string, listener Listener) {
	if t.listeners == nil {
		t.listeners = make(map[string][]Listener)
	}

	lAdded := reflect.ValueOf(listener)
	if list, has := t.listeners[eventName]; has {
		for _, l := range list {
			l1 := reflect.ValueOf(l)

			if l1.Pointer() == lAdded.Pointer() {
				return
			}
		}
	}
	t.listeners[eventName] = append(t.listeners[eventName], listener)
}

func (t *BasicEventTarget) RemoveEventListener(eventName string, listener Listener) {
	if t.listeners[eventName] == nil || len(t.listeners[eventName]) == 0 {
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

func (t *BasicEventTarget) DispatchEvent(e NonDispatchedEvent) (result bool) {
	result = true
	var cancelable = false
	if ce, isCancelableEvent := e.(CancelableEvent); isCancelableEvent {
		cancelable = ce.Cancelable()
	}

	event := BasicEvent{
		Original:  e,
		timestamp: time.Now(),
		target:    t,
	}

	list, has := t.listeners[e.Type()]
	if !has || len(list) == 0 {
		return
	}

	if cancelable {
		result = t.dispatchCancelableEvent(event)
	} else {
		t.dispatchNonCancelableEvent(event)
	}

	return
}

func (t *BasicEventTarget) dispatchCancelableEvent(e Event) (result bool) {
	for _, listener := range t.listeners[e.Type()] {
		result = listener(e)
		if !result {
			return
		}
	}
	return
}

func (t *BasicEventTarget) dispatchNonCancelableEvent(e Event) {
	var wg sync.WaitGroup
	for _, listener := range t.listeners[e.Type()] {
		wg.Add(1)
		go (func(l Listener) {
			defer wg.Done()

			l(e)
		})(listener)
	}
	wg.Wait()
}
