package handlers

import (
	"addressbook/internal/models"
	"addressbook/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	service services.ContactService
}

func NewContactHandler(service services.ContactService) *ContactHandler {
	return &ContactHandler{service}
}
func (h *ContactHandler) CreateContact(c *gin.Context) {
	var contact models.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateContact(&contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, contact)

}

func (h *ContactHandler) GetContactByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	contact, err := h.service.GetContactByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}
	c.JSON(http.StatusOK, contact)
}

func (h *ContactHandler) GetAllContact(c *gin.Context) {
	contact, err := h.service.GetAllContacts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contact)
}

// func (h *ContactHandler) UpdateContact(c *gin.Context) {
// 	var contact models.Contact
// 	if err := c.ShouldBindJSON(&contact); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	if err := h.service.UpdateContact(&contact); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, contact)
// }

func (h *ContactHandler) UpdateContact(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact ID"})
		return
	}

	var contact models.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contact.ID = uint(id) // this is critical

	if err := h.service.UpdateContact(&contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contact)
}

func (h *ContactHandler) DeleteContact(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.service.DeleteContact(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
