package main

import (
	"github.com/samcoded/go_crud_blog/initializers"
	"github.com/samcoded/go_crud_blog/models"
)

func init() {

	initializers.LoadEnv()
	initializers.ConnectDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
