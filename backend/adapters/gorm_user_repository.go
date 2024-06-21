package adapters

import (
	"Chinnasit/entities"
	"Chinnasit/usecases"
	"errors"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) usecases.UserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Create(user entities.User) error {
	result := r.db.Where("email = ?", user.Email).First(&user)

	if result.Error != gorm.ErrRecordNotFound {
		return errors.New("email is exist")
	}

	return r.db.Create(&user).Error
}

func (r *GormUserRepository) FindAll() ([]entities.User, error) {
	var users []entities.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *GormUserRepository) UpdateFull(user entities.User) error {
	result := r.db.First(&user, user.ID)

	if result.Error == gorm.ErrRecordNotFound {
		return errors.New("user not found")
	}
	result = r.db.Save(user)

	return result.Error
}

func (r *GormUserRepository) UpdatePatrial(id uint, data map[string]interface{}) error {
	var user entities.User
	result := r.db.First(&user, id)

	if result.Error == gorm.ErrRecordNotFound {
		return errors.New("user not found")
	}

	result = r.db.Model(&entities.User{ID: id}).Updates(data)

	return result.Error
}

func (r *GormUserRepository) Delete(id uint) error {
	var user entities.User
	result := r.db.First(&user, id)

	if result.Error != nil {
		return result.Error
	}

	result = r.db.Delete(&entities.User{ID: id})

	return result.Error
}
