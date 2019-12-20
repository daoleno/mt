package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

const defaultDir = "data"

func parseCmd() {
	if len(os.Args) == 1 {
		fmt.Println("My Thoughs")
		return
	}
	switch os.Args[1] {
	case "new":
		filename := os.Args[2]
		err := newfile(filename)
		if err != nil {
			panic(err)
		}
	case "ls":
		err := lsfile()
		if err != nil {
			panic(err)
		}
	case "open":
		filename := os.Args[2]
		err := openfile(filename)
		if err != nil {
			panic(err)
		}
	}

}

func newfile(name string) error {
	// Check if directory exist and create if does not exist
	mkDir(defaultDir)

	// Open file with vim
	cmd := exec.Command("vim", defaultDir+"/"+name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func lsfile() error {
	files, err := ioutil.ReadDir(defaultDir)
	if err != nil {
		return err
	}
	for _, f := range files {
		fmt.Println(f.Name(), f.ModTime().Format("2006-01-02 15:04:05"))
	}
	return nil
}

func openfile(name string) error {
	err := newfile(name)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	parseCmd()
}
