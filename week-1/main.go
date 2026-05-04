package main

import (
	"flag"
	"fmt"
)

func addTask(taskName string) {
	fmt.Printf("- %s\n", taskName)
}

func main() {
	task := flag.String("taskName", "task name", "Name of the task")
	flag.Parse()

	addTask(*task)
}
