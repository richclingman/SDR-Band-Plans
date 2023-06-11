package main

import (
	"encoding/json"
	"os"
)

type JsonRanges struct {
	Name        string      `json:"name"`
	CountryName string      `json:"country_name"`
	CountryCode string      `json:"country_code"`
	AuthorName  string      `json:"author_name"`
	AuthorUrl   string      `json:"author_url"`
	Ranges      []JsonRange `json:"bands"`
}

type JsonRange struct {
	Description  string `json:"name"`
	Color        string `json:"type"`
	MinFrequency int    `json:"start"`
	MaxFrequency int    `json:"end"`
	Mode         string `json:"mode,omitempty"`
	Step         string `json:"step,omitempty"`
}

func NewJsonRanges() JsonRanges {
	jsonRanges := JsonRanges{
		Name:        "USA (KN1E)",
		CountryName: "United States of America",
		CountryCode: "US",
		AuthorName:  "Arrin Clark - KN1E",
		AuthorUrl:   "https://github.com/Arrin-KN1E",
		Ranges:      []JsonRange{},
	}

	return jsonRanges
}

func (r JsonRanges) Write(jsonfile string) error {
	bytes, err := json.Marshal(r)
	if err != nil {
		return err
	}

	return os.WriteFile(jsonfile, bytes, os.ModePerm)
}
