package repositories

import (
	"MygarmProject/models"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db}
}

func (repo *CommentRepository) CreateComment(comment *models.Comment) error {
	return repo.db.Debug().Create(comment).Error
}

func (repo *CommentRepository) GetCommentByID(commentID uint) (models.Comment, error) {
	var result models.Comment
	err := repo.db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username")
	}).Preload("Photo", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, title, caption, photo_url")
	}).Find(&result, "id = ?", commentID).Error
	return result, err
}

func (repo *CommentRepository) GetAllComments() ([]models.Comment, error) {
	var results []models.Comment
	err := repo.db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username")
	}).Preload("Photo", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, title, caption, photo_url")
	}).Find(&results).Error
	return results, err
}

func (repo *CommentRepository) DeleteCommentByID(commentID uint, userID uint) (int64, error) {
	result := repo.db.Debug().Delete(&models.Comment{}, commentID).Where("user_id = ?", userID)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (repo *CommentRepository) UpdateCommentByID(comment *models.Comment, userID uint) (int64, error) {
	result := repo.db.Model(comment).Where("id = ?", comment.ID).Updates(models.Comment{Message: comment.Message}).Where("user_id = ?", userID)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
