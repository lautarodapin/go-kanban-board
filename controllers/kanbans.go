package controllers

import (
	"kanban-board/models"
	"net/http"
	"time"

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
		kanbans, err := models.KanbanManager.GetAll()
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
		id := ctx.GetUint("id")
		kanban, err := models.KanbanManager.GetById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, kanban)
	}
}

type KanbanBody struct {
	Name      string    `json:"name" binding:"required"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required"`
}

// @Summary Create a new kanban
// @Schemes
// @Description Create a new kanban
// @Tags CreateKanban
// @Accept json
// @Produce json
// @Param kanban body KanbanBody true "Kanban"
// @Success 200 {object} models.Kanban
// @Failure 404 {string} error
// @Router /kanbans [post]
func CreateKanban() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body KanbanBody
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		kanban := models.Kanban{
			Name:      body.Name,
			StartDate: body.StartDate,
			EndDate:   body.EndDate,
		}
		if err := models.KanbanManager.Create(&kanban); err != nil {
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
// @Param kanban body KanbanBody true "Kanban"
// @Success 200 {object} models.Kanban
// @Failure 404 {string} error
// @Router /kanbans/:id [put]
func UpdateKanban() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body KanbanBody
		id := ctx.GetUint("id")
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		kanban := models.Kanban{
			Name:      body.Name,
			StartDate: body.StartDate,
			EndDate:   body.EndDate,
		}
		if err := models.KanbanManager.Update(&kanban, id); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, kanban)
	}
}
