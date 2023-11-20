package controllers

import (
	"fmt"
	"main/config"
	"main/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary			Get Blogs.
// @Description		Return list of blog.
// @Tags			Blog
// @Router			/blog [get]
func GetBlogs(c *gin.Context) {
	var blogs []models.Blog
	config.GetDB().Find(&blogs)
	c.JSON(http.StatusOK, gin.H{"result": blogs})
}

// @Summary			Create Blogs.
// @Description		Return blog data.
// @Tags			Blog
// @Param request body forms.Blog true "Request body for creating a resource"
// @Router			/blog [post]
func CreateBlogs(c *gin.Context) {
	var blog models.Blog
	var err error

	err = c.ShouldBindJSON(&blog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err = config.GetDB().Create(&blog).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"result": blog})
}

// @Summary			Update Blogs.
// @Description		Return new blog after update data.
// @Tags			Blog
// @Param id path int true "ID of models.Blog"
// @Param request body forms.Blog true "Request body for update a resource"
// @Router			/blog/{id} [put]
func UpdateBlogs(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data, blog models.Blog

	r := config.GetDB().First(&blog, id)
	if r.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": r.Error.Error()})
		return
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	blog.Title = data.Title
	blog.Author = data.Author
	config.GetDB().Save(&blog)

	c.JSON(http.StatusCreated, gin.H{"result": blog})
}

// @Summary			Delete Blogs.
// @Description		Return no context status.
// @Tags			Blog
// @Param id path int true "ID of models.Blog"
// @Router			/blog/{id} [delete]
func DeletBlogs(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var blog models.Blog

	if err := config.GetDB().First(&blog, id).Error; err != nil {
		fmt.Println(err)
	}

	c.Status(http.StatusNoContent)
}