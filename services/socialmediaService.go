package services

import (
	"MygarmProject/models"
	"MygarmProject/repositories"
)

type SocialMediaService struct {
	socialmediaRepo *repositories.SocialMediaRepository
}

func NewSocialMediaService(socialmediaRepo *repositories.SocialMediaRepository) *SocialMediaService {
	return &SocialMediaService{
		socialmediaRepo: socialmediaRepo,
	}
}

func (s *SocialMediaService) CreateSocialMedia(socialmedia *models.SocialMedia) error {
	return s.socialmediaRepo.CreateSocialMedia(socialmedia)
}

func (s *SocialMediaService) GetSocialMediaByID(socialmediaID uint) (models.SocialMedia, error) {
	return s.socialmediaRepo.GetSocialMediaByID(socialmediaID)
}

func (s *SocialMediaService) GetAllSocialMedias() ([]models.SocialMedia, error) {
	return s.socialmediaRepo.GetAllSocialMedias()
}

func (s *SocialMediaService) UpdateSocialMediaByID(socialmedia *models.SocialMedia, userID uint) (int64, error) {
	return s.socialmediaRepo.UpdateSocialMediaByID(socialmedia, userID)
}

func (s *SocialMediaService) DeleteSocialMediaByID(socialmediaID uint, userID uint) (int64, error) {
	return s.socialmediaRepo.DeleteSocialMediaByID(socialmediaID, userID)
}
