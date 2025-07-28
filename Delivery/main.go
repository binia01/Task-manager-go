package main

import (
	"context"
	routers "task-manager-go/Delivery/router"
	"task-manager-go/Infrastructure"
	"task-manager-go/Repositories"
	"task-manager-go/Usecases"

	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

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

	taskCollection := client.Database("Task-Database").Collection("Tasks")
	userCollection := client.Database("Task-Database").Collection("Users")

	passwordService := Infrastructure.NewPasswordService()
	jwtSecret := []byte("wN7z@JrV3uK!dXp2qT$eGh8yF9cLb6mZ")
	jwtService := Infrastructure.NewJWTService(string(jwtSecret))

	authMiddleware := Infrastructure.AuthMiddleware(jwtSecret)

	taskRepo := Repositories.NewTaskRepository(taskCollection)
	userRepo := Repositories.NewUserRepository(userCollection)

	taskUsecase := Usecases.NewTaskUsecase(taskRepo)
	userUsecase := Usecases.NewUserUsecase(userRepo, passwordService, jwtService)

	r := routers.SetupRouter(*taskUsecase, *userUsecase, authMiddleware)
	r.Run(":8080")
}
