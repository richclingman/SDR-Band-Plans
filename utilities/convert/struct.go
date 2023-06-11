package main

import "encoding/xml"

type XmlRangeEntry struct {
	XMLName      xml.Name `xml:"RangeEntry"`
	RangeEntry   string   `xml:"RangeEntry"`
	MinFrequency int      `xml:"MinFrequency,attr"`
	MaxFrequency int      `xml:"maxFrequency,attr"`
	Mode         string   `xml:"mode,attr"`
	Step         string   `xml:"step,attr"`
	Color        string   `xml:"color,attr"`
}

type XmlArrayOfRangeEntry struct {
	XMLName           xml.Name `xml:"ArrayOfRangeEntry"`
	ArrayOfRangeEntry []XmlRangeEntry
}
