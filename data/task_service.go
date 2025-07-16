package data

import (
	"context"
	"fmt"
	"log"

	"task-manager-go/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db = connectDb()
var collection = db.Database("Task-Database").Collection("Tasks")

func GetAllTasks() []models.Task {
	findOptions := options.Find()

	var tasks []models.Task

	cur, err := collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem models.Task
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	return tasks
}

func GetTaskById(id string) (*models.Task, bool) {
	var task models.Task

	err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&task)
	if err != nil {
		return nil, false
	}
	return &task, true
}

func UpdateTask(id string, updatedTask models.Task) (*models.Task, bool) {
	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"due_date":    updatedTask.DueDate,
			"status":      updatedTask.Status,
		},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return nil, false
	}

	var task models.Task
	err = collection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return nil, false
	}
	return &task, true
}

func CreateTask(task models.Task) bool {
	// Check if a task with the same ID already exists
	var existingTask models.Task
	err := collection.FindOne(context.TODO(), bson.M{"id": task.ID}).Decode(&existingTask)
	if err == nil {
		// Task with the same ID already exists
		return false
	}

	// Proceed to insert the new task
	_, err = collection.InsertOne(context.TODO(), task)
	return err == nil
}

func DeleteTask(id string) (*models.Task, bool) {
	var task models.Task
	err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&task)
	if err != nil {
		return nil, false
	}

	_, err = collection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		return nil, false
	}
	return &task, true
}

func connectDb() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://<>username:<password>@cluster0.tj8um.mongodb.net/?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}
