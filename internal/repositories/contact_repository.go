package repositories

import (
	"addressbook/internal/models"

	"gorm.io/gorm"
)

type ContactRepository interface {
	Create(contact *models.Contact) error
	GetByID(id uint) (*models.Contact, error)
	GetAll() ([]models.Contact, error)
	Update(customer *models.Contact) error
	DeleteContact(id uint) error
}

type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{db: db}
}

func (r *contactRepository) Create(contact *models.Contact) error {
	return r.db.Create(contact).Error
}

func (r *contactRepository) GetByID(id uint) (*models.Contact, error) {
	var contact models.Contact
	if err := r.db.First(&contact, id).Error; err != nil {
		return nil, err
	}
	return &contact, nil
}

func (r *contactRepository) GetAll() ([]models.Contact, error) {
	var contacts []models.Contact
	err := r.db.Find(&contacts).Error
	return contacts, err
}

//	func (r *contactRepository) Update(contact *models.Contact) error {
//		return r.db.Save(contact).Error
//	}
func (r *contactRepository) Update(contact *models.Contact) error {
	var existingContact models.Contact

	// Step 1: Fetch the existing contact from DB using ID
	if err := r.db.First(&existingContact, contact.ID).Error; err != nil {
		return err
	}

	// Step 2: Update only non-empty fields
	if contact.FirstName != "" {
		existingContact.FirstName = contact.FirstName
	}
	if contact.LastName != "" {
		existingContact.LastName = contact.LastName
	}
	if contact.Email != "" {
		existingContact.Email = contact.Email
	}
	if contact.Phone != nil {
		existingContact.Phone = contact.Phone
	}
	if contact.AddressLine1 != "" {
		existingContact.AddressLine1 = contact.AddressLine1
	}

	if contact.AddressLine2 != "" {
		existingContact.AddressLine2 = contact.AddressLine2
	}
	if contact.City != "" {
		existingContact.City = contact.City
	}
	if contact.State != "" {
		existingContact.State = contact.State
	}
	if contact.Country != "" {
		existingContact.Country = contact.Country
	}
	if contact.Pincode != nil {
		existingContact.Pincode = contact.Pincode
	}

	// Step 3: Save updated contact
	return r.db.Save(&existingContact).Error
}

//	func (r *contactRepository) DeleteContact(id uint) error {
//		return r.db.Delete(&models.Contact{}, id).Error
//	}
func (r *contactRepository) DeleteContact(id uint) error {
	var contact models.Contact

	// Step 1: Find the contact by ID
	if err := r.db.First(&contact, id).Error; err != nil {
		return err // contact not found
	}

	// Step 2: Soft delete it
	return r.db.Delete(&contact).Error
}
