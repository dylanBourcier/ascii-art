package main

import (
	"fmt"
	"os"
)

func getContent(fileName string) []byte {
	content, err := os.ReadFile("templates/" + fileName + ".txt")
	if err != nil {
		fmt.Println("Template file not found. Exiting.")
		os.Exit(0)
	}
	return content
}
