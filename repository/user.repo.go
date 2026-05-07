package repository

import (
	"github.com/HMTCITS/hmtc-backend-2025/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	// Register(user model.User) (model.User, error)
	IsEmailExist(email string) (bool, error)
	FindUserByEmail(email string) (model.User, error)
	FindUserById(id string) (model.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

// func (ur *userRepository) Register(user model.User) (model.User, error) {
// 	if err := ur.DB.Create(&user).Error; err != nil {
// 		return model.User{}, err
// 	}

// 	return user, nil
// }

func (r *userRepository) FindUserById(id string) (model.User, error) {
	var user model.User
	if err := r.DB.Where("id = ?", id).Take(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.User{}, err
		}
		return model.User{}, err
	}

	return user, nil
}

func (r *userRepository) FindUserByEmail(email string) (model.User, error) {
	var user model.User
	if err := r.DB.Where("email = ?", email).Take(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.User{}, err
		}
		return model.User{}, err
	}

	return user, nil
}

func (r *userRepository) IsEmailExist(email string) (bool, error) {
	var user model.User
	if err := r.DB.Where("email = ?", email).Take(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}


