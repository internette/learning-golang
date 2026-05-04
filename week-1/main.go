package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Task struct {
	Name  string `json:"name"`
	Index int    `json:"index"`
}

func addTask(tasks []Task, taskName string) {
	index := len(tasks) + 1
	newTask := Task{Name: taskName, Index: index}
	tasks = append(tasks, newTask)
	bts, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
	os.WriteFile("tasks.json", bts, 0644)

	fmt.Println("Data written successfully.")
}

func getTasks() []Task {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("No file found. Starting with an empty list.")
		return []Task{}
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return []Task{}
	}
	return tasks
}

func listTasks(tasks []Task) {
	for i := 0; i < len(tasks); i++ {
		task := tasks[i]
		fmt.Printf("- %s\n", task.Name)
	}
}

func main() {
	cmd := flag.String("cmd", "get", "Command to execute")
	task := flag.String("taskName", "task name", "Name of the task")
	tasks := getTasks()
	flag.Parse()
	switch *cmd {
	case "get":
		listTasks(tasks)
	case "add":
		addTask(tasks, *task)
	default:
		listTasks(tasks)
	}
}
