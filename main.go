package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victoryeo/golang-restapi/controllers"
	"github.com/victoryeo/golang-restapi/database"
	"github.com/victoryeo/golang-restapi/middleware"
	"github.com/victoryeo/golang-restapi/models"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean room", Completed: false},
	{ID: "2", Item: "Record room", Completed: false},
	{ID: "3", Item: "Read room", Completed: false},
}

func getRoot(context *gin.Context) {
	fmt.Print("getRoot")
	// Call the HTML method of the Context to render a template
	context.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "Home Page",
		},
	)
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
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func addTodo(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func testPrivate(c *gin.Context) {
	uidStr := c.Param("uid")

	if uidStr != "" {
		c.JSON(200, gin.H{"uid": uidStr})
		return
	}

	c.JSON(200, gin.H{"error": "unknown uid"})
}

func main() {
	fmt.Print("Code is a ", " portal.\n")

	// Initialize Database
	database.Connect("host=localhost user=postgres password=password dbname=user_login port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	database.Migrate()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.POST("/login", middleware.Login)

	models.ConnectDatabase()
	router.GET("/books", controllers.FindBooks)
	router.GET("/books/:id", controllers.FindBook)
	router.POST("/books", controllers.CreateBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.GET("/", getRoot)
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	privateRouter := router.Group("/private")
	privateRouter.Use(middleware.JWTTokenCheck)
	privateRouter.GET("/test/:uid", testPrivate)

	router.Run("localhost:9090")
}
