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
