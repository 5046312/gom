package run

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/5046312/gom/util"

	"github.com/fsnotify/fsnotify"
)

var (
	cmd  *exec.Cmd
	dir  string
	file string
)

func watcher(dir string, file string) {
	path := filepath.Join(dir, file)
	exist := util.PathExists(path)
	if !exist {
		log.Println("Main File Not Exist:" + path)
		return
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok || cmd.Process == nil {
					return
				}
				log.Println("Event:", event)
				if kill() {
					run(dir, file)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Watcher Error:", err)
			}
		}
	}()

	// Start Process
	run(dir, file)

	err = watcher.Add(dir)
	if err != nil {
		log.Fatal("Watcher error:", err)
	}
	log.Println("Watcher:", dir)

	done := make(chan bool)
	<-done
}

func kill() bool {
	// Have Process
	if cmd != nil && cmd.Process != nil {
		log.Println("Kill Process:", cmd.Process.Pid)
		// if runtime.GOOS == "windows" {
		// 	cmd.Process.Signal(os.Kill)
		// } else {
		// 	cmd.Process.Signal(os.Interrupt)
		// }
		// cmd.Wait()
		err := cmd.Process.Kill()
		if err != nil {
			log.Printf("Error while killing cmd process: %s", err)
			return false
		} else {
			return true
		}
	}
	return false
}

// go run ...
func run(dir string, file string) {
	path := filepath.Join(dir, file)
	log.Println("Building Project:", path)
	bcmd := exec.Command("go", "build", path)
	bcmd.Stderr = os.Stderr
	bcmd.Stdout = os.Stdout
	err := bcmd.Run()
	if err != nil {
		log.Println("Failed To Build the Project:", err)
		os.Exit(2)
		return
	}

	exeFile := strings.Replace(file, `.go`, ``, -1)
	if runtime.GOOS == "windows" {
		exeFile += ".exe"
	}

	cmd = exec.Command(exeFile)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	log.Println("Start Process:", exeFile)
	go func() {
		err := cmd.Run()
		if err != nil {
			log.Println("Process Run Error: ", err)
		}
	}()
}
