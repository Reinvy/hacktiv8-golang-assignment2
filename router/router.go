package router

import (
	"assignment2/controllers"

	"github.com/gin-gonic/gin"
)

// InitRouter initializes the router
func InitRouter() *gin.Engine {
	// Create a new router
	r := gin.Default()

	// Define the order routes
	orderRoutes := r.Group("/order")
	{

		orderRoutes.POST("/", controllers.CreateOrder)
		orderRoutes.GET("/", controllers.GetOrder)
		orderRoutes.PUT("/:orderId", controllers.UpdateOrder)
		orderRoutes.DELETE("/:orderId", controllers.DeleteOrder)

	}

	// Return the router
	return r
}
