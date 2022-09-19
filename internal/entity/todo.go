package entity

type Todo struct {
	id           string
	name         string
	hasCompleted bool
}

func NewTodo(id string, name string) *Todo {
	return &Todo{
		id:           id,
		name:         name,
		hasCompleted: false,
	}
}

func (t *Todo) ID() string {
	if t == nil {
		return ""
	}
	return t.id
}

func (t *Todo) Name() string {
	if t == nil {
		return ""
	}
	return t.name
}

func (t *Todo) HasCompleted() bool {
	if t == nil {
		return false
	}
	return t.hasCompleted
}

type Todos []*Todo

func NewTodos(items []*Todo) Todos {
	todos := append(Todos{}, items...)
	return todos
}

func (t Todos) Completed() Todos {
	return t.filterByHasCompleted(true)
}

func (t Todos) NotCompleted() Todos {
	return t.filterByHasCompleted(false)
}

func (t Todos) Append(todos Todos) Todos {
	return append(t, todos...)
}

func (t Todos) filterByHasCompleted(hasCompleted bool) Todos {
	var todos Todos
	for _, todo := range t {
		if todo.hasCompleted == hasCompleted {
			todos = append(todos, todo)
		}
	}
	return todos
}
