package main

import (
	"context"
	"net/http"
	"log"
	"wonalabs/views"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, status int, template templ.Component) {
	c.Status(status)
	c.Header("Content-Type", "text/html; charset=utf-8")
	err := template.Render(context.Background(), c.Writer)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		Render(c, http.StatusOK, views.Layout())
	})

	// Binding port to 8080 for clean internal container mapping
	log.Println("Server running on http://localhost:8080")
	r.Run(":8080")
}