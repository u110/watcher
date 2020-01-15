package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func GetFileModTime(filepath string) time.Time {
	file, err := os.Open(filepath) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	// log.Println("target:", file.Name())
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	return fileInfo.ModTime()
}

func watch(target string, cmdStr string) {
	filename := target
	log.Println("target:", filename)
	last := GetFileModTime(filename)
	log.Println("last modified:", last)

	for {
		// TODO(u11): use parameter
		time.Sleep(time.Duration(5) * time.Second)
		current := GetFileModTime(filename)
		if current.After(last) {
			last = current

			cmds := strings.Split(cmdStr, " ")
			log.Println("Modified! cmd start:", cmds)

			cmd := exec.Command(cmds[0], cmds[1:]...)
			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Printf("%s\n", out)
		}
	}
}

func help() {
	fmt.Println(`NAME:
   watcher - watch file modified and execute something

USAGE:
   watcher {target-file} "{execute command}"

VERSION:
   0.0.1

GLOBAL OPTIONS:
   --help, -h     show help (default: false)`)
}

func main() {
	for _, v := range os.Args {
		if v == "--help" || v == "-h" {
			help()
			return
		}
	}
	if len(os.Args) < 2 {
		log.Fatalf("Error. args must be two. args: %s", os.Args)
		return
	}
	watch(os.Args[1], os.Args[2])
}
