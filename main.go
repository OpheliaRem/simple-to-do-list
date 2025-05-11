package main

import (
	"fmt"
	"toDoList/activator"
	"toDoList/task"
	"toDoList/userChoice"
)

func main() {
	var tasks []task.Task

	options := []string{
		"1. List all tasks",
		"2. Add task",
		"3. Select task",
		"4. Quit",
	}

	actions := []func(*[]task.Task){
		listTasks,
		addTask,
		selectTask,
	}

	var performer userChoice.Performer = &userChoice.ConsolePerformer{Options: options, Activator: &activator.MultiTaskActivator{
		Tasks: &tasks, Actions: actions,
	}}

	performer.Perform()
}

func listTasks(tasks *[]task.Task) {
	fmt.Printf("Name\t\tDate\tDescription\t\t\tStatus\n")
	for _, t := range *tasks {
		t.ConsoleWrite()
	}
}

func addTask(tasks *[]task.Task) {
	*tasks = append(*tasks, task.ConsoleRead())
}

func selectTask(tasks *[]task.Task) {
	fmt.Printf("id\tName\t\tDate\tDescription\t\t\tStatus\n")
	for i, t := range *tasks {
		fmt.Printf("%d\t", i+1)
		t.ConsoleWrite()
	}

	var id int
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		panic(err)
	}

	t := findTask(*tasks, id-1)

	options := []string{
		"1. Edit task",
		"2. Quit",
	}

	actions := []func(*task.Task){
		editTask,
	}

	var performer userChoice.Performer = &userChoice.ConsolePerformer{
		Options: options,
		Activator: &activator.SingleTaskActivator{
			Task: t, Actions: actions,
		},
	}

	performer.Perform()
}

func findTask(tasks []task.Task, id int) *task.Task {
	if id >= len(tasks) || id < 0 {
		panic("Task not found")
	}

	return &tasks[id]
}

func editTask(t *task.Task) {
	options := []string{
		"1. Edit task name",
		"2. Edit task description",
		"3. Edit task date",
		"4. Quit",
	}

	actions := []func(*task.Task){
		editTaskName,
		editTaskDescription,
		editTaskDate,
	}

	var performer userChoice.Performer = &userChoice.ConsolePerformer{
		Options: options,
		Activator: &activator.SingleTaskActivator{
			Task: t, Actions: actions,
		},
	}

	performer.Perform()
}

func editTaskName(t *task.Task) {
	fmt.Println("Please enter a new name")
	var name string
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		panic(err)
	}
	t.Name = name
}

func editTaskDescription(t *task.Task) {
	fmt.Println("Please enter a new description")
	var description string
	_, err := fmt.Scanf("%s", &description)
	if err != nil {
		panic(err)
	}
	t.Description = description
}

func editTaskDate(t *task.Task) {
	fmt.Println("Please enter a new date in format YYYY-MM-DD")
	var date task.Date
	_, err := fmt.Scanf("%d-%d-%d", &date.Year, &date.Month, &date.Day)
	if err != nil {
		panic(err)
	}
	t.Date = date
}
