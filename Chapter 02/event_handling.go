package main

import "fmt"

type EventHandler func()

// EventDispatcher keeps track of registered events and their handlers
type EventDispatcher struct {
 events map[string]EventHandler
}

// RegisterEvent associates an event with a handler
func (ed *EventDispatcher) RegisterEvent(event string, handler EventHandler) {
 ed.events[event] = handler
}

// TriggerEvent simulates an event by invoking its handler
func (ed *EventDispatcher) TriggerEvent(event string) {
 handler, exists := ed.events[event]
 if exists {
  handler() // Invoke the event handler
 } else {
  fmt.Println("Event not found:", event)
 }
}

func main() {
 // Create an EventDispatcher
 dispatcher := EventDispatcher{
  events: make(map[string]EventHandler),
 }

 // Define event handlers
 onLogin := func() {
  fmt.Println("User logged in")
 }

 onLogout := func() {
  fmt.Println("User logged out")
 }

 // Register events with their handlers
 dispatcher.RegisterEvent("login", onLogin)
 dispatcher.RegisterEvent("logout", onLogout)

 // Simulate events
 dispatcher.TriggerEvent("login")         // Simulate a user login event
 dispatcher.TriggerEvent("logout")        // Simulate a user logout event
 dispatcher.TriggerEvent("unknown_event") // Simulate an unknown event
}
