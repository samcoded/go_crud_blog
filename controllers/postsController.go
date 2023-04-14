package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/samcoded/go_crud_blog/initializers"
	"github.com/samcoded/go_crud_blog/models"
)

func PostsCreate(c *gin.Context) {
	// request
	var body struct {
		Title   string
		Content string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "Error",
		})
		return
	}

	// CREATE POST
	post := models.Post{Title: body.Title, Content: body.Content}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Error",
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	var post models.Post
	initializers.DB.First(&post, c.Param("id"))

	// SEND ERROR NO POST
	if post.ID == 0 {
		c.JSON(404, gin.H{
			"message": "Post not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	var post models.Post
	initializers.DB.First(&post, c.Param("id"))

	// SEND ERROR NO POST
	if post.ID == 0 {
		c.JSON(404, gin.H{
			"message": "Post not found",
		})
		return
	}

	// request

	var body struct {
		Title   string
		Content string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "Error",
		})
		return
	}

	post.Title = body.Title
	post.Content = body.Content

	initializers.DB.Save(&post)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	var post models.Post
	initializers.DB.First(&post, c.Param("id"))

	// SEND ERROR NO POST
	if post.ID == 0 {
		c.JSON(404, gin.H{
			"message": "Post not found",
		})
		return
	}

	// delete

	initializers.DB.Delete(&post)

	c.JSON(200, gin.H{
		"message": "Post deleted",
	})
}
