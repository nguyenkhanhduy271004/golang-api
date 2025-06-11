package main

import (
	"github.com/gin-gonic/gin"
	v1Handler "nguyenduy.com/golang-api/internal/api/v1/handler"
)

func main() {
	r := gin.Default()

	userHandler := v1Handler.NewUserHandler()
	productHandler := v1Handler.NewProductHandler()

	v1 := r.Group("/api/v1")
	{
		userApi := v1.Group("/users")
		{
			userApi.GET("/", userHandler.GetUsersV1)
			userApi.POST("/", userHandler.CreateUserV1)
			userApi.PUT("/:id", userHandler.UpdateUserV1)
			userApi.DELETE("/:id", userHandler.DeleteUserV1)
		}

		productApi := v1.Group("/products")
		{
			productApi.GET("/", productHandler.GetProductsV1)
			productApi.POST("/", productHandler.CreateProductV1)
			productApi.PUT("/:id", productHandler.UpdateProductV1)
			productApi.DELETE("/:id", productHandler.DeleteProductV1)
		}
	}

	r.Run(":8080")
}
