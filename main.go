package main

import (
	"github.com/gin-gonic/gin"
	"github.com/samcoded/go_crud_blog/controllers"
	"github.com/samcoded/go_crud_blog/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()

}

func main() {
	r := gin.Default()
	r.GET("/", controllers.Home)
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.Run() // listen and serve on 0.0.0.0:3000
}
