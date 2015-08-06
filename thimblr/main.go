package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/buddhamagnet/gothimbl"
)

var command string

func usage() {
	fmt.Printf("usage: %s %s\n", filepath.Base(os.Args[0]), gothimbl.ListCommands())
	os.Exit(1)
}

func usageSetup() {
	fmt.Printf("usage: %s setup <thimbl_user>\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}

func main() {

	if len(os.Args) == 1 {
		usage()
	}
	command = os.Args[1]
	switch command {
	case "setup":
		if len(os.Args) != 3 {
			usageSetup()
		}
		gothimbl.Setup(os.Args[2])
	case "version":
		fmt.Println(gothimbl.Version())
	default:
		usage()
	}
}
