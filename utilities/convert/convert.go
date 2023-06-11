package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	from := flag.String("from", "", "File converting from")
	to := flag.String("to", "", "File converting to")
	flag.Parse()

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

	ranges, _ := ReadXml(*from)

	for i, r := range ranges.Ranges {
		fmt.Printf("%3d: %+v\n", i, r)

		// TODO: convert xml to json

	}

	// TODO: write json

}
