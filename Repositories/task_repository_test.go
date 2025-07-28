package Repositories_test

import (
	"context"
	"testing"
	"time"

	"task-manager-go/Domain"
	"task-manager-go/Repositories"

	// "task-manager-go/Repositories"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupTaskTestDB() *mongo.Collection {
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://username:password@cluster0.tj8um.mongodb.net/?retryWrites=true&w=majority"))
	col := client.Database("task_test_db").Collection("tasks")
	col.DeleteMany(context.TODO(), map[string]interface{}{})
	return col
}

func TestTaskRepository_CreateAndGet(t *testing.T) {
	col := setupTaskTestDB()
	repo := Repositories.NewTaskRepository(col)

	task := Domain.Task{
		ID:          "task1",
		Title:       "Task Title",
		Description: "Task Description",
		DueDate:     time.Now(),
		Status:      "open",
	}

	err := repo.CreateTask(task)
	assert.NoError(t, err)

	result, err := repo.GetTaskById("task1")
	assert.NoError(t, err)
	assert.Equal(t, "Task Title", result.Title)
}

func TestTaskRepository_UpdateAndDelete(t *testing.T) {
	col := setupTaskTestDB()
	repo := Repositories.NewTaskRepository(col)

	task := Domain.Task{ID: "task2", Title: "Old", Description: "Old Desc", DueDate: time.Now(), Status: "open"}
	_ = repo.CreateTask(task)

	updated := Domain.Task{Title: "New", Description: "New Desc", DueDate: task.DueDate, Status: "done"}
	updatedTask, err := repo.UpdateTask("task2", updated)
	assert.NoError(t, err)
	assert.Equal(t, "New", updatedTask.Title)

	deletedTask, err := repo.DeleteTask("task2")
	assert.NoError(t, err)
	assert.Equal(t, "New", deletedTask.Title)
}
