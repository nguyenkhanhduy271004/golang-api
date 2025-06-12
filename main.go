package main

import (
	"github.com/gin-gonic/gin"
	v1Handler "nguyenduy.com/golang-api/internal/api/v1/handler"
)

func main() {
	r := gin.Default()

	userHandler := v1Handler.NewUserHandler()
	productHandlerV1 := v1Handler.NewProductHandler()

	v1 := r.Group("/api/v1")
	{
		userApi := v1.Group("/users")
		{
			userApi.GET("/", userHandler.GetUsersV1)
			// userApi.GET("/:uuid", userHandler.GetUsersByUuidV1)
			userApi.GET("/:id", userHandler.GetUserByIdV1)
			userApi.POST("/", userHandler.CreateUserV1)
			userApi.PUT("/:id", userHandler.UpdateUserV1)
			userApi.DELETE("/:id", userHandler.DeleteUserV1)
		}

		productApi := v1.Group("/products")
		{
			productApi.GET("", productHandlerV1.GetProductsV1)
			productApi.GET("/:slug", productHandlerV1.GetProductsBySlugV1)
			productApi.POST("", productHandlerV1.PostProductsV1)
			productApi.PUT("/:id", productHandlerV1.PutProductsV1)
			productApi.DELETE("/:id", productHandlerV1.DeleteProductsV1)
		}

	}

	r.Run(":8080")
}
