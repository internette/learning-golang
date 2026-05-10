package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Folder struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

var rootCmd = &cobra.Command{
	Use:   "algorithmsLogbook",
	Short: "Create a logbook of algorithm projects",
	Long:  `A CLI tool to generate a JSON logbook of weekly algorithm projects.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is a simple CLI tool to create a logbook of algorithm projects and track them.")
	},
}

func createLogbook() error {
	entries, err := os.ReadDir("../")
	if err != nil {
		return fmt.Errorf("error reading directory: %w", err)
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
		return fmt.Errorf("error marshaling to JSON: %w", err)
	}
	err = os.WriteFile("logbook.json", bts, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}
	return nil
}

var createLogbookCmd = &cobra.Command{
	Use:   "create-logbook",
	Short: "Create a logbook of algorithm projects",
	Long:  "Create a logbook of algorithm projects and track them.",
	Run: func(cmd *cobra.Command, args []string) {
		err := createLogbook()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("Logbook created successfully.")
	},
}

func init() {
	rootCmd.AddCommand(createLogbookCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
