package types

import (
	"encoding/xml"
)

// ModulesDir set the directory for modules
const ModulesDir = "./src/Modules"

// ModuleXML struct - xml interface
type ModuleXML struct {
	XMLName      xml.Name     `xml:"module"`
	Name         string       `xml:"name"`
	Func         string       `xml:"func"`
	Version      string       `xml:"version"`
	Events       Events       `xml:"events"`
	Dependencies Dependencies `xml:"dependencies"`
}

type ModulesConfigXML struct {
	XMLName xml.Name       `xml:"modules"`
	Modules []ModuleConfig `xml:"module"`
}

type ModuleConfig struct {
	XMLName xml.Name `xml:"module"`
	Name    string   `xml:"name"`
	Enabled bool     `xml:"enabled"`
}

// Module interface
type Module interface {
	GetEventFuncs() map[string]EventFunc
	Install()
	Upgrade(version string)
	MapRoutes()
}
