package main

import (
	"github.com/MeLiger/backpack-bcgow6-german-torres/GoTesting/Arquitectura.GoWeb/cmd/server/handler"
	"github.com/MeLiger/backpack-bcgow6-german-torres/GoTesting/Arquitectura.GoWeb/internal/users"
	"github.com/MeLiger/backpack-bcgow6-german-torres/GoTesting/Arquitectura.GoWeb/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

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
	userRouter.PATCH("/:id", userController.UpdateNameAge)
	userRouter.DELETE("/:id", userController.Delete)
	router.Run()
}
