package main

import (
	"testing"
)

const raw string = `ToCountry=CA
&ToState=Quebec
&SmsMessageSid=SM0cee2f5a06c9726753c7231249572309
&NumMedia=0
&ToCity=
&FromZip=J7Y+3W2
&SmsSid=SM0cee2f5a06c9726753c7231249572309
&FromState=QC
&SmsStatus=received
&FromCity=MONTREAL
&Body=Plop
&FromCountry=CA
&To=%2B14388004514
&ToZip=
&NumSegments=1
&MessageSid=SM0cee2f5a06c9726753c7231249572309
&AccountSid=ACa0e4c544ea66c6ff500da67e39730dfb
&From=%2B15145692911
&ApiVersion=2010-04-01`

func TestParseBadQuery(t *testing.T) {
	t.Log("Parsing query string")

	_, err := NewRequestInfo("")

	if err != nil {
		t.Error("Parsing empty query string should return an error")
	}
}

func TestParseQuery(t *testing.T) {
	t.Log("Parsing query string")

	query, err := NewRequestInfo(raw)

	if err != nil {
		t.Error(err)
	}

	if query.ToState != "Quebec" {
		t.Error("ToState from query string was not parsed properly")
	}

	if query.SmsMessageSid != "SM0cee2f5a06c9726753c7231249572309" {
		t.Error("SmsMessageSid from query string was not parsed properly")
	}

	if query.NumMedia != "0" {
		t.Error("NumMedia from query string was not parsed properly")
	}

	if query.FromZip != "J7Y 3W2" {
		t.Error("FromZip from query string was not parsed properly")
	}

	if query.SmsSid != "SM0cee2f5a06c9726753c7231249572309" {
		t.Error("SmsMessageSid from query string was not parsed properly")
	}

	if query.FromState != "QC" {
		t.Error("FromState from query string was not parsed properly")
	}

	if query.SmsStatus != "received" {
		t.Error("SmsStatus from query string was not parsed properly")
	}

	if query.FromCity != "MONTREAL" {
		t.Error("FromCity from query string was not parsed properly")
	}

	if query.Body != "Plop" {
		t.Error("Body from query string was not parsed properly")
	}

	if query.FromCountry != "CA" {
		t.Error("FromCountry from query string was not parsed properly")
	}

	if query.To != "+14388004514" {
		t.Error("To from query string was not parsed properly")
	}

	if query.NumSegments != "1" {
		t.Error("NumSegments from query string was not parsed properly")
	}
}
