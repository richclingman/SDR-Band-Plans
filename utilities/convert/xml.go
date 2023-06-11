package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type XmlRanges struct {
	XMLName xml.Name   `xml:"ArrayOfRangeEntry"`
	Ranges  []XmlRange `xml:"RangeEntry"`
}

type XmlRange struct {
	XMLName      xml.Name `xml:"RangeEntry"`
	Description  string   `xml:",chardata"`
	MinFrequency int      `xml:"minFrequency,attr"`
	MaxFrequency int      `xml:"maxFrequency,attr"`
	Mode         string   `xml:"mode,attr"`
	Step         string   `xml:"step,attr"`
	Color        string   `xml:"color,attr"`
}

func ReadXml(xmlfile string) (XmlRanges, error) {
	bytes, err := os.ReadFile(xmlfile)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(bytes)[:1000], "\n\n")

	var ranges XmlRanges
	err = xml.Unmarshal(bytes, &ranges)
	if err != nil {
		panic(err)
	}

	return ranges, err
}
