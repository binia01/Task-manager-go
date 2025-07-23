package Repositories

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"task-manager-go/Domain"
)

type TaskRepository interface {
	GetAllTasks() ([]Domain.Task, error)
	GetTaskById(id string) (*Domain.Task, error)
	CreateTask(task Domain.Task) error
	UpdateTask(id string, task Domain.Task) (*Domain.Task, error)
	DeleteTask(id string) (*Domain.Task, error)
}

type taskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository() TaskRepository {
	client := connectDb()
	collection := client.Database("Task-Database").Collection("Tasks")
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

func connectDb() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://<username>:<password>@cluster0.tj8um.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
	return client
}
