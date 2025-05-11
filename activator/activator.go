package activator

import "toDoList/task"

type Activator interface {
	Act(indexToActivate int)
}

type SingleTaskActivator struct {
	Task    *task.Task
	Actions []func(*task.Task)
}

func (s *SingleTaskActivator) Act(indexToActivate int) {
	s.Actions[indexToActivate](s.Task)
}

type MultiTaskActivator struct {
	Tasks   *[]task.Task
	Actions []func(*[]task.Task)
}

func (m *MultiTaskActivator) Act(indexToActivate int) {
	m.Actions[indexToActivate](m.Tasks)
}
