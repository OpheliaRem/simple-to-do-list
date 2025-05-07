package main

import "fmt"
import "os"

type month uint

func (m month) String() string {
	switch m {
	case 1:
		return "January"
	case 2:
		return "February"
	case 3:
		return "March"
	case 4:
		return "April"
	case 5:
		return "May"
	case 6:
		return "June"
	case 7:
		return "July"
	case 8:
		return "August"
	case 9:
		return "September"
	case 10:
		return "October"
	case 11:
		return "November"
	case 12:
		return "December"
	default:
		return "Unknown"
	}
}

type date struct {
	year int
	month
	day int
}

type taskType uint

func (t taskType) String() string {
	switch t {
	case PLANNED:
		return "PLANNED"
	case DONE:
		return "DONE"
	case DISCARDED:
		return "DISCARDED"
	case MISSED:
		return "MISSED"
	default:
		return "UNKNOWN"
	}
}

const (
	PLANNED taskType = iota
	DONE
	DISCARDED
	MISSED
)

type task struct {
	name string
	date
	description string
	taskType
}

func main() {
	var tasks []task

	const quitNumber int = 3
	userChoice := 0

	for userChoice != quitNumber {
		fmt.Println("Please select an option below:")
		fmt.Println("1. Create a new task")
		fmt.Println("2. List all tasks")
		fmt.Printf("%d. Quit\n", quitNumber)

		_, err := fmt.Fscan(os.Stdin, &userChoice)
		if err != nil {
			return
		}

		switch userChoice {
		case 1:
			addTask(&tasks)
		case 2:
			listTasks(&tasks)
		case quitNumber:
			fmt.Println("Bye")
		default:
			fmt.Println("Invalid option")
			break
		}
	}
}

func listTasks(tasks *[]task) {
	fmt.Printf("Name\tDate\tDescription\tStatus\n")
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
