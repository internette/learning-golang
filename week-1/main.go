package main

import (
	"flag"
	"fmt"
)

func main() {
	task := flag.String("taskName", "task name", "Name of the task")
	flag.Parse()
	fmt.Printf("- %s\n", *task)
}
