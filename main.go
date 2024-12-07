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
	{
		ID:        "1",
		Item:      "Clean Room",
		Completed: false,
	},
	{
		ID:        "2",
		Item:      "Read book",
		Completed: false,
	},
	{
		ID:        "3",
		Item:      "Code",
		Completed: false,
	},
}

func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context ){
	var newTodo todo

	if err := context.BindJSON(&newTodo); err!= nil{
		return
	}
	
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(id string) (*todo, error) {
	for i, t:= range todos{
		if t.ID == id{
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getTodo(context *gin.Context){
	id:=context.Param("id")
	todo, err:= getTodoById(id)
	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message":"Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}



func main(){
	//create a default gin router
	rotuer := gin.Default()
	rotuer.GET("/todos", getTodos)
	rotuer.GET("/todos/:id", getTodo)
	rotuer.POST("/todos", addTodo)
	rotuer.Run("localhost:9090")
}

