package repository

import (
	"github.com/ktruedat/taleBack/internal/model"
	"gorm.io/gorm"
)

type TaleRepository struct {
	dbClient *gorm.DB
}

func NewTaleRepository(dbClient *gorm.DB) *TaleRepository {
	return &TaleRepository{dbClient: dbClient}
}

func (repo *TaleRepository) Create(tale *model.Tale) (uint, error) {
	if err := repo.dbClient.Debug().Model(model.Tale{}).Create(tale).Error; err != nil {
		return 0, err
	}
	return tale.ID, nil
}

func (repo *TaleRepository) GetAll() ([]model.Tale, error) {
	var tales []model.Tale
	if err := repo.dbClient.Debug().Find(&tales).Error; err != nil {
		return nil, err
	}
	return tales, nil
}

func (repo *TaleRepository) Update(tale *model.Tale) error {
	existingTale := &model.Tale{}
	if err := repo.dbClient.Debug().Where("title = ?", tale.ID).First(&existingTale).Error; err != nil {
		return err
	}
	existingTale.CategoryID = tale.CategoryID
	existingTale.Contents = tale.Contents
	if err := repo.dbClient.Debug().Save(&existingTale).Error; err != nil {
		return err
	}
	return nil
}

func (repo *TaleRepository) Delete(title string) (model.Tale, error) {
	var tale model.Tale
	if err := repo.dbClient.Debug().Model(model.Tale{}).Where("title = ?", title).First(&tale).Error; err != nil {
		return model.Tale{}, err
	}
	if err := repo.dbClient.Debug().Delete(&tale).Error; err != nil {
		return model.Tale{}, err
	}
	return tale, nil
}
