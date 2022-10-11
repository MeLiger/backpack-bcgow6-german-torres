package main

import (
	"github.com/MeLiger/backpack-bcgow6-german-torres/Arquitectura.GoWeb/cmd/server/handler"

	"github.com/MeLiger/backpack-bcgow6-german-torres/Arquitectura.GoWeb/internal/users"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()
	db := store.New(store.FileType, "./users.json")
	userRepository := users.NewRepository(db)
	userService := users.NewService(userRepository)
	userController := handler.NewUserController(userService)

	router := gin.Default()

	userRouter := router.Group("/users")
	userRouter.GET("/", userController.GetAll)
	userRouter.POST("/", userController.Store)
	userRouter.PUT("/:id", userController.Update)
	userRouter.PATCH("/id", userController.UpdateNameAge)
	userRouter.DELETE("/:id", userController.Delete)
	router.Run()
}
