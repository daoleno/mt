package main

import (
	"gopkg.in/russross/blackfriday.v2"
)

// Render markdown to html
func render(input []byte) (output []byte) {
	return blackfriday.Run(input)
}
