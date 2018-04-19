package model

import "encoding/xml"

type XMLResult struct {
	XMLName xml.Name `xml:"mediawiki"`
	Page    []Page   `xml:"page"`
}
