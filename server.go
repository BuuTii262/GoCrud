package main

import (
	"net/http"

	"exampl.com/goCrud/config"
	"exampl.com/goCrud/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := config.OpenDB()

	//It will do before every route start
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Min Ga Lar Par!"})
	})

	r.GET("/api/users", controller.GetAllUsers)
	r.POST("/api/users", controller.CreateUser)
	r.GET("/api/users/:id", controller.FindUser)
	r.PATCH("/api/users/:id", controller.UpdateUser)
	r.DELETE("/api/users/:id", controller.DeleteUser)

	r.Run()
}
