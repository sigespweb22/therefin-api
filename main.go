package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Cria um router
	router := gin.Default()

	// Rota raiz
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok retornado!\n")
	})

	// Rota que retorna um nome
	router.GET("api/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Meu nome é, %s\n", name)
	})

	router.Run(":8080")
}
