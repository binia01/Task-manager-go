package Repositories

import (
	"context"
	"fmt"
	"log"

	"task-manager-go/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type taskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(collection *mongo.Collection) Domain.ITaskRepository {
	return &taskRepository{collection: collection}
}

func (tr *taskRepository) GetAllTasks() ([]Domain.Task, error) {
	findOptions := options.Find()
	var tasks []Domain.Task

	cur, err := tr.collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var elem Domain.Task
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err)
			continue
		}
		tasks = append(tasks, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (tr *taskRepository) GetTaskById(id string) (*Domain.Task, error) {
	var task Domain.Task
	err := tr.collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&task)
	if err != nil {
		return nil, fmt.Errorf("task not found")
	}
	return &task, nil
}

func (tr *taskRepository) CreateTask(task Domain.Task) error {
	// Check if task exists
	var existingTask Domain.Task
	err := tr.collection.FindOne(context.TODO(), bson.M{"id": task.ID}).Decode(&existingTask)
	if err == nil {
		return fmt.Errorf("task with this ID already exists")
	}

	_, err = tr.collection.InsertOne(context.TODO(), task)
	return err
}

func (tr *taskRepository) UpdateTask(id string, updatedTask Domain.Task) (*Domain.Task, error) {
	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"due_date":    updatedTask.DueDate,
			"status":      updatedTask.Status,
		},
	}

	_, err := tr.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, fmt.Errorf("task not found")
	}

	var task Domain.Task
	err = tr.collection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return nil, fmt.Errorf("task not found")
	}
	return &task, nil
}

func (tr *taskRepository) DeleteTask(id string) (*Domain.Task, error) {
	var task Domain.Task
	err := tr.collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&task)
	if err != nil {
		return nil, fmt.Errorf("task not found")
	}

	_, err = tr.collection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		return nil, fmt.Errorf("error deleting task")
	}
	return &task, nil
}
