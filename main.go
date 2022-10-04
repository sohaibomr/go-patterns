package main

import (
	"fmt"

	"github.com/sohaibomr/test-coverage/internal/api"
	"github.com/sohaibomr/test-coverage/internal/event"
)

func main() {
	a := api.Sum(10, 10)
	fmt.Println(a)
	listner := event.NewEventListner()
	listner.RegisterHandlers()
	listner.Handle("FEEDBACK_EMAIL")
	listner.Handle("FEEDBACK_EMAIL")
	listner.Handle("AUDIT_LOGS")
}
