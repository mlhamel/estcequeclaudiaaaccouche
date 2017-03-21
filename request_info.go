package main

import (
	"net/url"
	"reflect"
	"strings"

	"github.com/gorilla/schema"
)

// RequestInfo is a representation of a parsed query string from twilio
type RequestInfo struct {
	ToCountry     string
	ToState       string
	SmsMessageSid string
	NumMedia      string
	ToCity        string
	FromZip       string
	SmsSid        string
	FromState     string
	SmsStatus     string
	FromCity      string
	Body          string
	FromCountry   string
	To            string
	ToZip         string
	NumSegments   string
	MessageSid    string
	AccountSid    string
	From          string
	ApiVersion    string
	MediaUrl0     string
}

// NewRequestInfo parse raw requery string and turns it to a RequestInfo
//
// To use it, just type:
//
//	info, err := NewRequestInfo(req.Body)
//
func NewRequestInfo(raw string) (*RequestInfo, error) {
	info := new(RequestInfo)
	decoder := schema.NewDecoder()

	decoder.RegisterConverter(info.Body,
		func(s string) reflect.Value { return reflect.ValueOf(strings.TrimSpace(s)) })

	parsed, err := url.ParseQuery(raw)

	if err != nil {
		return nil, err
	}

	err = decoder.Decode(info, parsed)

	return info, err
}
