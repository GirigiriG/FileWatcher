package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

func main() {
	// type MyDate time.Time
	filePath := "/Users/gideongirigiri/golang/fileWatcher/src/main.go"
	var processIDFile = "process.txt"
	fileUpdateMap := make(map[string]time.Time)

	for {

		time.Sleep(1 * time.Second)

		info, _ := os.Stat(filePath)

		if fileUpdateMap[filePath] != info.ModTime() {
			fileUpdateMap[filePath] = info.ModTime()
			
			b, _ := ioutil.ReadFile(processIDFile)
			pid := string(b)
			if len(b) != 0 {

				var killCommand = []string{"kill", "-9", pid}
				cmd := exec.Command(killCommand[0], killCommand...)

				cmd.Run()
				println(fileUpdateMap[filePath] != info.ModTime())
			}

			execCommand := []string{"go", "run", filePath}
			cmd := exec.Command(execCommand[0], execCommand...)
			println("started program on process: ", pid)
			cmd.Run()
		}

	}
}
