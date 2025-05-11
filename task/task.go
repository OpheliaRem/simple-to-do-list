package task

import "fmt"

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

type Date struct {
	Year  int
	Month month
	Day   int
}

type taskType uint

const (
	PLANNED taskType = iota
	DONE
	DISCARDED
	MISSED
)

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

type Task struct {
	Name        string
	Date        Date
	Description string
	Type        taskType
}

func (t *Task) ConsoleWrite() {
	fmt.Printf(
		"%s\t%d\t%s\t%d\t%s\t%s\n",
		t.Name,
		t.Date.Year,
		t.Date.Month,
		t.Date.Day,
		t.Description,
		t.Type,
	)
}

func ConsoleRead() Task {
	var name string
	var date Date
	var description string

	fmt.Println("Please enter name of task:")
	_, err := fmt.Scan(&name)
	if err != nil {
		panic(err)
	}

	fmt.Println("Please enter date of task in format YYYY-MM-DD:")
	_, err = fmt.Scanf("%d-%d-%d", &date.Year, &date.Month, &date.Day)
	if err != nil {
		panic(err)
	}

	fmt.Println("Please enter description of task:")
	_, err = fmt.Scan(&description)
	if err != nil {
		panic(err)
	}

	return Task{name, date, description, PLANNED}
}
