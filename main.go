package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Learn Go", Completed: false},
	{ID: "2", Item: "Learn GraphQL", Completed: false},
	{ID: "3", Item: "Learn GraphQL with Go", Completed: false},
}

func getTodos(context *gin.Context) {
	//Convert to JSON and return
	context.IndentedJSON(http.StatusOK, todos)

}

func addToDo(context *gin.Context) {
	var newTodo todo
	//Bind JSON to struct
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	//Append to slice
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)

}

func getToDo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoByID(id string) (*todo, error) {
	for index, todo := range todos {
		if todo.ID == id {
			return &todos[index], nil
		}
	}
	return nil, errors.New("todo not found")

}

func toggleToDoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)

}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getToDo)
	router.PATCH("/todos/:id", toggleToDoStatus)
	router.POST("/todos", addToDo)
	router.Run("localhost:9090")
}
