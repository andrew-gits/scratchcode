package service

import "github.com/andrew-gits/scratchcode/backend/graph/model"

type TodoService interface {
	Todos() []*model.Todo
}
