package main

import (
	"flag"
	"log"
	"os"
)

const (
	// Default content directory.
	DATA_DIR string = "."
)

var dataDir string = DATA_DIR

func main() {
	parseCliArgs()
}

// Parse command line arguments.
func parseCliArgs() {
	flag.Parse()
	args := flag.Args()

	// Verify data directory.
	if len(args) > 0 {
		dir, err := os.Open(args[len(args)-1])
		if err != nil {
			log.Fatalln(err)
		}
		defer dir.Close()

		di, err := dir.Stat()
		if err != nil {
			log.Fatalln(err)
		}

		if !di.IsDir() {
			log.Fatalln(err)
		}

		dataDir = args[len(args)-1]
	}
}
