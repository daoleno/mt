package main

import (
	"fmt"
	"log"
	"os"

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
					files, err := lsFile()
					if err != nil {
						return err
					}
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
					err := encryptFile(key)
					if err != nil {
						return err
					}
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
