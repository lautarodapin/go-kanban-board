package controllers

import (
	"kanban-board/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetKanbans returns all kanbans

// @Summary get all kanbans
// @Schemes
// @Description Returns all kanbans
// @Tags GetKanbans
// @Accept json
// @Produce json
// @Success 200 {array} []models.Kanban
// @Failure 404 {string} error
// @Router /kanbans [get]
func GetKanbans() gin.HandlerFunc {
	return func(c *gin.Context) {
		var kanbans []models.Kanban
		err := models.DB.Find(&kanbans).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusOK, kanbans)
	}
}

// GetKanban returns a kanban by id

// @Summary Return a kanban by id
// @Schemes
// @Description Return a kanban by id
// @Tags GetKanban
// @Accept json
// @Produce json
// @Param id path string true "Kanban ID"
// @Success 200 {object} models.Kanban
// @Failure 404 {string} error
// @Router /kanbans/:id [get]
func GetKanban() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var kanban models.Kanban
		id := ctx.Param("id")
		err := models.DB.First(&kanban, id).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, kanban)
	}
}

// @Summary Create a new kanban
// @Schemes
// @Description Create a new kanban
// @Tags CreateKanban
// @Accept json
// @Produce json
// @Param kanban body models.CreateKanban true "Kanban"
// @Success 200 {object} models.Kanban
// @Failure 404 {string} error
// @Router /kanbans [post]
func CreateKanban() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var kanban models.CreateKanban
		err := ctx.ShouldBindJSON(&kanban)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		err = models.DB.Model(models.Kanban{}).Create(&kanban).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, kanban)
	}
}

// @Summary Update a kanban
// @Schemes
// @Description Update a kanban
// @Tags UpdateKanban
// @Accept json
// @Produce json
// @Param id path string true "Kanban ID"
// @Param kanban body models.UpdateKanban true "Kanban"
// @Success 200 {object} models.Kanban
// @Failure 404 {string} error
// @Router /kanbans/:id [put]
func UpdateKanban() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var kanban models.UpdateKanban
		id := ctx.Param("id")
		err := ctx.ShouldBindJSON(&kanban)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		err = models.DB.Model(models.Kanban{}).Where("id = ?", id).Updates(&kanban).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, kanban)
	}
}
