package repository

import "github.com/AaronBrownDev/ContactManagementCLI/domain"

type jsonContactRepository struct {
	content []byte
}

func GetJsonContactRepository(content []byte) domain.ContactRepository {
	return &jsonContactRepository{content: content}
}

// Create implements domain.ContactRepository.
func (r *jsonContactRepository) Create(name string, phoneNumber string, emailAddress string) error {
	panic("unimplemented")
}

// Delete implements domain.ContactRepository.
func (r jsonContactRepository) Delete(contactID int) error {
	panic("unimplemented")
}

// GetAll implements domain.ContactRepository.
func (r jsonContactRepository) GetAll() ([]domain.Contact, error) {
	panic("unimplemented")
}

// GetByID implements domain.ContactRepository.
func (r jsonContactRepository) GetByID(contactID int) (domain.Contact, error) {
	panic("unimplemented")
}

// GetByName implements domain.ContactRepository.
func (r jsonContactRepository) GetByName(name string) ([]domain.Contact, error) {
	panic("unimplemented")
}

// Update implements domain.ContactRepository.
func (r jsonContactRepository) Update(contact domain.Contact) error {
	panic("unimplemented")
}
