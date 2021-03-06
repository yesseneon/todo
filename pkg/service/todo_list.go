package service

import (
	"github.com/yesseneon/todo"
	"github.com/yesseneon/todo/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func newTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userID int, list todo.TodoList) (int, error) {
	return s.repo.Create(userID, list)
}

func (s *TodoListService) GetAll(userID int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userID)
}

func (s *TodoListService) GetByID(userID, listID int) (todo.TodoList, error) {
	return s.repo.GetByID(userID, listID)
}

func (s *TodoListService) Update(userID, listID int, list todo.TodoListUpdate) error {
	if err := list.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userID, listID, list)
}

func (s *TodoListService) Delete(userID, listID int) error {
	return s.repo.Delete(userID, listID)
}
