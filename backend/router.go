package backend

import (
"github.com/gin-gonic/gin"
	"net/http"
)

func setRouter() *gin.Engine {
	// Creates default gin router with Logger and Recovery middleware already attached
	router := gin.Default()

	// Create API route group
	api := router.Group("/api")
	{
		// Add /hello GET route to router and define route handler function
		api.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"msg": "world"})
		})

		api.POST("/signup", signUp)
		api.POST("/signin", signIn)


	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}