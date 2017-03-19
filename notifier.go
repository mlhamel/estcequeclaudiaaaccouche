package main

import (
	"log"

	"github.com/sfreiberg/gotwilio"
)

type Notifier interface {
	NotifyInline(from string, to string) error
	NotifyLater(from string, to string)
}

type TwilioNotifier struct {
	sid           string
	token         string
	errs          chan error
	statusManager *StatusManager
	sendMessage   func(from string, to string, message string) error
}

func NewTwilioNotifier(sid string, token string, statusManager *StatusManager) *TwilioNotifier {
	errs := make(chan error, 1)

	return &TwilioNotifier{
		sid:           sid,
		token:         token,
		errs:          errs,
		statusManager: statusManager,
		sendMessage: func(from string, to string, message string) error {
			client := gotwilio.NewTwilioClient(sid, token)
			_, _, err := client.SendSMS(from, to, message, "", "")
			return err
		},
	}
}

func (notifier *TwilioNotifier) NotifyInline(from string, to string) error {
	log.Printf("Sending notification from %s to %s...\n", from, to)

	message := notifier.statusManager.Value()
	err := notifier.sendMessage(from, to, message)

	log.Print("Done sending notification")

	return err
}

func (notifier *TwilioNotifier) NotifyLater(from string, to string) {
	go func() {
		err := notifier.NotifyInline(from, to)

		notifier.errs <- err
	}()
}
