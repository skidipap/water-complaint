package api

import (
	"example/complaint_service/pkg/handlers"
	"example/complaint_service/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// Group API routes under /api/v1
	v1Router := router.Group("/api/v1")
	{
		// Register handler for user registration and login
		v1Router.POST("/register", handlers.RegisterHandler)
		v1Router.POST("/login", handlers.LoginHandler)

	}

	// Gruop API for user
	userRouter := v1Router.Group("/user")
	// Use JWT Middleware
	userRouter.Use(middleware.JWTAuth)
	{
		userRouter.DELETE("/", handlers.DeleteUser)
		userRouter.GET("/", handlers.ListAllUsers)
	}

	// Gruop API for Complaint
	complaintRouter := v1Router.Group("/complaint")
	// Use JWT Middleware
	complaintRouter.Use(middleware.JWTAuth)
	{
		complaintRouter.GET("/", handlers.ListAllComplaintsHandler)
		complaintRouter.POST("/", handlers.SubmitComplaintHandler)
		complaintRouter.DELETE("/:id", handlers.DeleteComplaintHandler)
		complaintRouter.PUT("/:id", handlers.UpdateComplaintHandler)
	}
	return router
}
