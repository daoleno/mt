package utils

import "github.com/russross/blackfriday"

// Render markdown to html
func Render(input []byte) (output []byte) {
	return blackfriday.Run(input)
}
