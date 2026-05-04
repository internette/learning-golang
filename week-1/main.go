package main

import (
	"flag"
	"fmt"
	"os"
)

func addTask(taskName string) {
	var file, err = os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// Write a string to the file
	_, err = file.WriteString(taskName)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data written successfully.")
}

func main() {
	task := flag.String("taskName", "task name", "Name of the task")
	flag.Parse()

	addTask(*task)
}
