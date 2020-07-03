package graphql

import (
	"context"
	"fmt"
	"testing"

	"github.com/machinebox/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestDockerServer(t *testing.T) {
	if testing.Short() {
		t.Skip("skip docker mapdb integration test")
	}
	ctx := context.Background()

	// start container
	containerReq := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    ".",
			Dockerfile: "Dockerfile",
		},
		ExposedPorts: []string{"8080/tcp"},
		Env: map[string]string{
			"SCHEMA_FILE":   "example/mapdb/schema/todo.graphql",
			"DATABASE_TYPE": "MapDB",
		},
		Name:       "graphql-crud-server-test",
		WaitingFor: wait.ForLog("start graphql server"),
		AutoRemove: true,
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: containerReq,
		Started:          true,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer container.Terminate(ctx)
	ip, err := container.Host(ctx)
	if err != nil {
		t.Fatal(err)
	}
	port, err := container.MappedPort(ctx, "8080")
	if err != nil {
		t.Fatal(err)
	}

	// start test
	execIntegrationTests(t, ctx, fmt.Sprintf("http://%s:%s/graphql", ip, port.Port()))
}

func execIntegrationTests(t *testing.T, ctx context.Context, address string) {
	client := graphql.NewClient(address)

	// test create
	req := graphql.NewRequest(`
		mutation {
		  createTodo(todo: { text: "Hello World" }) {
			id
			text
		  }
		}
	`)
	var resp map[string]interface{}
	err := client.Run(ctx, req, &resp)
	if err != nil {
		t.Fatal(err)
	}
	createTodo, ok := resp["createTodo"].(map[string]interface{})
	if !ok {
		t.Fatal("response has no createTodo field")
	}
	assert.NotEmpty(t, createTodo["id"], "id should not be empty")
	assert.Equal(t, "Hello World", createTodo["text"], "text should be \"Hello World\"")

	// test update
	req = graphql.NewRequest(`
		mutation ($id: ID!) {
		  updateTodo(id: $id, todo: { text: "Hello World!" }) {
			id
			text
		  }
		}
	`)
	req.Var("id", createTodo["id"].(string))
	err = client.Run(ctx, req, &resp)
	if err != nil {
		t.Fatal(err)
	}
	updateTodo, ok := resp["updateTodo"].(map[string]interface{})
	if !ok {
		t.Fatal("response has no updateTodo field")
	}
	assert.Equal(t, createTodo["id"], updateTodo["id"], "id should be the same as on create")
	assert.Equal(t, "Hello World!", updateTodo["text"], "text should be \"Hello World!\"")

	// test read
	req = graphql.NewRequest(`
		query ($id: ID!) {
		  readTodo(id: $id) {
			id
			text
		  }
		}
	`)
	req.Var("id", updateTodo["id"].(string))
	err = client.Run(ctx, req, &resp)
	if err != nil {
		t.Fatal(err)
	}
	readTodo, ok := resp["readTodo"].(map[string]interface{})
	if !ok {
		t.Fatal("response has no readTodo field")
	}
	assert.Equal(t, updateTodo["id"], readTodo["id"], "id should be the same as on update")
	assert.Equal(t, "Hello World!", readTodo["text"], "text should be \"Hello World!\"")

	// test read all
	req = graphql.NewRequest(`
		{
		  readTodos {
			id
			text
		  }
		}
	`)
	err = client.Run(ctx, req, &resp)
	if err != nil {
		t.Fatal(err)
	}
	readTodos, ok := resp["readTodos"].([]interface{})
	if !ok {
		t.Fatal("response has no readTodos field")
	}
	assert.Len(t, readTodos, 1, "readTodos should have len 1")

	// test delete
	req = graphql.NewRequest(`
		mutation ($id: ID!) {
		  deleteTodo(id: $id) {
			id
			text
		  }
		}
	`)
	req.Var("id", readTodo["id"].(string))
	err = client.Run(ctx, req, &resp)
	if err != nil {
		t.Fatal(err)
	}
	deleteTodo, ok := resp["deleteTodo"].(map[string]interface{})
	if !ok {
		t.Fatal("response has no updateTodo field")
	}
	assert.Equal(t, readTodo["id"], deleteTodo["id"], "id should be the same as on read")
	assert.Equal(t, "Hello World!", deleteTodo["text"], "text should be \"Hello World!\"")

	// test read all
	req = graphql.NewRequest(`
		{
		  readTodos {
			id
			text
		  }
		}
	`)
	err = client.Run(ctx, req, &resp)
	if err != nil {
		t.Fatal(err)
	}
	readTodos, ok = resp["readTodos"].([]interface{})
	if !ok {
		t.Fatal("response has no readTodos field")
	}
	assert.Len(t, readTodos, 0, "readTodos should have len 0")
}
