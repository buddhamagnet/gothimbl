package gothimbl

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
)

var ThimblFolder string

var commands = []string{"setup", "version", "print", "post", "follow"}

func init() {
	home, _ := homedir.Dir()
	ThimblFolder = home + "/.thimbl"
}

func ListCommands() string {
	return strings.Join(commands, "|")
}

func Setup(user string) {
	if _, err := os.Stat(ThimblFolder); os.IsNotExist(err) {
		os.Mkdir(ThimblFolder, 0777)
	}
	file, err := os.Create(ThimblFolder + "/" + "actual")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(user)
	if err != nil {
		log.Fatal(err)
	}
	writer.Flush()
}

func Version() string {
	return "0.1.0"
}
