package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/urfave/cli/v2"
)

func parseCmd() {
	app := &cli.App{
		Name:  "My Thought",
		Usage: "Rocord all my thoughts",
		Commands: []*cli.Command{
			{
				Name:  "open",
				Usage: "Open a thought",
				Action: func(c *cli.Context) error {
					filename := c.Args().First()
					// log.Println("Open thought: ", filename)
					err := openFile(filename)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:  "cat",
				Usage: "View a thought",
				Action: func(c *cli.Context) error {
					filename := c.Args().First()
					err := catFile(filename)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "Delete a thought",
				Action: func(c *cli.Context) error {
					filename := c.Args().First()
					// log.Println("Delete thought: ", filename)
					err := deleteFile(filename)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:  "list",
				Usage: "List all thoughts",
				Action: func(c *cli.Context) error {
					// log.Println("List all thoughts")
					files, err := listDataFile()
					if err != nil {
						return err
					}
					// Sort file by file modtime
					sort.Slice(files, func(i, j int) bool {
						return files[i].ModTime().After(files[j].ModTime())
					})

					// Print file
					for _, f := range files {
						fmt.Println(f.Name(), f.ModTime().Format("2006-01-02 15:04:05"))
					}
					return nil
				},
			},
			{
				Name:  "clean",
				Usage: "Clean all thoughts",
				Action: func(c *cli.Context) error {
					// log.Println("Clean all thoughts")
					err := deleteAllFile()
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:  "encrypt",
				Usage: "Encrypt all thoughts",
				Action: func(c *cli.Context) error {
					key := c.Args().First()
					err := encryptFile(key)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:  "decrypt",
				Usage: "Decrypt all thoughts",
				Action: func(c *cli.Context) error {
					key := c.Args().First()
					err := decryptFile(key)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:  "render",
				Usage: "Render all markdown to beautiful html",
				Action: func(c *cli.Context) error {
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
							htmltext := render(plaintext)
							mkDir(renderDir())
							err = ioutil.WriteFile(renderDir()+"/"+strings.TrimSuffix(fileinfo.Name(), path.Ext(fileinfo.Name()))+".html", htmltext, 0644)
							if err != nil {
								return err
							}
						}
					}
					fmt.Printf("All your thoughts is generated in %s\n", renderDir())
					return nil
				},
			},
			{
				Name:  "web",
				Usage: "Server static html page",
				Action: func(c *cli.Context) error {
					serveStatic()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	parseCmd()
}
