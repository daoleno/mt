package main

import (
	"io/ioutil"
	"mt/vcs"
	"os"
	"os/exec"
	"strings"
)

func newFile(name string) error {
	// Check if directory exist and create if does not exist
	mkDir(dataDir())

	// Init git repo
	git := vcs.ByCmd("git")
	err := git.Init(dataDir())
	if err != nil {
		panic(err)
	}

	// Monitor git repo
	go Monitor(dataDir())

	// Open file with vim
	cmd := exec.Command("vim", dataDir()+"/"+name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func listFile(dir string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for idx, f := range files {
		if strings.Contains(f.Name(), ".git") {
			files = append(files[:idx], files[idx+1:]...)
		}
	}
	return files, nil
}

func listDataFile() ([]os.FileInfo, error) {
	files, err := listFile(dataDir())
	if err != nil {
		return nil, err
	}
	return files, nil
}

func listRenderFile() ([]os.FileInfo, error) {
	files, err := listFile(renderDir())
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

func catFile(name string) error {
	// Check if directory exist and create if does not exist
	mkDir(dataDir())

	// Exec cat command
	cmd := exec.Command("cat", dataDir()+"/"+name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func renameFile(oldName, newName string) error {
	oldName = dataDir() + "/" + oldName
	newName = dataDir() + "/" + newName
	err := os.Rename(oldName, newName)
	if err != nil {
		return err
	}
	return nil
}

func deleteFile(name string) error {
	go Monitor(dataDir())
	// Delete data file
	dataFiles, err := listDataFile()
	if err != nil {
		return err
	}
	for _, fileinfo := range dataFiles {
		if fileinfo.Name() == name {
			os.Remove(dataDir() + "/" + fileinfo.Name())
		}
	}

	// Delete render file
	renderFiles, err := listRenderFile()
	for _, fileinfo := range renderFiles {
		if fileinfo.Name() == name {
			os.Remove(renderDir() + "/" + fileinfo.Name())
		}
	}

	return nil
}

func deleteAllFile() error {
	dataFiles, err := listDataFile()
	if err != nil {
		return err
	}
	renderFiles, err := listRenderFile()
	if err != nil {
		return err
	}
	files := append(dataFiles, renderFiles...)
	for _, fileinfo := range files {
		err := deleteFile(fileinfo.Name())
		if err != nil {
			return err
		}
	}
	return nil
}

func encryptFile(key string) error {
	files, err := listDataFile()
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
	files, err := listDataFile()
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
