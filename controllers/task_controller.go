package controllers

import (
	"net/http"

	"task-manager-go/data"
	"task-manager-go/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	tasks := data.GetAllTasks()
	c.IndentedJSON(http.StatusOK, tasks)
}
func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task, found := data.GetTaskById(id)
	if !found {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	if ok := data.CreateTask(newTask); !ok {
		c.JSON(http.StatusConflict, gin.H{"message": "Task with this ID already exists"})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "task created successfully!", "task": newTask})
}
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	task, found := data.DeleteTask(id)
	if !found {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted successfully!", "task": task})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask models.Task
	if err := c.BindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	newTask, ok := data.UpdateTask(id, updatedTask)
	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "task updated successfully!", "task": newTask})
}
