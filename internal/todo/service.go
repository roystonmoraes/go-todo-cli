package todo

import (
	"errors"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) List() ([]Todo, error) {
	return s.repo.Load()
}

func (s *Service) Add(task string) error {
	todos, err := s.repo.Load()
	if err != nil {
		return err
	}

	nextID := 1

	if len(todos) > 0 {
		nextID = todos[len(todos)-1].ID + 1
	}

	todos = append(todos, Todo{
		ID:        nextID,
		Task:      task,
		Completed: false,
	})

	return s.repo.Save(todos)
}

func (s *Service) Complete(id int) error {
	todos, err := s.repo.Load()
	if err != nil {
		return err
	}

	found := false

	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = true
			found = true
			break
		}
	}

	if !found {
		return errors.New("todo not found")
	}

	return s.repo.Save(todos)
}

func (s *Service) Delete(id int) error {
	todos, err := s.repo.Load()
	if err != nil {
		return err
	}

	index := -1

	for i := range todos {
		if todos[i].ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("todo not found")
	}

	todos = append(
		todos[:index],
		todos[index+1:]...,
	)

	return s.repo.Save(todos)
}
