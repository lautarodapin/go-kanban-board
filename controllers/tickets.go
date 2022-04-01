package controllers

import (
	"kanban-board/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all tickets
// @Schemes
// @Description Returns all tickets
// @Tags Tickets list
// @Accept json
// @Produce json
// @Success 200 {array} []models.Ticket
// @Failure 404 {string} error
// @Router /tickets [get]
func GetTickets() gin.HandlerFunc {
	return func(c *gin.Context) {
		tickets, err := models.TicketManager.GetAll()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusOK, tickets)
	}
}

// @Summary Get a ticket by id
// @Schemes
// @Description Returns a ticket by id
// @Tags Get Ticket
// @Accept json
// @Produce json
// @Param id path string true "Ticket ID"
// @Success 200 {object} models.Ticket
// @Failure 404 {string} error
// @Router /tickets/:id [get]
func GetTicket() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		ticket, err := models.TicketManager.GetById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, ticket)
	}
}

// @Summary Create a new ticket
// @Schemes
// @Description Create a new ticket
// @Tags Create Ticket
// @Accept json
// @Produce json
// @Param ticket body models.Ticket true "Ticket"
// @Success 200 {object} models.Ticket
// @Failure 404 {string} error
// @Router /tickets [post]
func CreateTicket() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ticket models.Ticket
		err := ctx.ShouldBindJSON(&ticket)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		err = models.TicketManager.Create(&ticket)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, ticket)
	}
}

// @Summary Update a ticket
// @Schemes
// @Description Update a ticket
// @Tags Update Ticket
// @Accept json
// @Produce json
// @Param id path string true "Ticket ID"
// @Param ticket body models.Ticket true "Ticket"
// @Success 200 {object} models.Ticket
// @Failure 404 {string} error
// @Router /tickets/:id [put]
func UpdateTicket() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ticket models.Ticket
		id := ctx.Param("id")
		err := ctx.ShouldBindJSON(&ticket)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}
		err = models.TicketManager.Update(&ticket, id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, ticket)
	}
}
