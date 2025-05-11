package userChoice

import (
	"fmt"
	"strconv"
	"toDoList/activator"
)

type Performer interface {
	Perform()
}

type ConsolePerformer struct {
	Options   []string
	Activator activator.Activator
}

func (p *ConsolePerformer) Perform() {
	var exitChoiceStr = p.Options[len(p.Options)-1][0:1]

	exitChoice, err := strconv.Atoi(exitChoiceStr)
	if err != nil {
		panic(err)
	}

	choice := -1
	for choice != exitChoice {
		for _, option := range p.Options {
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
			p.Activator.Act(choice - 1)
		}
	}
}
