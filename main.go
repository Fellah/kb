package main

// Project packages.
import (
	"kb/web"
)

// Third side packages.
import (
	"gopkg.in/fsnotify.v1"
)

// Base packages.
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
	web.Serve()
	watchDataDir()
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
