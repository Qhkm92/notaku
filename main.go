package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"notaku/controller"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.NoRoute(notFound)
	api := r.Group("api")
	{
		v1 := api.Group("v1")
		v1.GET("/post/:id", controller.GetPost)
		v1.GET("/posts", controller.GetAllPosts)
		v1.GET("/posts/add", controller.AddPost)
		v1.GET("/posts/update", controller.UpdatePost)
		v1.GET("/posts/delete", controller.DeletePost)
	}
	r.Run()
}

// not found page
func notFound(c *gin.Context) {
	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}

// nomethods commentj
func NoMethods(route *gin.Engine) {
	route.NoMethod(func(c *gin.Context) {
		c.AbortWithStatusJSON(405, "not allowed")
	})
}

type QueryError struct {
	Query string
	Err   error
}
