package main

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func newFile(name string) error {
	// Check if directory exist and create if does not exist
	mkDir(dataDir())

	// Open file with vim
	cmd := exec.Command("vim", dataDir()+"/"+name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func listFile() ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dataDir())
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
	files, err := listFile()
	if err != nil {
		return err
	}
	for _, fileinfo := range files {
		if fileinfo.Name() == name {
			os.Remove(dataDir() + "/" + fileinfo.Name())
		}
	}
	return nil
}

func deleteAllFile() error {
	files, err := listFile()
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
	files, err := listFile()
	if err != nil {
		return err
	}

	for _, fileinfo := range files {
		if fileinfo.Mode().IsRegular() {
			plaintext, err := ioutil.ReadFile(dataDir() + "/" + fileinfo.Name())
			if err != nil {
				return err
			}
			encryptedText, err := encrypt(plaintext, []byte(key))
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(dataDir()+"/"+fileinfo.Name(), encryptedText, 0644)
			if err != nil {
				return nil
			}
		}
	}
	return nil
}

func decryptFile(key string) error {
	files, err := listFile()
	if err != nil {
		return err
	}

	for _, fileinfo := range files {
		if fileinfo.Mode().IsRegular() {
			plaintext, err := ioutil.ReadFile(dataDir() + "/" + fileinfo.Name())
			if err != nil {
				return err
			}
			decryptedText, err := decrypt(plaintext, []byte(key))
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(dataDir()+"/"+fileinfo.Name(), decryptedText, 0644)
			if err != nil {
				return nil
			}
		}
	}
	return nil
}
