package cloudmailin

import (
	"encoding/json"
	"io"
)

// A SPF (Sender Policy Framework) result for the given IP address and Domain.
type SPF struct {
	Result string `json:"result"`
	Domain string `json:"domain"`
}

// Envelope contains the data sent or gathered from the remote server.
type Envelope struct {
	To         string   `json:"to"`
	Recipients []string `json:"recipients"`
	From       string   `json:"from"`
	HeloDomain string   `json:"helo_domain"`
	RemoteIP   string   `json:"remote_ip"`
	TLS        bool     `json:"tls"`
	SPF        SPF      `json:"spf"`
}

// Headers contains all of the message headers extracted from the email.
type Headers struct {
	Received                []string `json:"received"`
	Date                    string   `json:"date"`
	From                    string   `json:"from"`
	To                      string   `json:"to"`
	MessageID               string   `json:"message_id"`
	Subject                 string   `json:"subject"`
	MimeVersion             string   `json:"mime_version"`
	ContentType             string   `json:"content_type"`
	ContentTransferEncoding string   `json:"content_transfer_encoding"`
	XOriginatingIP          string   `json:"x_originating_ip"`
	XDomainSigner           string   `json:"x_domain_signer"`
	DkimSignature           string   `json:"dkim_signature"`
}

// Attachments to the message
type Attachments struct {
	Content     string      `json:"content"`
	FileName    string      `json:"file_name"`
	ContentType string      `json:"content_type"`
	Size        string      `json:"size"`
	Disposition string      `json:"disposition"`
	ContentID   interface{} `json:"content_id"`
}

// A Data struct contains all the fields of a decoded message
type Data struct {
	Headers     Headers       `json:"headers"`
	Envelope    Envelope      `json:"envelope"`
	Plain       *string       `json:"plain"`
	HTML        *string       `json:"html"`
	ReplyPlain  *string       `json:"reply_plain"`
	Attachments []Attachments `json:"attachments"`
}

// Decode a message into a Data struct
func Decode(r io.Reader) (Data, error) {
	var ret Data

	err := json.NewDecoder(r).Decode(&ret)

	return ret, err
}
