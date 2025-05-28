package services

import (
	"addressbook/internal/models"
	"addressbook/internal/repositories"
)

type ContactService interface {
	CreateContact(contact *models.Contact) error
	GetContactByID(id uint) (*models.Contact, error)
	GetAllContacts() ([]models.Contact, error)
	UpdateContact(contact *models.Contact) error
	DeleteContact(id uint) error
}

type contactService struct {
	repo repositories.ContactRepository
}

func NewContactService(repo repositories.ContactRepository) ContactService {
	return &contactService{repo}
}

func (s *contactService) CreateContact(contact *models.Contact) error {
	return s.repo.Create(contact)
}

func (s *contactService) GetContactByID(id uint) (*models.Contact, error) {
	return s.repo.GetByID(id)
}

func (s *contactService) GetAllContacts() ([]models.Contact, error) {
	return s.repo.GetAll()
}

func (s *contactService) UpdateContact(contact *models.Contact) error {
	return s.repo.Update(contact)
}

func (s *contactService) DeleteContact(id uint) error {
	return s.repo.DeleteContact(id)
}
