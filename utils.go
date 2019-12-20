package main

import "os"

// Check if directory exist and create if does not exist
func mkDir(defaultDir string) {
	_, err := os.Stat(defaultDir)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(defaultDir, 0755)
		if errDir != nil {
			panic(err)
		}
	}
}
