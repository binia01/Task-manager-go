package routers

import (
	"task-manager-go/Delivery/controllers"
	"task-manager-go/Domain"
	"task-manager-go/Infrastructure"
	"task-manager-go/Usecases"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	taskUC Usecases.TaskUsecase,
	userUC Usecases.UserUsecase,
	authMw gin.HandlerFunc,
) *gin.Engine {

	router := gin.Default()
	taskController := controllers.NewTaskController(taskUC)
	userController := controllers.NewUserController(userUC)
	// Auth routes
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	// User management
	router.PATCH("/promote/:username",
		authMw,
		Infrastructure.RoleMiddleware(Domain.AdminRole.String()),
		userController.UpdateRole,
	)

	// Task routes
	taskRoutes := router.Group("/tasks", authMw)
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
