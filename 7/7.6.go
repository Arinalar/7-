// Ларионова Арина 363
package main

import (
	"fmt"
	"sync"
)

type EventBus struct {
	mu          sync.RWMutex
	subscribers map[string][]func(interface{})
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]func(interface{})),
	}
}

func (eb *EventBus) Subscribe(event string, handler func(interface{})) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	eb.subscribers[event] = append(eb.subscribers[event], handler)
}

func (eb *EventBus) Publish(event string, data interface{}) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()
	if handlers, found := eb.subscribers[event]; found {
		for _, handler := range handlers {
			handler(data)
		}
	}
}

func main() {
	eventBus := NewEventBus()
	eventBus.Subscribe("hehe", func(data interface{}) {
		fmt.Printf("Созданное пользователем событие имеет следующие даныыен: %v\n", data)
	})
	eventBus.Publish("bla bla", map[string]interface{}{"id": 1, "name": "Arnir"})
}
