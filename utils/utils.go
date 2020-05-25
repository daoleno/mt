package utils

import (
	"log"
	"os"
	"os/user"
)

// Check if directory exist and create if does not exist
func MkDir(dir string) {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dir, 0755)
		if errDir != nil {
			panic(err)
		}
	}
}

func homeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func DataDir() string {
	return homeDir() + "/" + ".mt/data"
}

func RenderDir() string {
	return homeDir() + "/" + ".mt/render"
}
