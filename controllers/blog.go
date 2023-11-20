package controllers

import (
	"main/config"
	"main/models"
	"main/utils"
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
// @Param file formData file true "File to upload"
// @Param title formData string true "Title data"
// @Param author formData string true "Author data"
// @Router			/blog [post]
func CreateBlogs(c *gin.Context) {
	blog := models.Blog{}

	blog.Title = c.PostForm("title")
	blog.Author = c.PostForm("author")
	blog.File = utils.UploadFile(c)

	config.GetDB().Create(&blog)
	c.JSON(http.StatusCreated, gin.H{"result": blog})
}

// @Summary			Update Blogs.
// @Description		Return new blog after update data.
// @Tags			Blog
// @Param id path int true "ID of models.Blog"
// @Param file formData file true "File"
// @Param title formData string true "Title"
// @Param author formData string true "Author"
// @Router			/blog/{id} [put]
func UpdateBlogs(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var blog models.Blog

	r := config.GetDB().First(&blog, id)
	if r.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": r.Error.Error()})
		return
	}

	utils.RemoveFile(blog.Title)

	blog.Title = c.PostForm("title")
	blog.Author = c.PostForm("author")
	blog.File = utils.UploadFile(c)
	config.GetDB().Save(&blog)

	c.JSON(http.StatusCreated, gin.H{"result": blog})
}

// @Summary			Delete Blogs.
// @Description		Return no context status.
// @Tags			Blog
// @Param id path int true "ID of models.Blog"
// @Router			/blog/{id} [delete]
func DeletBlogs(c *gin.Context) {
	id := c.Param("id")
	var blog models.Blog

	r := config.GetDB().First(&blog, id)
	if r.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": r.Error.Error()})
		return
	}

	config.GetDB().Delete(&blog, id)
	c.Status(http.StatusNoContent)
}
