package Usecases

import (
	"task-manager-go/Domain"
)

type TaskUsecase struct {
	taskRepo Domain.ITaskRepository
}

func NewTaskUsecase(taskRepo Domain.ITaskRepository) *TaskUsecase {
	return &TaskUsecase{taskRepo: taskRepo}
}

func (tu *TaskUsecase) GetAllTasks() ([]Domain.Task, error) {
	return tu.taskRepo.GetAllTasks()
}

func (tu *TaskUsecase) GetTaskById(id string) (*Domain.Task, error) {
	return tu.taskRepo.GetTaskById(id)
}

func (tu *TaskUsecase) CreateTask(task Domain.Task) error {
	return tu.taskRepo.CreateTask(task)
}

func (tu *TaskUsecase) UpdateTask(id string, task Domain.Task) (*Domain.Task, error) {
	return tu.taskRepo.UpdateTask(id, task)
}

func (tu *TaskUsecase) DeleteTask(id string) (*Domain.Task, error) {
	return tu.taskRepo.DeleteTask(id)
}
