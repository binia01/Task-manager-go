package routers

import (
	"task-manager-go/Delivery/controllers"
	"task-manager-go/Domain"
	"task-manager-go/Infrastructure"
	"task-manager-go/Repositories"
	"task-manager-go/Usecases"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Initialize repositories
	taskRepo := Repositories.NewTaskRepository()
	userRepo := Repositories.NewUserRepository()

	// Initialize use cases
	taskUsecase := Usecases.NewTaskUsecase(taskRepo)
	userUsecase := Usecases.NewUserUsecase(userRepo)

	// Initialize controllers
	taskController := controllers.NewTaskController(taskUsecase)
	userController := controllers.NewUserController(userUsecase)

	// Auth routes
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	// User management
	router.PATCH("/promote/:username",
		Infrastructure.AuthMiddleware(),
		Infrastructure.RoleMiddleware(Domain.AdminRole.String()),
		userController.UpdateRole,
	)

	// Task routes
	taskRoutes := router.Group("/tasks", Infrastructure.AuthMiddleware())
	{
		taskRoutes.GET("", taskController.GetTasks)
		taskRoutes.GET("/:id", taskController.GetTaskById)
		taskRoutes.POST("",
			Infrastructure.RoleMiddleware(Domain.AdminRole.String()),
			taskController.CreateTask,
		)
		taskRoutes.PUT("/:id",
			Infrastructure.RoleMiddleware(Domain.AdminRole.String()),
			taskController.UpdateTask,
		)
		taskRoutes.DELETE("/:id",
			Infrastructure.RoleMiddleware(Domain.AdminRole.String()),
			taskController.DeleteTask,
		)
	}

	return router
}
