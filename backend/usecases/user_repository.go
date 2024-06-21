package usecases

import "Chinnasit/entities"

// C R U D = create read update delete
type UserRepository interface {
	Create(user entities.User) error
	UpdateFull(user entities.User) error
	UpdatePatrial(id uint, data map[string]interface{}) error
	Delete(id uint) error
	FindAll() ([]entities.User, error)
}
