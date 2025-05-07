package main

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
