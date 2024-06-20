package main

import (
	"bytes"
	"strings"
)

func PrintAscii(input string, style string) string {
	var result string

	template := getContent(style)
	template = bytes.ReplaceAll(template, []byte{13, 10}, []byte{10})
	contentTable := bytes.Split(template, []byte{10})
	textToTransform := strings.Split(input, `\n`)
	for k := 0; k < len(textToTransform); k++ {
		lines := make([]string, 8)
		for i := 0; i < 8; i++ {
			for j := 0; j < len(textToTransform[k]); j++ {
				startChar := ((int(textToTransform[k][j]) - 32) * 9) + 1
				lines[i] += string(contentTable[startChar+i])
			}
		}
		if lines[0] != "" {
			result += strings.Join(lines, "\n")
		}
		if k != len(textToTransform) {
			result += "\n"
		}
	}
	return result
}
