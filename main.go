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
		"4. Remove all discarded tasks",
		"5. Quit",
	}

	actions := []func(*[]task.Task){
		listTasks,
		addTask,
		selectTask,
		removeAllDiscardedTasks,
	}

	var performer userChoice.Performer = &userChoice.ConsolePerformer{
		Options: options,
		Activator: &activator.MultiTaskActivator{
			Tasks:   &tasks,
			Actions: actions,
		}}

	performer.Perform()
}

func listTasks(tasks *[]task.Task) {
	if len(*tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Printf("Name\t\tDate\t\t\tDescription\t\t\tStatus\n")
	for _, t := range *tasks {
		t.ConsoleWrite()
	}
}

func addTask(tasks *[]task.Task) {
	*tasks = append(*tasks, task.ConsoleRead())
	fmt.Println("The task has been added")
}

func selectTask(tasks *[]task.Task) {
	if len(*tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

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
		"1. Edit",
		"2. Mark as done",
		"3. Discard",
		"4. Restore (after being discarded)",
		"5. Quit",
	}

	actions := []func(*task.Task){
		editTask,
		markTaskAsDone,
		discardTask,
		restoreTask,
	}

	var performer userChoice.Performer = &userChoice.ConsolePerformer{
		Options: options,
		Activator: &activator.SingleTaskActivator{
			Task: t, Actions: actions,
		},
	}

	performer.Perform()
}

func removeAllDiscardedTasks(tasks *[]task.Task) {

	if len(*tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Println("This action will delete all discarded tasks. There is no going back. Are you sure? y/n")

	var choice = ""

	for choice != "y" && choice != "n" {
		_, err := fmt.Scanf("%s", &choice)
		if err != nil {
			panic(err)
		}
	}

	if choice == "n" {
		fmt.Println("Cancelled")
		return
	}

	var newTasks []task.Task

	for _, t := range *tasks {
		if t.Type != task.DISCARDED {
			newTasks = append(newTasks, t)
		}
	}

	*tasks = newTasks

	fmt.Println("All discarded tasks were removed")
}

func markTaskAsDone(t *task.Task) {
	if t.Type == task.DONE {
		fmt.Println("Task is already done")
		return
	}

	if t.Type == task.DISCARDED {
		fmt.Println("The task is discarded, so it cannot be marked as done. Restore the task first")
		return
	}

	t.Type = task.DONE
	fmt.Println("Task is marked as done")
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
	fmt.Println("Name of the task was successfully changed")
}

func editTaskDescription(t *task.Task) {
	fmt.Println("Please enter a new description")
	var description string
	_, err := fmt.Scanf("%s", &description)
	if err != nil {
		panic(err)
	}
	t.Description = description
	fmt.Println("Description of the task was successfully changed")
}

func editTaskDate(t *task.Task) {
	fmt.Println("Please enter a new date in format YYYY-MM-DD")
	var date task.Date
	_, err := fmt.Scanf("%d-%d-%d", &date.Year, &date.Month, &date.Day)
	if err != nil {
		panic(err)
	}
	t.Date = date
	fmt.Println("Date of the task was successfully changed")
}

func discardTask(t *task.Task) {
	t.Type = task.DISCARDED
	fmt.Println("The task was successfully discarded")
}

func restoreTask(t *task.Task) {
	if t.Type != task.DISCARDED {
		fmt.Println("The task is present")
		return
	}

	t.Type = task.PLANNED
	fmt.Println("The task was successfully restored")
}
