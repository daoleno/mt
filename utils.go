package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/urfave/cli/v2"
)

// Check if directory exist and create if does not exist
func mkDir(dir string) {
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

func dataDir() string {
	return homeDir() + "/" + ".mt/data"
}

func renderDir() string {
	return homeDir() + "/" + ".mt/render"
}

func bashCompleteFile(c *cli.Context) {
	if c.NArg() > 0 {
		return
	}
	files, err := listDataFile()
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
