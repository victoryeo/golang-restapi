package main

import (
  "net/http"
  "fmt"
	"errors"
  "github.com/gin-gonic/gin"
)

type todo struct {
	ID string `json:"id"`
	Item string `json:"title"`
	Completed bool `json:"completed"`
}

var todos = []todo{
  {ID:"1", Item:"Clean room", Completed: false},
  {ID:"2", Item:"Record room", Completed: false},
  {ID:"3", Item:"Read room", Completed: false},
}

func getHello(context *gin.Context) {
  fmt.Print("getHello")
  context.IndentedJSON(http.StatusOK, "hello")
}

func getTodos(context *gin.Context) {
  context.IndentedJSON(http.StatusOK, todos)
}

func getTodoByID(id string) (*todo, error) {

	for i, t := range todos {
			if t.ID == id {
					return &todos[i], nil
			}
	}
  return nil, errors.New("todo not found")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message":"not found"})
	  return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func addTodo(context *gin.Context) {
  var newTodo todo
  if err := context.BindJSON(&newTodo); err!= nil {
    return
  }

  todos = append(todos, newTodo)
  context.IndentedJSON(http.StatusCreated, newTodo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message":"not found"})
	  return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}
 
func main() {
  fmt.Print("Code is a ", " portal.\n")
  router := gin.Default()

  router.GET("/", getHello)
  router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
  router.POST("/todos", addTodo)
  router.NoRoute(func(c *gin.Context) {
    c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
  })
  router.Run("localhost:9090")
}
