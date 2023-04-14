package repositories

import (
	"MygarmProject/models"
	"gorm.io/gorm"
)

type PhotoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *PhotoRepository {
	return &PhotoRepository{db}
}

func (repo *PhotoRepository) CreatePhoto(photo *models.Photo) error {
	return repo.db.Debug().Create(photo).Error
}

func (repo *PhotoRepository) GetPhotoByID(photoID uint) (models.Photo, error) {
	var result models.Photo
	err := repo.db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username")
	}).Find(&result, "id = ?", photoID).Error
	return result, err
}

func (repo *PhotoRepository) GetPhotos() ([]models.Photo, error) {
	var results []models.Photo
	err := repo.db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, age")
	}).Find(&results).Error
	return results, err
}

func (repo *PhotoRepository) DeletePhotoByID(photoID uint, userID uint) (int64, error) {
	result := repo.db.Debug().Delete(&models.Photo{}, photoID).Where("user_id = ?", userID)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (repo *PhotoRepository) UpdatePhotoByID(photo *models.Photo, userID uint) (int64, error) {
	result := repo.db.Model(photo).Where("id = ?", photo.ID).Updates(models.Photo{Title: photo.Title, Caption: photo.Caption, PhotoUrl: photo.PhotoUrl}).Where("user_id = ?", userID)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
