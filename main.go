package main

import (
	"kanban-board/controllers"
	docs "kanban-board/docs"
	"kanban-board/models"

	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func main() {
	godotenv.Load()

	models.ConnectDatabase()
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	kanbans := r.Group("/kanbans")
	{
		kanbans.GET("/", controllers.GetKanbans())
		kanbans.POST("/", controllers.CreateKanban())
		kanbans.GET("/:id", controllers.GetKanban())
		kanbans.PUT("/:id", controllers.UpdateKanban())
	}
	dropzones := r.Group("/dropzones")
	{
		dropzones.GET("/", controllers.GetDropzones())
		dropzones.POST("/", controllers.CreateDropzone())
		dropzones.GET("/:id", controllers.GetDropzone())
		dropzones.PUT("/:id", controllers.UpdateDropzone())
	}
	columns := r.Group("/columns")
	{
		columns.GET("/", controllers.GetColumns())
		columns.POST("/", controllers.CreateColumn())
		columns.GET("/:id", controllers.GetColumn())
		columns.PUT("/:id", controllers.UpdateColumn())
	}
	tickets := r.Group("/tickets")
	{
		tickets.GET("/", controllers.GetTickets())
		tickets.POST("/", controllers.CreateTicket())
		tickets.GET("/:id", controllers.GetTicket())
		tickets.PUT("/:id", controllers.UpdateTicket())
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
