package main

import (
	// Project packages.
	"github.com/fellah/kb/markdown"
	"github.com/fellah/kb/web"

	// Third side packages.
	"github.com/fellah/version"
	"gopkg.in/fsnotify.v1"

	// Base packages.
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	// Default content directory.
	DATA_DIR string = "."
)

var dataDir string = DATA_DIR

var (
	flgVer *bool
)

func main() {
	parseCliArgs()

	if *flgVer {
		fmt.Println("Version:", version.GetVersion())
		fmt.Println("DateTime:", version.GetDateTime())
		fmt.Println("Commit:", version.GetCommit())
		fmt.Println("Branch:", version.GetBranch())
		fmt.Println("Author:", version.GetAuthor())
		return
	}

	log.Println("Parse markdown data...")
	markdown.SetBaseDir(dataDir)
	filepath.Walk(dataDir, markdown.Walk)

	web.Serve()
	watchDataDir()
}

// Parse command line arguments.
func parseCliArgs() {
	flgVer = flag.Bool("v", false, "Output version information.")

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

// Monitoring data directory for changes.
func watchDataDir() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalln(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(dataDir)
	if err != nil {
		log.Fatalln(err)
	}

	<-done
}
