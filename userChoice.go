package main

import (
	"fmt"
	"strconv"
)

func performUserChoiceWithListOfTasks(tasks *[]task, options []string, actions []func(*[]task)) {
	if len(options) != len(actions)+1 {
		panic("Incorrect number of options")
	}

	var exitChoiceStr = options[len(options)-1][0:1]

	exitChoice, err := strconv.Atoi(exitChoiceStr)

	if err != nil {
		panic(err)
	}

	choice := -1

	for choice != exitChoice {
		for _, option := range options {
			fmt.Println(option)
		}

		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			panic(err)
		}

		if choice < 1 || choice > exitChoice {
			fmt.Println("Invalid choice")
			break
		}

		if choice != exitChoice {
			actions[choice-1](tasks)
		}
	}
}

func performUserChoiceWithSingleTask(task *task, options []string, actions []func(*task)) {
	if len(options) != len(actions)+1 {
		panic("Incorrect number of options")
	}

	var exitChoiceStr = options[len(options)-1][0:1]

	exitChoice, err := strconv.Atoi(exitChoiceStr)

	if err != nil {
		panic(err)
	}

	choice := -1

	for choice != exitChoice {
		for _, option := range options {
			fmt.Println(option)
		}

		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			panic(err)
		}

		if choice < 1 || choice > exitChoice {
			fmt.Println("Invalid choice")
			break
		}

		if choice != exitChoice {
			actions[choice-1](task)
		}
	}
}
