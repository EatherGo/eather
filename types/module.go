package types

import (
	"encoding/xml"
)

// ConfigDir the directory for config files
const ConfigDir = "./config/"

// ModuleXML struct - xml interface
type ModuleXML struct {
	XMLName      xml.Name     `xml:"module"`
	Name         string       `xml:"name"`
	Version      string       `xml:"version"`
	Events       Events       `xml:"events"`
	Dependencies Dependencies `xml:"dependencies"`
	Dir          string
}

// ModulesConfigXML structure for XML of modules config
type ModulesConfigXML struct {
	XMLName xml.Name       `xml:"modules"`
	Modules []ModuleConfig `xml:"module"`
}

// ModuleConfig structure for specific module
type ModuleConfig struct {
	XMLName xml.Name `xml:"module"`
	Name    string   `xml:"name"`
	Enabled bool     `xml:"enabled"`
}

// Module interface
type Module interface{}

type Eventable interface {
	GetEventFuncs() map[string]EventFunc
}

type Installable interface {
	Install()
}

type Upgradable interface {
	Upgrade(version string)
}

type Routable interface {
	MapRoutes()
}
