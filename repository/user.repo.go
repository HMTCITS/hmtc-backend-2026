package repository

import (
	"github.com/HMTCITS/hmtc-backend-2025/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user model.User) (model.User, error)
	IsNRPExist(nrp string) (bool, error)
	FindUserByNRP(nrp string) (model.User, error)
	FindUserById(id string) (model.User, error)
	FindDepartementByName(name string) (*model.Departement, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) Register(user model.User) (model.User, error) {
	if err := ur.DB.Create(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

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

func (r *userRepository) FindUserByNRP(nrp string) (model.User, error) {
	var user model.User
	if err := r.DB.Where("nrp = ?", nrp).Take(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.User{}, err
		}
		return model.User{}, err
	}

	return user, nil
}

func (r *userRepository) IsNRPExist(nrp string) (bool, error) {
	var user model.User
	if err := r.DB.Where("nrp = ?", nrp).Take(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, err
		}
		return false, err
	}

	return true, nil
}

func (r *userRepository) FindDepartementByName(name string) (*model.Departement, error) {
	var departement model.Departement
	err := r.DB.Where("name = ?", name).First(&departement).Error
	if err != nil {
		return nil, err
	}
	return &departement, nil
}
