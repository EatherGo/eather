package types

import "encoding/xml"

// Dependencies struct - dependencies xml interface
type Dependencies struct {
	XMLName      xml.Name `xml:"dependencies"`
	Dependencies []string `xml:"dependency"`
}
