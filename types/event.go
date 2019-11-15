package types

import "encoding/xml"

// Events struct - events xml interface
type Events struct {
	XMLName   xml.Name   `xml:"events"`
	Listeners []Listener `xml:"listener"`
}

// Listener struct - listener xml interface
type Listener struct {
	XMLName xml.Name `xml:"listener"`
	For     string   `xml:"for,attr"`
	Call    string   `xml:"call,attr"`
	Name    string   `xml:"name,attr"`
}

// EventFunc type of events func
type EventFunc func(data ...interface{})
