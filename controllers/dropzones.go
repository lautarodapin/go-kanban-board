package controllers

import (
	"kanban-board/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all dropzones
// @Schemes
// @Description Returns all dropzones
// @Tags Dropzones list
// @Accept json
// @Produce json
// @Success 200 {array} []models.Dropzone
// @Failure 404 {string} error
// @Router /dropzones [get]
func GetDropzones() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dropzones []models.Dropzone
		err := models.DB.Find(&dropzones).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusOK, dropzones)
	}
}

// @Summary Get a dropzone by id
// @Schemes
// @Description Returns a dropzone by id
// @Tags Get Dropzone
// @Accept json
// @Produce json
// @Param id path string true "Dropzone ID"
// @Success 200 {object} models.Dropzone
// @Failure 404 {string} error
// @Router /dropzones/:id [get]
func GetDropzone() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dropzone models.Dropzone
		id := ctx.Param("id")
		err := models.DB.Preload("Column").First(&dropzone, id).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, dropzone)
	}
}

// @Summary Create a new dropzone
// @Schemes
// @Description Create a new dropzone
// @Tags Create Dropzone
// @Accept json
// @Produce json
// @Param dropzone body models.BaseDropzone true "Dropzone"
// @Success 200 {object} models.Dropzone
// @Failure 404 {string} error
// @Router /dropzones [post]
func CreateDropzone() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dropzone models.BaseDropzone
		err := ctx.BindJSON(&dropzone)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		err = models.DB.Create(&dropzone).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, dropzone)
	}
}

// @Summary Update a dropzone
// @Schemes
// @Description Update a dropzone
// @Tags Update Dropzone
// @Accept json
// @Produce json
// @Param id path string true "Dropzone ID"
// @Param dropzone body models.BaseDropzone true "Dropzone"
// @Success 200 {object} models.Dropzone
// @Failure 404 {string} error
// @Router /dropzones/:id [put]
func UpdateDropzone() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dropzone models.BaseDropzone
		id := ctx.Param("id")
		err := ctx.ShouldBindJSON(&dropzone)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		err = models.DB.Model(models.Dropzone{}).Where("id = ?", id).Updates(dropzone).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, dropzone)
	}
}