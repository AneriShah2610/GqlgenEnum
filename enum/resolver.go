package enum

import (
	"context"
	"fmt"
)

type Resolver struct {
	todos []ToDo
}

func (r *Resolver) MyMutation() MyMutationResolver {
	return &myMutationResolver{r}
}
func (r *Resolver) MyQuery() MyQueryResolver {
	return &myQueryResolver{r}
}

type myMutationResolver struct{ *Resolver }

func (r *myMutationResolver) CreateTodo(ctx context.Context, todo TodoInput, state State) (ToDo, error) {
	newID := fmt.Sprintf("Todo:%d", len(r.todos)+1)

	newToDo := &ToDo{
		ID:    newID,
		Text:  todo.Text,
		State: state,
	}
	r.todos = append(r.todos, *newToDo)
	return *newToDo, nil
}

type myQueryResolver struct{ *Resolver }

func (r *myQueryResolver) Todos(ctx context.Context) ([]ToDo, error) {
	return r.todos, nil
}
func (r *myQueryResolver) Todo(ctx context.Context, id string) (*ToDo, error) {
	panic("not implemented")
}
