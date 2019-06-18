package run

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

var restarting bool = false
var (
	cmd *exec.Cmd
)

func watcher(dir string, file string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// TODO CHANEL
				if restarting {
					return
				}
				log.Println("Event", event)
				restarting = true
				kill()
				run(dir, file)
				restarting = false

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Watcher:", dir)

	// Start Process
	run(dir, file)

	done := make(chan bool)
	<-done
}

func kill() {
	// Have Process
	if cmd != nil && cmd.Process != nil {
		cmd.Process.Kill()
	}
}

// go run ...
func run(dir string, file string) {
	path := filepath.Join(dir, file)
	cmd := exec.Command("go", "run", path)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	log.Println("Starting", path)
	go cmd.Run()
}
