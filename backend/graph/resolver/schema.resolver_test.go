package resolver_test

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/andrew-gits/scratchcode/backend/graph/generated"
	"github.com/andrew-gits/scratchcode/backend/graph/mocks"
	"github.com/andrew-gits/scratchcode/backend/graph/model"

	"github.com/andrew-gits/scratchcode/backend/graph/resolver"
	"github.com/stretchr/testify/require"
)

func TestQueryResolver_GetTodo(t *testing.T) {

	t.Run("should receive a Todo", func(t *testing.T) {
		testTodoService := new(mocks.MockedTodoService)
		resolvers := resolver.Resolver{TodoService: testTodoService}
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers})))
		todo := model.Todo{ID: "123", Text: "hello world", Done: true, User: &model.User{ID: "1234", Name: "John"}}
		ret := []*model.Todo{&todo}
		testTodoService.On("Todos").Return(ret)

		var resp struct {
			Todos []struct {
				ID   string
				Text string
				Done bool
			}
		}
		q := `
query test {
    todos {
        id,
        text,
        done
    }
}
	`
		c.MustPost(q, &resp)
		testTodoService.AssertNumberOfCalls(t, "Todos", 1)
		testTodoService.AssertExpectations(t)
		require.Equal(t, "123", resp.Todos[0].ID)
		require.Equal(t, "hello world", resp.Todos[0].Text)
		require.Equal(t, true, resp.Todos[0].Done)
	})
}
