package router

import (
	"task-manager-go/controllers"
	"task-manager-go/middleware"
	"task-manager-go/models"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Auth routes
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	// User management
	router.PATCH("/promote/:username",
		middleware.AuthMiddleware(),
		middleware.RoleMiddleware(models.AdminRole.String()),
		controllers.UpdateRole,
	)

	// Task routes
	taskRoutes := router.Group("/tasks", middleware.AuthMiddleware())
	{
		taskRoutes.GET("", controllers.GetTasks)
		taskRoutes.GET("/:id", controllers.GetTaskById)
		taskRoutes.POST("",
			middleware.RoleMiddleware(models.AdminRole.String()),
			controllers.CreateTask,
		)
		taskRoutes.PUT("/:id",
			middleware.RoleMiddleware(models.AdminRole.String()),
			controllers.UpdateTask,
		)
		taskRoutes.DELETE("/:id",
			middleware.RoleMiddleware(models.AdminRole.String()),
			controllers.DeleteTask,
		)
	}

	return router
}
