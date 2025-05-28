package app

import (
	v1 "addressbook/api/v1"
	"addressbook/config"
	"addressbook/internal/handlers"
	"addressbook/internal/repositories"
	"addressbook/internal/services"
	"addressbook/internal/utils"

	// "addressbook/internal/utils"
	// "log"

	"github.com/gin-gonic/gin"
)

func Run(config *config.Config) error {
	// Connect to the database
	db, err := utils.ConnectDB(config)
	if err != nil {
		return err
	}

	// // Auto-migrate models
	// if err := db.AutoMigrate(&models.Contact{}); err != nil {
	// 	log.Fatalf("AutoMigrate failed: %v", err)
	// }

	// Setup Gin
	router := gin.Default()

	contactRepository := repositories.NewContactRepository(db)
	contactService := services.NewContactService(contactRepository)
	contactHandler := handlers.NewContactHandler(contactService)
	// Register routes
	v1.SetupRoutes(router, contactHandler)

	// Run server
	return router.Run(":8080") //default port
}
