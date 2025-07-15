package data

import (
	"time"

	"task-manager-go/models"
)

var tasks = []models.Task{
    {ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
    {ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
    {ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}


func GetAllTasks() []models.Task {
	return tasks
}

func GetTaskById(id string) (*models.Task, bool) {
	for _, t := range tasks {
		if t.ID == id {
			return &t, true
		}
	}
	return nil, false
}

func UpdateTask(id string, updatedTask models.Task) (*models.Task, bool) {
	for i, t := range tasks {
		if t.ID == id {
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			if !updatedTask.DueDate.IsZero() {
				tasks[i].DueDate = updatedTask.DueDate
			}
			if updatedTask.Status != "" {
				tasks[i].Status = updatedTask.Status
			}
			return &tasks[i], true
		}
	}
	return nil, false
}

func CreateTask(task models.Task) bool {
	for _, t := range tasks {
		if t.ID == task.ID {
			return false
		}
	}
	tasks = append(tasks, task)
	return true
}

func DeleteTask(id string) (*models.Task, bool) {
	for i, t := range tasks {
		if t.ID == id {
			deleted := tasks[i]
			tasks = append(tasks[:i], tasks[i+1:]...)
			return &deleted, true
		}
	}
	return nil, false
}