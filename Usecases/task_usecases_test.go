package Usecases_test

import (
	"task-manager-go/Domain"
	"task-manager-go/Usecases"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskRepo struct {
	mock.Mock
}

func (m *MockTaskRepo) GetAllTasks() ([]Domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]Domain.Task), args.Error(1)
}
func (m *MockTaskRepo) GetTaskById(id string) (*Domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*Domain.Task), args.Error(1)
}
func (m *MockTaskRepo) CreateTask(task Domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}
func (m *MockTaskRepo) UpdateTask(id string, task Domain.Task) (*Domain.Task, error) {
	args := m.Called(id, task)
	return args.Get(0).(*Domain.Task), args.Error(1)
}
func (m *MockTaskRepo) DeleteTask(id string) (*Domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*Domain.Task), args.Error(1)
}

func TestGetAllTasks(t *testing.T) {
	mockRepo := new(MockTaskRepo)
	usecase := Usecases.NewTaskUsecase(mockRepo)

	mockTasks := []Domain.Task{
		{ID: "1", Title: "Test 1"},
		{ID: "2", Title: "Test 2"},
	}

	mockRepo.On("GetAllTasks").Return(mockTasks, nil)
	tasks, err := usecase.GetAllTasks()

	assert.NoError(t, err)
	assert.Len(t, tasks, 2)
	mockRepo.AssertExpectations(t)
}
