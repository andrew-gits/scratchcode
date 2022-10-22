package mocks

import (
	"github.com/andrew-gits/scratchcode/backend/graph/model"
	"github.com/stretchr/testify/mock"
)

type MockedTodoService struct {
	mock.Mock
}

func (s *MockedTodoService) Todos() []*model.Todo {
	args := s.Called()
	args.String()
	todo := model.Todo{ID: "123", Text: "hello world", Done: true, User: &model.User{ID: "1234", Name: "John"}}
	todos := []*model.Todo{&todo}
	return todos //.(*model.Todo)
}
