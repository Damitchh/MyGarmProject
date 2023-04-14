package services

import (
	"MygarmProject/models"
	"MygarmProject/repositories"
)

type PhotoService struct {
	photoRepo *repositories.PhotoRepository
}

func NewPhotoService(photoRepo *repositories.PhotoRepository) *PhotoService {
	return &PhotoService{
		photoRepo: photoRepo,
	}
}

func (s *PhotoService) CreatePhoto(photo *models.Photo) error {
	return s.photoRepo.CreatePhoto(photo)
}

func (s *PhotoService) GetPhotoByID(photoID uint) (models.Photo, error) {
	return s.photoRepo.GetPhotoByID(photoID)
}

func (s *PhotoService) GetPhotos() ([]models.Photo, error) {
	return s.photoRepo.GetPhotos()
}

func (s *PhotoService) UpdatePhotoByID(photo *models.Photo, UserID uint) (int64, error) {
	return s.photoRepo.UpdatePhotoByID(photo, UserID)
}

func (s *PhotoService) DeletePhotoByID(photoID uint, userID uint) (int64, error) {
	return s.photoRepo.DeletePhotoByID(photoID, userID)
}
