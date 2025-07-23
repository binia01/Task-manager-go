package Usecases

import (
	"task-manager-go/Domain"
	"task-manager-go/Repositories"
)

type TaskUsecase interface {
	GetAllTasks() ([]Domain.Task, error)
	GetTaskById(id string) (*Domain.Task, error)
	CreateTask(task Domain.Task) error
	UpdateTask(id string, task Domain.Task) (*Domain.Task, error)
	DeleteTask(id string) (*Domain.Task, error)
}

type taskUsecase struct {
	taskRepo Repositories.TaskRepository
}

func NewTaskUsecase(taskRepo Repositories.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo: taskRepo}
}

func (tu *taskUsecase) GetAllTasks() ([]Domain.Task, error) {
	return tu.taskRepo.GetAllTasks()
}

func (tu *taskUsecase) GetTaskById(id string) (*Domain.Task, error) {
	return tu.taskRepo.GetTaskById(id)
}

func (tu *taskUsecase) CreateTask(task Domain.Task) error {
	return tu.taskRepo.CreateTask(task)
}

func (tu *taskUsecase) UpdateTask(id string, task Domain.Task) (*Domain.Task, error) {
	return tu.taskRepo.UpdateTask(id, task)
}

func (tu *taskUsecase) DeleteTask(id string) (*Domain.Task, error) {
	return tu.taskRepo.DeleteTask(id)
}
