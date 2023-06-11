package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func ReadXml(xmlfile string) (XmlArrayOfRangeEntry, error) {
	bytes, err := os.ReadFile(xmlfile)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(bytes)[:1000], "\n\n")

	var ranges XmlArrayOfRangeEntry
	err = xml.Unmarshal(bytes, &ranges)
	if err != nil {
		panic(err)
	}

	return ranges, err
}
