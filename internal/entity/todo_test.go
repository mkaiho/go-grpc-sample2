package entity

import (
	"reflect"
	"testing"
)

func TestTodos_Completed(t *testing.T) {
	testTodos := Todos{
		{
			id:           "todo_id_001",
			name:         "todo_name_001",
			hasCompleted: true,
		},
		{
			id:           "todo_id_002",
			name:         "todo_name_002",
			hasCompleted: false,
		},
		{
			id:           "todo_id_003",
			name:         "todo_name_003",
			hasCompleted: true,
		},
		{
			id:           "todo_id_004",
			name:         "todo_name_004",
			hasCompleted: false,
		},
	}
	tests := []struct {
		name string
		tr   Todos
		want Todos
	}{
		{
			name: "return only tasks that have completed",
			tr:   testTodos,
			want: Todos{
				testTodos[0],
				testTodos[2],
			},
		},
		{
			name: "return nil when todos is nil",
			tr:   nil,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Completed(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Todos.Completed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodos_NotCompleted(t *testing.T) {
	testTodos := Todos{
		{
			id:           "todo_id_001",
			name:         "todo_name_001",
			hasCompleted: true,
		},
		{
			id:           "todo_id_002",
			name:         "todo_name_002",
			hasCompleted: false,
		},
		{
			id:           "todo_id_003",
			name:         "todo_name_003",
			hasCompleted: true,
		},
		{
			id:           "todo_id_004",
			name:         "todo_name_004",
			hasCompleted: false,
		},
	}
	tests := []struct {
		name string
		tr   Todos
		want Todos
	}{
		{
			name: "return only tasks that have not completed",
			tr:   testTodos,
			want: Todos{
				testTodos[1],
				testTodos[3],
			},
		},
		{
			name: "return nil when todos is nil",
			tr:   nil,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.NotCompleted(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Todos.NotCompleted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodos_Append(t *testing.T) {
	testTodos := Todos{
		{
			id:           "todo_id_001",
			name:         "todo_name_001",
			hasCompleted: true,
		},
		{
			id:           "todo_id_002",
			name:         "todo_name_002",
			hasCompleted: false,
		},
		{
			id:           "todo_id_003",
			name:         "todo_name_003",
			hasCompleted: true,
		},
		{
			id:           "todo_id_004",
			name:         "todo_name_004",
			hasCompleted: false,
		},
	}
	type args struct {
		todos Todos
	}
	tests := []struct {
		name string
		tr   Todos
		args args
		want Todos
	}{
		{
			name: "return todos added new todo",
			tr:   testTodos[0:2],
			args: args{
				todos: testTodos[2:4],
			},
			want: testTodos[0:4],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Append(tt.args.todos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Todos.Append() = %v, want %v", got, tt.want)
			}
		})
	}
}
