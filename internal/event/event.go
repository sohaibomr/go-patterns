package event

import (
	"fmt"

	"github.com/sohaibomr/test-coverage/internal/event/email"
	"github.com/sohaibomr/test-coverage/internal/event/logs"
)

type EventHandler func(string) error

type EventListner struct {
	handlers map[string]EventHandler
}

func NewEventListner() *EventListner {
	handlers := map[string]EventHandler{}
	return &EventListner{handlers: handlers}
}

func (el *EventListner) registerEventHandler(eventKey string, handlerFunc EventHandler) {
	el.handlers[eventKey] = handlerFunc
}

func (el *EventListner) RegisterHandlers() {
	el.registerEventHandler("FEEDBACK_EMAIL", email.FeedbackEmail)
	el.registerEventHandler("AUDIT_LOGS", logs.HandleAuditLog)

}

func (el *EventListner) Handle(eventKey string) error {
	if handler, ok := el.handlers[eventKey]; ok {
		return handler(eventKey)
	}
	fmt.Println("Handler not found for given key")
	return nil
}
