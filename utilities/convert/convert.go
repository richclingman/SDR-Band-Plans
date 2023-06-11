package main

import (
	"flag"
	"fmt"
	"os"
)

// TODO - convert the colors using "Config BandColors Modification.txt"
// Example: 33FFC800": "#FFC80033"
// C# has ALPHA first. Simply MOVE first 2 bytes to end, prepend "#"
// TODO: Can we determine and use "amateur" and other SDR++ predefined colors?

// TODO - In json, set name, country_name, country_code

func main() {
	from := flag.String("from", "", "File converting from")
	to := flag.String("to", "", "File converting to")
	flag.Parse()

	// TODO: Do we want to keep these empty bands?

	if *from == "" || *to == "" {
		fmt.Print("\n\nUsage:\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if (*from)[len(*from)-4:] != ".xml" || (*to)[len(*to)-5:] != ".json" {
		fmt.Printf("F: '%s' T: '%s'\n", (*from)[len(*from)-4:], (*to)[len(*to)-5:])
		fmt.Print("\nCurrently, only support conversion from '.xml' to '.json'.\n\n")
		os.Exit(2)
	}

	fmt.Printf("\nConverting '%s' to '%s'...", *from, *to)

	xmlRanges, _ := ReadXml(*from)
	jsonRanges := NewJsonRanges()

	for i, r := range xmlRanges.Ranges {
		fmt.Printf("%3d: %+v\n", i, r)

		if r.Description == "" {
			continue
		}

		// TODO: convert xml to json
		jsonRange := JsonRange{
			Description:  r.Description,
			Color:        r.Color,
			MinFrequency: r.MinFrequency,
			MaxFrequency: r.MaxFrequency,
			Mode:         r.Mode,
			Step:         r.Step,
		}
		jsonRanges.Ranges = append(jsonRanges.Ranges, jsonRange)

		fmt.Printf("%3d: %+v\n", i, jsonRange)
	}

	// TODO: write json
	jsonRanges.Write(*to)
}
