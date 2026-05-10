package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Folder struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func createLogbook() {
	entries, err := os.ReadDir("../")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	var projectFolders []Folder
	for _, entry := range entries {
		if entry.IsDir() && strings.Contains(entry.Name(), "week-") {
			var formattedName = strings.ReplaceAll(entry.Name(), "-", " ")
			formattedName = strings.ReplaceAll(formattedName, "w", "W")
			projectFolders = append(projectFolders, Folder{Name: formattedName, Path: entry.Name()})
		}
	}
	bts, err := json.Marshal(projectFolders)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
	os.WriteFile("logbook.json", bts, 0644)
}

func main() {
	createLogbook()
}
