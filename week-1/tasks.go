package tasks

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Task struct {
	Name   string `json:"name"`
	Index  int    `json:"index"`
	Status string `json:"status"`
}

func addTask(tasks []Task, taskName string, taskStatus string) {
	index := len(tasks) + 1
	newTask := Task{Name: taskName, Index: index, Status: taskStatus}
	tasks = append(tasks, newTask)
	bts, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
	os.WriteFile("tasks.json", bts, 0644)

	fmt.Println("Data written successfully.")
}

func delTask(tasks []Task, taskIndex int) {
	tasks = append(tasks[:taskIndex-1], tasks[taskIndex:]...)
	bts, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
	os.WriteFile("tasks.json", bts, 0644)

	fmt.Println("Task deleted successfully.")
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
	fmt.Printf("%-5s %-25s %-20s\n", "Index", "Name", "Status")
	fmt.Printf("%-5s %-25s %-20s\n", "----", "----", "----")
	for _, task := range tasks {
		fmt.Printf("%-5d %-25s %-20s\n", task.Index, task.Name, task.Status)
	}
}

func listFilterTasks(tasks []Task, taskFilter string) {
	fmt.Printf("%-5s %-25s %-20s\n", "Index", "Name", "Status")
	fmt.Printf("%-5s %-25s %-20s\n", "----", "----", "----")
	for _, task := range tasks {
		if task.Status == taskFilter {
			fmt.Printf("%-5d %-25s %-20s\n", task.Index, task.Name, task.Status)
		}
	}
}

func main() {
	cmd := flag.String("cmd", "get", "Command to execute")
	task := flag.String("taskName", "task name", "Name of the task")
	taskIndex := flag.Int("taskIndex", 1, "Index of the task")
	taskStatus := flag.String("taskStatus", "", "Status of the task")
	tasks := getTasks()
	flag.Parse()
	switch *cmd {
	case "get":
		if *taskStatus != "" {
			listFilterTasks(tasks, *taskStatus)
		} else {
			listTasks(tasks)
		}
	case "add":
		addTask(tasks, *task, *taskStatus)
		listTasks(tasks)
	case "del":
		delTask(tasks, *taskIndex)
		listTasks(tasks)
	default:
		if *taskStatus != "" {
			listFilterTasks(tasks, *taskStatus)
		} else {
			listTasks(tasks)
		}
	}
}
