package usecases

import (
	"Chinnasit/entities"
	"errors"
)

type UserUseCase interface {
	CreateUser(user entities.User) error
	GetUsers() ([]entities.User, error)
	UpdateUserFull(user entities.User) error
	UpdateUserPartial(id uint, data map[string]interface{}) error
	DeleteUser(id uint) error
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserUseCase {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user entities.User) error {
	if user.FirstName == nil || user.LastName == nil || user.Email == nil || user.RoleId == nil {
		return errors.New("invalid request: Some field is empty")
	}

	if *user.RoleId > 3 {
		return errors.New("role id must be in range [0,3]")
	}

	return s.repo.Create(user)
}

func (s *UserService) GetUsers() ([]entities.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) UpdateUserFull(user entities.User) error {
	if user.FirstName == nil || user.LastName == nil || user.Email == nil || user.RoleId == nil {
		return errors.New("invalid request: Some field is empty")
	}

	if *user.RoleId > 3 {
		return errors.New("invalid request: RoleId must be between [0,3]")
	}

	return s.repo.UpdateFull(user)
}

func (s *UserService) UpdateUserPartial(id uint, data map[string]interface{}) error {
	if len(data) == 0 {
		return errors.New("invalid request: empty data")
	}

	if data["RoleId"] != nil {
		if data["RoleId"].(float64) > float64(3) {
			return errors.New("invalid request: RoleId must be between [0,3]")
		}
	}

	return s.repo.UpdatePatrial(id, data)
}

func (s *UserService) DeleteUser(id uint) error {
	if id < 0 {
		return errors.New("invalid request: ID must be more than 0")
	}

	return s.repo.Delete(id)
}
