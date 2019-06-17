package main 

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"fmt"
)

type Todo struct {
	ID 	string `json:"id"`
	Title	string `json:"title"`
	Status	string `json:"status"`  
}

var store []Todo

func main() {

	routes := gin.Default()
	api := routes.Group("/api")
	api.Handle(http.MethodPost, "/todos", postToDo)
	api.Handle(http.MethodGet, "todos/:id", getByID)
	api.Handle(http.MethodGet, "todos", todos)
	api.Handle(http.MethodPut, "todos/:id", putToDo)
	api.Handle(http.MethodDelete, "todos/:id", deleteToDo)
       
	routes.Run(":1234")
}

func deleteToDo(c *gin.Context) {
	param := c.Param("id")
	for _, todo := range store {
		if todo.ID == param {
			index, err := strconv.Atoi(todo.ID)

			if err != nil {
				c.JSON(http.StatusBadRequest, "cannot delete")
				return
			}

			if index == len(store) {
				store = store[:index - 1]
			} else {
				store = append(store[:index], store[index+1:]...)
			}
		
			c.JSON(200, gin.H {
				"status": "success",
			})
			return
		}
	}
		c.JSON(http.StatusBadRequest, "cannot delete")
}

func putToDo(c *gin.Context) {
	param := c.Param("id")

	var req Todo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	for _, todo := range store {
		if todo.ID == param {
			todo.Title = req.Title
			todo.Status = req.Status
			c.JSON(200, todo)
			return
		}
	}
	fmt.Println(store)
	c.JSON(http.StatusBadRequest, "")
}

func todos(c *gin.Context) {
	c.JSON(http.StatusOK, store)
}

func getByID(c *gin.Context) {
	param := c.Param("id")
	for _, todo := range store {
		if todo.ID == param {
			c.JSON(http.StatusOK, todo)
			return
		}
	}
	c.JSON(http.StatusBadRequest, "")
}

func postToDo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	todo.ID = strconv.Itoa(len(store) + 1)
	store = append(store, todo)
	c.JSON(201, todo)

}