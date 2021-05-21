package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func check(c *gin.Context) {
	c.JSON(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"status": "OK",
		},
	)
}

func main() {
	router := gin.Default()
	api := router.Group("/")

	api.GET("/abc", func(c *gin.Context) {

		c.JSON(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"status": "OK",
			},
		)

	})
	api.GET("/", check)

	router.Run(":8000")
}
