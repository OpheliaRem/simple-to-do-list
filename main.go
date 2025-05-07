package main

import "fmt"
import "os"

func main() {
	var tasks []task

	options := []string{
		"1. Create a new task",
		"2. List all tasks",
		"3. Edit task",
		"4. Quit",
	}

	actions := []func(*[]task){
		addTask,
		listTasks,
		editTask,
	}

	performUserChoiceWithListOfTasks(&tasks, options, actions)
}

func listTasks(tasks *[]task) {
	fmt.Printf("Name\t\tDate\tDescription\t\t\tStatus\n")
	for _, task := range *tasks {
		fmt.Println(task.name, task.month.String(), task.day, task.year, task.description, task.taskType.String())
	}
}

func addTask(tasks *[]task) {
	fmt.Println("Please enter a task name")
	var name string
	_, err := fmt.Fscan(os.Stdin, &name)
	if err != nil {
		return
	}

	fmt.Println("Please enter the deadline date of your task in format dd.mm.yyyy")
	var date date
	_, err = fmt.Scanf("%d.%d.%d", &date.day, &date.month, &date.year)
	if err != nil {
		return
	}

	fmt.Println("Please enter a task description")
	var description string
	_, err = fmt.Fscan(os.Stdin, &description)
	if err != nil {
		return
	}

	*tasks = append(*tasks, task{name, date, description, PLANNED})
}

func editTask(tasks *[]task) {
	correctChoice := false
	var choice int

	for !correctChoice {
		for i, task := range *tasks {
			fmt.Println(i+1, task.name, task.month.String(), task.day, task.description, task.taskType.String())
		}

		fmt.Println("Please choose a task to edit")

		_, err := fmt.Fscan(os.Stdin, &choice)
		if err != nil {
			return
		}

		if choice-1 >= len(*tasks) || choice-1 < 0 {
			fmt.Println("Invalid choice: No task with this number exists. Try again.")
		} else {
			correctChoice = true
		}
	}

	options := []string{
		"1. Edit name",
		"2. Edit date",
		"3. Edit description",
		"4. Discard task",
		"5. Exit",
	}

	actions := []func(*task){
		editTaskName,
		editTaskDate,
		editTaskDescription,
		discardTask,
	}

	performUserChoiceWithSingleTask(&(*tasks)[choice-1], options, actions)
}

func editTaskName(task *task) {
	fmt.Println("Please enter a task name")
	var name string
	_, err := fmt.Fscan(os.Stdin, &name)
	if err != nil {
		panic("Error reading task name")
	}

	task.name = name
}

func editTaskDescription(task *task) {
	fmt.Println("Please enter a task description")
	var description string
	_, err := fmt.Fscan(os.Stdin, &description)
	if err != nil {
		panic("Error reading task description")
	}

	task.description = description
}

func editTaskDate(task *task) {
	fmt.Println("Please enter a task date")
	_, err := fmt.Scanf("%d.%d.%d", &task.day, &task.month, &task.year)
	if err != nil {
		panic("Error reading task date")
	}
}

func discardTask(task *task) {
	task.taskType = DISCARDED
}
