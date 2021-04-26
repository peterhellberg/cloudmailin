package cloudmailin

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
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

// Headers contains all of the message headers extracted from the email.
type HeadersStringReceived struct {
	Received string `json:"received"`
	Headers
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

type DataHeadersStringReceived struct {
	Headers HeadersStringReceived `json:"headers"`
	Data
}

// Decode a message into a Data struct
func Decode(r io.Reader) (Data, error) {
	var ret Data
	var tmp DataHeadersStringReceived

	body, _ := ioutil.ReadAll(r)

	reader1 := bytes.NewReader(body)
	reader2 := bytes.NewReader(body)

	err := json.NewDecoder(reader1).Decode(&ret)
	if err != nil {
		tmpErr := json.NewDecoder(reader2).Decode(&tmp)
		if tmpErr != nil {
			return ret, tmpErr
		}

		ret.Headers.Received = []string{tmp.Headers.Received}
		ret.Headers.Date = tmp.Headers.Date
		ret.Headers.From = tmp.Headers.From
		ret.Headers.To = tmp.Headers.To
		ret.Headers.MessageID = tmp.Headers.MessageID
		ret.Headers.Subject = tmp.Headers.Subject
		ret.Headers.MimeVersion = tmp.Headers.MimeVersion
		ret.Headers.ContentType = tmp.Headers.ContentType
		ret.Headers.ContentTransferEncoding = tmp.Headers.ContentTransferEncoding
		ret.Headers.XOriginatingIP = tmp.Headers.XOriginatingIP
		ret.Headers.XDomainSigner = tmp.Headers.XDomainSigner
		ret.Headers.DkimSignature = tmp.Headers.DkimSignature

		return ret, nil
	}

	return ret, err
}
