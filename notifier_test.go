package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestNewNotifier(t *testing.T) {
	t.Log("Testing creating a new notifier")

	manager := makeStatusManager()
	_, notifier := makeTwilioNotifier(manager)

	if notifier.statusManager == nil {
		t.Error("Notifer has an invalid StatusManager")
	}
}

func TestNotifyInline(t *testing.T) {
	t.Log("Testing NotifyInline")

	manager := makeStatusManager()
	logger, notifier := makeTwilioNotifier(manager)

	manager.Enable()

	err := notifier.NotifyInline("me", "you")

	if err != nil {
		t.Errorf("NotifyInline failed: %s", err)
	}

	body, err := ioutil.ReadAll(logger)

	if err != nil {
		t.Errorf("Cannot read logger: %s", err)
	}

	if string(body) != "Send `oui` from `me` to `you`" {
		t.Errorf("Invalid message: %s", body)
	}

	manager.Toggle()
	err = notifier.NotifyInline("me", "you")

	if err != nil {
		t.Errorf("NotifyInline failed: %s", err)
	}

	body, _ = ioutil.ReadAll(logger)

	if string(body) != "Send `non` from `me` to `you`" {
		t.Errorf("Invalid message: %s", body)
	}
}

func makeTwilioNotifier(statusManager *StatusManager) (*bytes.Buffer, *TwilioNotifier) {
	errs := make(chan error, 1)
	logger := bytes.NewBufferString("")
	notifier := &TwilioNotifier{
		statusManager: statusManager,
		errs:          errs,
		sendMessage: func(from string, to string, message string) error {
			fmt.Fprintf(logger, "Send `%s` from `%s` to `%s`", message, from, to)
			return nil
		},
	}

	return logger, notifier
}
