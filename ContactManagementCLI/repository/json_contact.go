package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/AaronBrownDev/ContactManagementCLI/domain"
)

type jsonContactRepository struct {
	jsonContacts []map[string]string
}

func GetJsonContactRepository(jsonContacts []map[string]string) domain.ContactRepository {
	return &jsonContactRepository{jsonContacts: jsonContacts}
}

// Create implements domain.ContactRepository.
func (r *jsonContactRepository) Create(name string, phoneNumber string, emailAddress string) (err error) {

	newContact := map[string]string{
		"name":         name,
		"phoneNumber":  phoneNumber,
		"emailAddress": emailAddress,
	}

	r.jsonContacts = append(r.jsonContacts, newContact)

	updatedBytes, err := json.MarshalIndent(r.jsonContacts, "", " ")
	if err != nil {
		return err
	}

	if err = os.WriteFile("contacts.json", updatedBytes, 0644); err != nil {
		return err
	}

	return nil
}

// Delete implements domain.ContactRepository.
func (r jsonContactRepository) Delete(contactID int) (err error) {

	if len(r.jsonContacts) <= contactID {
		return fmt.Errorf("invalid contactID: the contactID goes up to %d", len(r.jsonContacts) - 1)
	}

	r.jsonContacts = append(r.jsonContacts[:contactID], r.jsonContacts[contactID+1:]...)

	updatedBytes, err := json.MarshalIndent(r.jsonContacts, "", " ")
	if err != nil {
		return err
	}

	if err = os.WriteFile("contacts.json", updatedBytes, 0644); err != nil {
		return err
	}

	return
}

// GetAll implements domain.ContactRepository.
func (r jsonContactRepository) GetAll() (contacts []domain.Contact, err error) {

	for i, jsonContact := range r.jsonContacts {
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

	if len(r.jsonContacts) <= contactID {
		return domain.Contact{}, fmt.Errorf("invalid contactID: the contactID goes up to %d", len(r.jsonContacts) - 1)
	}

	return domain.Contact{
		ContactID:    contactID,
		Name:         r.jsonContacts[contactID]["name"],
		PhoneNumber:  r.jsonContacts[contactID]["phoneNumber"],
		EmailAddress: r.jsonContacts[contactID]["emailAddress"],
	}, err
}

// GetByName implements domain.ContactRepository.
func (r jsonContactRepository) GetByName(name string) (contacts []domain.Contact, err error) {

	for i, jsonContact := range r.jsonContacts {

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
func (r jsonContactRepository) Update(contact domain.Contact) (err error) {

	if len(r.jsonContacts) <= contact.ContactID {
		return fmt.Errorf("invalid contactID: the contactID goes up to %d", len(r.jsonContacts) - 1)
	}

	r.jsonContacts[contact.ContactID]["name"] = contact.Name
	r.jsonContacts[contact.ContactID]["phoneNumber"] = contact.PhoneNumber
	r.jsonContacts[contact.ContactID]["emailAddress"] = contact.EmailAddress

	updatedBytes, err := json.MarshalIndent(r.jsonContacts, "", " ")
	if err != nil {
		return err
	}

	if err = os.WriteFile("contacts.json", updatedBytes, 0644); err != nil {
		return err
	}

	return
}
