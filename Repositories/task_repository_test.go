package Repositories_test

import (
	"context"
	"log"
	"os"
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

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGODB_URI not set in environment")
	}
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
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
