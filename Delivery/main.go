package main

import (
	"context"
	"os"
	routers "task-manager-go/Delivery/router"
	"task-manager-go/Infrastructure"
	"task-manager-go/Repositories"
	"task-manager-go/Usecases"

	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGODB_URI not set in environment")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
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

	jwtSecret := os.Getenv("MONGODB_URI")
	if jwtSecret == "" {
		log.Fatal("Jwt_secret not set in environment")
	}

	jwtService := Infrastructure.NewJWTService(string(jwtSecret))

	authMiddleware := Infrastructure.AuthMiddleware([]byte(jwtSecret))

	taskRepo := Repositories.NewTaskRepository(taskCollection)
	userRepo := Repositories.NewUserRepository(userCollection)

	taskUsecase := Usecases.NewTaskUsecase(taskRepo)
	userUsecase := Usecases.NewUserUsecase(userRepo, passwordService, jwtService)

	r := routers.SetupRouter(*taskUsecase, *userUsecase, authMiddleware)
	r.Run(":8080")
}
