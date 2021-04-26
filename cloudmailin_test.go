package cloudmailin

import (
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		Fn, To, From, Bcc, Subject string
	}{
		{
			Fn:      "testdata/example.json",
			To:      "to@example.co.uk",
			From:    "Message Sender <sender@example.com>",
			Bcc:     "",
			Subject: "Test Subject",
		},
	}

	for _, tt := range tests {
		f, err := os.Open(tt.Fn)
		defer f.Close()

		d, err := Decode(f)
		if err != nil {
			t.Errorf("got unexpected error %v", err)
		}

		if d.Headers.To != tt.To {
			t.Errorf("d.Headers.To = %v, want %v", d.Headers.To, tt.To)
		}

		if d.Headers.From != tt.From {
			t.Errorf("d.Headers.From = %v, want %v", d.Headers.From, tt.From)
		}

		if d.Headers.Subject != tt.Subject {
			t.Errorf("d.Headers.Subject = %v, want %v", d.Headers.Subject, tt.Subject)
		}
	}
}

func TestDecodeCloudmailin(t *testing.T) {
	tests := []struct {
		Fn, To, From, Bcc, Subject string
	}{
		{
			Fn:      "testdata/cloudmailin.json",
			To:      "asdf9390933023+32222222-2222-2222-2222-222222222222@cloudmailin.net",
			From:    "example The Prelude <noreply@example.dev>",
			Bcc:     "",
			Subject: "example account confirmation",
		},
	}

	for _, tt := range tests {
		f, err := os.Open(tt.Fn)
		defer f.Close()

		d, err := Decode(f)
		if err != nil {
			t.Errorf("got unexpected error %v", err)
		}

		if d.Headers.To != tt.To {
			t.Errorf("d.Headers.To = %v, want %v", d.Headers.To, tt.To)
		}

		if d.Headers.From != tt.From {
			t.Errorf("d.Headers.From = %v, want %v", d.Headers.From, tt.From)
		}

		if d.Headers.Subject != tt.Subject {
			t.Errorf("d.Headers.Subject = %v, want %v", d.Headers.Subject, tt.Subject)
		}
	}
}

func TestDecodeCloudmailin2(t *testing.T) {
	tests := []struct {
		Fn, To, From, Bcc, Subject string
	}{
		{
			Fn:      "testdata/cloudmailin_string_headers_received.json",
			To:      "12345678+makkara2@cloudmailin.net",
			From:    "Hermanni Hiiri <Hermanni.Hiiri@iki.fi>",
			Bcc:     "",
			Subject: "asdfsdf",
		},
	}

	for _, tt := range tests {
		f, err := os.Open(tt.Fn)
		defer f.Close()

		d, err := Decode(f)
		if err != nil {
			t.Errorf("got unexpected error %v", err)
		}

		if d.Headers.To != tt.To {
			t.Errorf("d.Headers.To = %v, want %v", d.Headers.To, tt.To)
		}

		if d.Headers.From != tt.From {
			t.Errorf("d.Headers.From = %v, want %v", d.Headers.From, tt.From)
		}

		if d.Headers.Subject != tt.Subject {
			t.Errorf("d.Headers.Subject = %v, want %v", d.Headers.Subject, tt.Subject)
		}
	}
}
