package internal

import (
	"context"
	"fmt"
	"my-app/internal/repo"
)

type IService interface {
	AddTodo(ctx context.Context, name string) (data repo.Todo, err error)
	GetTodos(ctx context.Context, search string) (data []repo.Todo, err error)
	UpdateTodo(ctx context.Context, arg repo.UpdateTodoParams) (data repo.Todo, err error)
}

type Service struct {
	querier repo.Querier
}

func NewService(querier repo.Querier) *Service {
	return &Service{querier: querier}
}

func (s Service) AddTodo(ctx context.Context, name string) (data repo.Todo, err error) {
	data, err = s.querier.AddTodo(ctx, name)
	return
}

func (s Service) GetTodos(ctx context.Context, search string) (data []repo.Todo, err error) {
	data, err = s.querier.GetTodos(ctx, fmt.Sprintf("%%%s%%", search))
	return
}

func (s Service) UpdateTodo(ctx context.Context, arg repo.UpdateTodoParams) (data repo.Todo, err error) {
	data, err = s.querier.UpdateTodo(ctx, arg)
	return
}
