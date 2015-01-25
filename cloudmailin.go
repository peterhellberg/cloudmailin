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
	From       string   `json:"From"`
	HeloDomain string   `json:"helo_domain"`
	RemoteIP   string   `json:"remote_ip"`
	SPF        SPF      `json:"spf"`
}

// Headers contains all of the message headers extracted from the email.
type Headers struct {
	ReturnPath            string   `json:"Return-Path"`
	Received              []string `json:"Received"`
	Date                  string   `json:"Date"`
	From                  string   `json:"From"`
	To                    string   `json:"To"`
	Bcc                   string   `json:"Bcc"`
	MessageID             string   `json:"Message-ID"`
	Subject               string   `json:"Subject"`
	MimeVersion           string   `json:"Mime-Version"`
	ContentType           string   `json:"Content-Type"`
	DeliveredTo           string   `json:"Delivered-To"`
	ReceivedSPF           string   `json:"Received-SPF"`
	AuthenticationResults string   `json:"Authentication-Results"`
	UserAgent             string   `json:"User-Agent"`
}

// Attachments to the message
type Attachments struct {
	Content     string `json:"content"`
	URL         string `json:"url"`
	FileName    string `json:"file_name"`
	ContentType string `json:"content-type"`
	Size        int    `json:"size"`
	Disposition string `json:"disposition"`
}

// A Data struct contains all the fields of a decoded message
type Data struct {
	Headers     Headers       `json:"headers"`
	Envelope    Envelope      `json:"envelope"`
	Plain       string        `json:"plain"`
	HTML        string        `json:"html"`
	ReplyPlain  string        `json:"reply_plain"`
	Attachments []Attachments `json:"attachments"`
}

// Decode a message into a Data struct
func Decode(r io.ReadCloser) (Data, error) {
	var ret Data

	err := json.NewDecoder(r).Decode(&ret)

	return ret, err
}
