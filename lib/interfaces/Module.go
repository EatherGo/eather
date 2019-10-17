package interfaces

import (
	"encoding/xml"
)

type module string

// ModulesDir set the directory for modules
const ModulesDir = "./src/Modules"

// ModuleXML struct - xml interface
type ModuleXML struct {
	XMLName      xml.Name     `xml:"module"`
	Name         string       `xml:"name"`
	Path         string       `xml:"path"`
	Func         string       `xml:"func"`
	Version      string       `xml:"version"`
	Events       Events       `xml:"events"`
	Dependencies Dependencies `xml:"dependencies"`
}

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

type Dependencies struct {
	XMLName      xml.Name `xml:"dependencies"`
	Dependencies []string `xml:"dependency"`
}

// Module interface
type Module interface {
	GetEventFuncs() map[string]func()
	Install()
	Upgrade(version string)
	MapRoutes()
}
