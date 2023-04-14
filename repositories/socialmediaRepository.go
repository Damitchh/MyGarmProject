package repositories

import (
	"MygarmProject/models"
	"gorm.io/gorm"
)

type SocialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *SocialMediaRepository {
	return &SocialMediaRepository{db}
}

func (repo *SocialMediaRepository) CreateSocialMedia(socialmedia *models.SocialMedia) error {
	return repo.db.Debug().Create(socialmedia).Error
}

func (repo *SocialMediaRepository) GetSocialMediaByID(socialmediaID uint) (models.SocialMedia, error) {
	var results models.SocialMedia
	err := repo.db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username")
	}).Find(&results, "id = ?", socialmediaID).Error
	return results, err
}

func (repo *SocialMediaRepository) GetAllSocialMedias() ([]models.SocialMedia, error) {
	var results []models.SocialMedia
	err := repo.db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, age")
	}).Find(&results).Error
	return results, err
}

func (repo *SocialMediaRepository) DeleteSocialMediaByID(socialmediaID uint, userID uint) (int64, error) {
	result := repo.db.Debug().Delete(&models.SocialMedia{}, socialmediaID).Where("user_id = ?", userID)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (repo *SocialMediaRepository) UpdateSocialMediaByID(socialmedia *models.SocialMedia, userID uint) (int64, error) {
	result := repo.db.Model(socialmedia).Where("id = ?", socialmedia.ID).Updates(models.SocialMedia{Name: socialmedia.Name, SocialMediaUrl: socialmedia.SocialMediaUrl}).Where("user_id = ?", userID)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
