package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
)

const defaultDirPath = ".mt/data"

func defaultDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir + "/" + defaultDirPath
}

func newFile(name string) error {
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

func lsFile() ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(defaultDir())
	if err != nil {
		return nil, err
	}
	return files, nil
}

func openFile(name string) error {
	err := newFile(name)
	if err != nil {
		return err
	}
	return nil
}

func deleteFile(name string) error {
	files, err := lsFile()
	if err != nil {
		return err
	}
	for _, fileinfo := range files {
		if fileinfo.Name() == name {
			os.Remove(defaultDir() + "/" + fileinfo.Name())
		}
	}
	return nil
}

func deleteAllFile() error {
	files, err := lsFile()
	if err != nil {
		return nil
	}
	for _, fileinfo := range files {
		err := deleteFile(fileinfo.Name())
		if err != nil {
			return err
		}
	}
	return nil
}

func encryptFile(key string) error {
	files, err := lsFile()
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

func decryptFile(key string) error {
	files, err := lsFile()
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
