package utils

import (
	"log"
	"os"
	"os/user"
)

// MkDir Check if directory exist and create if does not exist
func MkDir(dir string) {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dir, 0755)
		if errDir != nil {
			panic(err)
		}
	}
}

// AddGitIgnore add a .gitignore file in data dir
func AddGitIgnore() {
	ignoreFile := DataDir() + "/.gitignore"
	_, err := os.Stat(ignoreFile)
	if err == nil {
		return
	}

	f, err := os.OpenFile(ignoreFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	f.WriteString(".*")
}

// TODO: Replace with os.ExpandEnv()
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
