package file

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/daoleno/mt/utils"
	"github.com/daoleno/mt/vcs"
)

func newFile(name string) error {
	// Check if directory exist and create if does not exist
	utils.MkDir(utils.DataDir())

	// Place a .gitignore file
	utils.AddGitIgnore()

	// Init git repo
	git := vcs.ByCmd("git")
	err := git.Init(utils.DataDir())
	if err != nil {
		panic(err)
	}

	// Monitor git repo
	go Monitor(utils.DataDir())

	// Open file with vim
	cmd := exec.Command("vim", utils.DataDir()+"/"+name)
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
	var filtedFiles []os.FileInfo
	for _, f := range files {
		if strings.HasPrefix(f.Name(), ".") {
			continue
		}
		filtedFiles = append(filtedFiles, f)
	}
	return filtedFiles, nil
}

func ListDataFile() ([]os.FileInfo, error) {
	files, err := listFile(utils.DataDir())
	if err != nil {
		return nil, err
	}
	return files, nil
}

func ListRenderFile() ([]os.FileInfo, error) {
	files, err := listFile(utils.RenderDir())
	if err != nil {
		return nil, err
	}
	return files, nil
}

func OpenFile(name string) error {
	err := newFile(name)
	if err != nil {
		return err
	}
	return nil
}

func CatFile(name string) error {
	// Check if directory exist and create if does not exist
	utils.MkDir(utils.DataDir())

	// Exec cat command
	cmd := exec.Command("cat", utils.DataDir()+"/"+name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func RenameFile(oldName, newName string) error {
	oldName = utils.DataDir() + "/" + oldName
	newName = utils.DataDir() + "/" + newName
	err := os.Rename(oldName, newName)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFile(name string) error {
	go Monitor(utils.DataDir())
	// Delete data file
	dataFiles, err := ListDataFile()
	if err != nil {
		return err
	}
	for _, fileinfo := range dataFiles {
		if fileinfo.Name() == name {
			os.Remove(utils.DataDir() + "/" + fileinfo.Name())
		}
	}

	// Delete render file
	renderFiles, err := ListRenderFile()
	for _, fileinfo := range renderFiles {
		if fileinfo.Name() == name {
			os.Remove(utils.RenderDir() + "/" + fileinfo.Name())
		}
	}

	return nil
}

func DeleteAllFile() error {
	dataFiles, err := ListDataFile()
	if err != nil {
		return err
	}
	renderFiles, err := ListRenderFile()
	if err != nil {
		return err
	}
	files := append(dataFiles, renderFiles...)
	for _, fileinfo := range files {
		err := DeleteFile(fileinfo.Name())
		if err != nil {
			return err
		}
	}
	return nil
}

func EncryptFile(key string) error {
	files, err := ListDataFile()
	if err != nil {
		return err
	}

	for _, fileinfo := range files {
		if fileinfo.Mode().IsRegular() {
			plaintext, err := ioutil.ReadFile(utils.DataDir() + "/" + fileinfo.Name())
			if err != nil {
				return err
			}
			encryptedText, err := Encrypt(plaintext, []byte(key))
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(utils.DataDir()+"/"+fileinfo.Name(), encryptedText, 0644)
			if err != nil {
				return nil
			}
		}
	}
	return nil
}

func DecryptFile(key string) error {
	files, err := ListDataFile()
	if err != nil {
		return err
	}

	for _, fileinfo := range files {
		if fileinfo.Mode().IsRegular() {
			plaintext, err := ioutil.ReadFile(utils.DataDir() + "/" + fileinfo.Name())
			if err != nil {
				return err
			}
			decryptedText, err := Decrypt(plaintext, []byte(key))
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(utils.DataDir()+"/"+fileinfo.Name(), decryptedText, 0644)
			if err != nil {
				return nil
			}
		}
	}
	return nil
}
