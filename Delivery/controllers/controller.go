package controllers

import (
	"net/http"

	"task-manager-go/Domain"
	"task-manager-go/Usecases"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskUsecase Usecases.TaskUsecase
}

func NewTaskController(taskUsecase Usecases.TaskUsecase) *TaskController {
	return &TaskController{taskUsecase: taskUsecase}
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.taskUsecase.GetAllTasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error fetching tasks"})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.taskUsecase.GetTaskById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var newTask Domain.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	err := tc.taskUsecase.CreateTask(newTask)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "task created successfully!", "task": newTask})
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.taskUsecase.DeleteTask(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted successfully!", "task": task})
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask Domain.Task
	if err := c.BindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	task, err := tc.taskUsecase.UpdateTask(id, updatedTask)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "task updated successfully!", "task": task})
}

type UserController struct {
	userUsecase Usecases.UserUsecase
}

func NewUserController(userUsecase Usecases.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}

func (uc *UserController) Login(c *gin.Context) {
	var reqUser Domain.User
	if err := c.BindJSON(&reqUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	token, err := uc.userUsecase.Login(reqUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Login Successful",
		"access_token": token,
	})
}

func (uc *UserController) Register(c *gin.Context) {
	var newUser Domain.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	user, err := uc.userUsecase.Register(newUser)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Register Successful",
		"username": user.Username,
	})
}

func (uc *UserController) UpdateRole(c *gin.Context) {
	username := c.Param("username")
	err := uc.userUsecase.UpdateRole(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update Successful"})
}
