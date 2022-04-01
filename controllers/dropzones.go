package controllers

import (
	"kanban-board/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DropzoneBody struct {
	Name     string `json:"name" binding:"required" validate:"required,min=3,max=255"`
	ColumnID uint   `json:"column_id" binding:"required" validate:"required"`
	Order    uint   `json:"order" binding:"required" validate:"required,min=0"`
}

func (body *DropzoneBody) ToModel() models.Dropzone {
	return models.Dropzone{
		Name:     body.Name,
		ColumnID: body.ColumnID,
		Order:    body.Order,
	}
}

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
		dropzones, err := models.DropzoneManager.GetAll()
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
		dropzone, err := models.DropzoneManager.GetById(id)
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
// @Param dropzone body DropzoneBody true "Dropzone"
// @Success 200 {object} models.Dropzone
// @Failure 404 {string} error
// @Router /dropzones [post]
func CreateDropzone() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body DropzoneBody
		if err := ctx.BindJSON(&body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		dropzone := body.ToModel()
		if err := models.DropzoneManager.Create(&dropzone); err != nil {
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
// @Param dropzone body models.Dropzone true "Dropzone"
// @Success 200 {object} models.Dropzone
// @Failure 404 {string} error
// @Router /dropzones/:id [put]
func UpdateDropzone() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body DropzoneBody
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		id := ctx.Param("id")
		dropzone := body.ToModel()
		if err := models.DropzoneManager.Update(&dropzone, id); err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, dropzone)
	}
}
