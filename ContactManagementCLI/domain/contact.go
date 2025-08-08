package domain

type Contact struct {
	ContactID    int    `json:"contactID"`
	Name         string `json:"name"`
	PhoneNumber  string `json:"phoneNumber"`
	EmailAddress string `json:"emailAddress"`
}

type ContactRepository interface {
	Create(name, phoneNumber, emailAddress string) error

	GetAll() ([]Contact, error)
	GetByID(contactID int) (Contact, error)
	GetByName(name string) ([]Contact, error)

	Update(contact Contact) error

	Delete(contactID int) error
}
