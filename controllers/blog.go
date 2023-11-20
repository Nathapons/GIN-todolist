package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary			Get Blogs.
// @Description		Return list of blog.
// @Tags			Blog
// @Router			/blog [get]
func GetBlogs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}


// @Summary			Get Blogs.
// @Description		Return list of blog.
// @Tags			Blog
// @Router			/blog [post]
func CreateBlogs(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "ok"})
}


// @Summary			Get Blogs.
// @Description		Return list of blog.
// @Tags			Blog
// @Router			/blog/:id [put]
func UpdateBlogs(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "ok"})
}

// @Summary			Get Blogs.
// @Description		Return list of blog.
// @Tags			Blog
// @Router			/blog/:id [delete]
func DeletBlogs(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "ok"})
}
