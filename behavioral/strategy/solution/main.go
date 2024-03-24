package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (e EmailNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: Email)\n", message)
}

type SmsNotifier struct{}

func (e SmsNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: SMS)\n", message)
}

type NotifierService struct {
	notifier Notifier
}

func (ns NotifierService) Notification(message string) {
	ns.notifier.Send(message)
}

func main() {
	s := NotifierService{
		notifier: EmailNotifier{},
	}

	s.Notification("Hello HA GIA KINH")
}
