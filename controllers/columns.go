package controllers

import (
	"kanban-board/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all columns
// @Schemes
// @Description Returns all columns
// @Tags GetColumns
// @Accept json
// @Produce json
// @Success 200 {array} []models.Column
// @Failure 404 {string} error
// @Router /columns [get]
func GetColumns() gin.HandlerFunc {
	return func(c *gin.Context) {
		columns, err := models.ColumnManager.GetAll()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusOK, columns)
	}
}

// @Summary Get a column by id
// @Schemes
// @Description Returns a column by id
// @Tags GetColumn
// @Accept json
// @Produce json
// @Param id path string true "Column ID"
// @Success 200 {object} models.Column
// @Failure 404 {string} error
// @Router /columns/:id [get]
func GetColumn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var column models.Column
		id := ctx.Param("id")
		err := models.DB.First(&column, id).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, column)
	}
}

// @Summary Create a new column
// @Schemes
// @Description Create a new column
// @Tags CreateColumn
// @Accept json
// @Produce json
// @Param column body models.Column true "Column"
// @Success 200 {object} models.Column
// @Failure 404 {string} error
// @Router /columns [post]
func CreateColumn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var column models.Column
		err := ctx.BindJSON(&column)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		err = models.DB.Create(&column).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, column)
	}
}

// @Summary Update a column
// @Schemes
// @Description Update a column
// @Tags UpdateColumn
// @Accept json
// @Produce json
// @Param id path string true "Column ID"
// @Param column body models.Column true "Column"
// @Success 200 {object} models.Column
// @Failure 404 {string} error
// @Router /columns/:id [put]
func UpdateColumn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var column models.Column
		id := ctx.Param("id")
		err := ctx.ShouldBindJSON(&column)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		err = models.DB.Model(&column).Where("id = ?", id).Updates(column).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, column)
	}
}
