package repository

import (
	"errors"
	"fmt"
	"github.com/ktruedat/taleBack/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	dbClient *gorm.DB
}

func NewUserRepository(dbClient *gorm.DB) *UserRepository {
	return &UserRepository{
		dbClient: dbClient,
	}
}

func (repo *UserRepository) Create(user *model.User) (uint, error) {
	err := repo.dbClient.Debug().Model(model.User{}).Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (repo *UserRepository) Get(email string) (*model.User, error) {
	user := &model.User{}
	err := repo.dbClient.Debug().Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Update(user *model.User) error {
	existingUser := &model.User{}

	err := repo.dbClient.Where("email = ?", user.ID).First(&existingUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with ID %d not found", user.ID)
		}
		return err
	}
	existingUser.Password = user.Password

	err = repo.dbClient.Save(&existingUser).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) Delete(email string) (model.User, error) {
	var user model.User
	err := repo.dbClient.Debug().Model(model.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		return model.User{}, err
	}

	err = repo.dbClient.Debug().Delete(&user).Error

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
