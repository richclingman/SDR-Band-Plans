package main

import (
	"encoding/xml"
)

type XmlArrayOfRangeEntry struct {
	XMLName xml.Name        `xml:"ArrayOfRangeEntry"`
	Ranges  []XmlRangeEntry `xml:"RangeEntry"`
}

type XmlRangeEntry struct {
	XMLName      xml.Name `xml:"RangeEntry"`
	Description  string   `xml:",chardata"`
	MinFrequency int      `xml:"minFrequency,attr"`
	MaxFrequency int      `xml:"maxFrequency,attr"`
	Mode         string   `xml:"mode,attr"`
	Step         string   `xml:"step,attr"`
	Color        string   `xml:"color,attr"`
}
