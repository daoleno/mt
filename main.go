package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
)

const defaultDirPath = ".mt/data"

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
		files, err := lsfile()
		for _, f := range files {
			fmt.Println(f.Name(), f.ModTime().Format("2006-01-02 15:04:05"))
		}
		if err != nil {
			panic(err)
		}
	case "open":
		filename := os.Args[2]
		err := openfile(filename)
		if err != nil {
			panic(err)
		}
	case "encrypt":
		key := os.Args[2]
		err := encryptfile(key)
		if err != nil {
			panic(err)
		}
	case "decrypt":
		key := os.Args[2]
		err := decryptfile(key)
		if err != nil {
			panic(err)
		}
	}

}

func defaultDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir + "/" + defaultDirPath
}

func newfile(name string) error {
	// Check if directory exist and create if does not exist
	mkDir(defaultDir())

	// Open file with vim
	cmd := exec.Command("vim", defaultDir()+"/"+name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func lsfile() ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(defaultDir())
	if err != nil {
		return nil, err
	}
	return files, nil
}

func openfile(name string) error {
	err := newfile(name)
	if err != nil {
		return err
	}
	return nil
}

func encryptfile(key string) error {
	files, err := lsfile()
	if err != nil {
		return err
	}

	for _, fileinfo := range files {
		if fileinfo.Mode().IsRegular() {
			plaintext, err := ioutil.ReadFile(defaultDir() + "/" + fileinfo.Name())
			if err != nil {
				return err
			}
			encryptedText, err := encrypt(plaintext, []byte(key))
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(defaultDir()+"/"+fileinfo.Name(), encryptedText, 0644)
			if err != nil {
				return nil
			}
		}
	}
	return nil
}

func decryptfile(key string) error {
	files, err := lsfile()
	if err != nil {
		return err
	}

	for _, fileinfo := range files {
		if fileinfo.Mode().IsRegular() {
			plaintext, err := ioutil.ReadFile(defaultDir() + "/" + fileinfo.Name())
			if err != nil {
				return err
			}
			decryptedText, err := decrypt(plaintext, []byte(key))
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(defaultDir()+"/"+fileinfo.Name(), decryptedText, 0644)
			if err != nil {
				return nil
			}
		}
	}
	return nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	parseCmd()
}
