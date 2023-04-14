package services

import (
	"MygarmProject/models"
	"MygarmProject/repositories"
)

type CommentService struct {
	commentRepo *repositories.CommentRepository
}

func NewCommentService(commentRepo *repositories.CommentRepository) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
	}
}

func (s *CommentService) CreateComment(comment *models.Comment) error {
	return s.commentRepo.CreateComment(comment)
}

func (s *CommentService) GetCommentByID(commentID uint) (models.Comment, error) {
	return s.commentRepo.GetCommentByID(commentID)
}

func (s *CommentService) GetAllComments() ([]models.Comment, error) {
	return s.commentRepo.GetAllComments()
}

func (s *CommentService) UpdateCommentByID(comment *models.Comment, userID uint) (int64, error) {
	return s.commentRepo.UpdateCommentByID(comment, userID)
}

func (s *CommentService) DeleteCommentByID(commentID uint, userID uint) (int64, error) {
	return s.commentRepo.DeleteCommentByID(commentID, userID)
}
