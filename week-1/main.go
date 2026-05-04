package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Task struct {
	Name string `json:"name"`
}

func addTask(tasks []Task, taskName string) {
	tasks = append(tasks, Task{Name: taskName})
	bts, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	os.WriteFile("tasks.json", bts, 0644)

	fmt.Println("Data written successfully.")
}

func main() {
	task := flag.String("taskName", "task name", "Name of the task")
	flag.Parse()
	var dataSlice = make([]Task, 0)

	addTask(dataSlice, *task)
}
