package main

import (
	"fmt"
	"os"
	"strings"
)

func getProjects() {
	entries, err := os.ReadDir("../")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	var projectFolders []string
	for _, entry := range entries {
		if entry.IsDir() && strings.Contains(entry.Name(), "week-") {
			projectFolders = append(projectFolders, entry.Name())
		}
	}
	for _, folder := range projectFolders {
		fmt.Println(folder)
	}
}

func main() {
	getProjects()
}
