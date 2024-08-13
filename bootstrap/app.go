package bootstrap

import (
	"log"
	"task/api/controller"
	"task/api/middleware"
	"task/api/router"
	"task/repository"
	"task/usecase"

	"github.com/gin-gonic/gin"
)

// StartApp initializes the application and starts the server
func StartApp() {
	// Load environment variables
	env := LoadEnv()

	// Initialize the database connection
	db, err := ConnectDB(env)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Repositories
	userRepo := repository.NewUserRepository(db)
	taskRepo := repository.NewTaskRepository(db)

	// Usecases
	loginUsecase := usecase.NewLoginUsecase(userRepo)
	registerUsecase := usecase.NewRegisterUsecase(userRepo)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)

	// Controllers
	loginController := controller.NewLoginController(loginUsecase)
	signUpController := controller.NewSignUpController(registerUsecase)
	taskController := controller.NewTaskController(taskUsecase)

	// Middleware
	authMiddleware := middleware.NewAuthMiddleware(loginUsecase)

	// Initialize Gin router
	r := gin.Default()

	// Setup routes
	router.SetupRouter(r, signUpController, loginController, taskController, authMiddleware)

	// Start server
	if err := r.Run(":" + env.ServerPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
