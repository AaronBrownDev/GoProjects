package repository

import (
	"encoding/json"
	"os"

	"github.com/AaronBrownDev/ContactManagementCLI/domain"
)

type jsonContactRepository struct {
	content []byte
}

func GetJsonContactRepository(content []byte) domain.ContactRepository {
	return &jsonContactRepository{content: content}
}

// Create implements domain.ContactRepository.
func (r *jsonContactRepository) Create(name string, phoneNumber string, emailAddress string) (err error) {
	var jsonContacts []map[string]string
	err = json.Unmarshal(r.content, &jsonContacts)
	if err != nil {
		return err
	}

	newContact := map[string]string{
		"name":         name,
		"phoneNumber":  phoneNumber,
		"emailAddress": emailAddress,
	}

	jsonContacts = append(jsonContacts, newContact)

	updatedBytes, err := json.MarshalIndent(jsonContacts, "", " ")
	if err != nil {
		return err
	}

	if err = os.WriteFile("contacts.json", updatedBytes, 0644); err != nil {
		return err
	}

	return nil
}

// Delete implements domain.ContactRepository.
func (r jsonContactRepository) Delete(contactID int) error {
	panic("unimplemented")
}

// GetAll implements domain.ContactRepository.
func (r jsonContactRepository) GetAll() (contacts []domain.Contact, err error) {
	var jsonContacts []map[string]string
	err = json.Unmarshal(r.content, &jsonContacts)
	for i, jsonContact := range jsonContacts {
		contacts = append(contacts, domain.Contact{
			ContactID:    i,
			Name:         jsonContact["name"],
			PhoneNumber:  jsonContact["phoneNumber"],
			EmailAddress: jsonContact["emailAddress"],
		})
	}
	return contacts, err
}

// GetByID implements domain.ContactRepository.
func (r jsonContactRepository) GetByID(contactID int) (contact domain.Contact, err error) {
	var jsonContacts []map[string]string
	err = json.Unmarshal(r.content, &jsonContacts)

	return domain.Contact{
		ContactID:    contactID,
		Name:         jsonContacts[contactID]["name"],
		PhoneNumber:  jsonContacts[contactID]["phoneNumber"],
		EmailAddress: jsonContacts[contactID]["emailAddress"],
	}, err
}

// GetByName implements domain.ContactRepository.
func (r jsonContactRepository) GetByName(name string) (contacts []domain.Contact, err error) {
	var jsonContacts []map[string]string
	err = json.Unmarshal(r.content, &jsonContacts)
	for i, jsonContact := range jsonContacts {

		if jsonContact["name"] == name {
			contacts = append(contacts, domain.Contact{
				ContactID:    i,
				Name:         jsonContact["name"],
				PhoneNumber:  jsonContact["phoneNumber"],
				EmailAddress: jsonContact["emailAddress"],
			})
		}

	}
	return contacts, err
}

// Update implements domain.ContactRepository.
func (r jsonContactRepository) Update(contact domain.Contact) error {
	panic("unimplemented")
}
