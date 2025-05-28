package v1

import (
	"addressbook/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, contactHandler *handlers.ContactHandler) {

	//TO CHECK SERVER IS WORKING
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "WELCOME TO ADDRESS BOOK",
	// 	})
	// })

	v1 := router.Group("/api/v1")
	{
		v1.POST("/contact", contactHandler.CreateContact)
		v1.GET("/contact/:id", contactHandler.GetContactByID)
		v1.GET("/contacts", contactHandler.GetAllContact)
		v1.PUT("/contact/:id", contactHandler.UpdateContact)
		v1.DELETE("/contact/:id", contactHandler.DeleteContact)
	}
}
