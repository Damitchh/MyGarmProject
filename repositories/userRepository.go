package repositories

import (
	"MygarmProject/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) LoginUser(user *models.User) error {
	return repo.db.Where("username = ?", user.Username).Take(&user).Error
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	return repo.db.Debug().Create(user).Error
}

func (repo *UserRepository) GetUserByID(userID uint) ([]models.User, error) {
	var results []models.User
	err := repo.db.Debug().Find(&results, "id = ?", userID).Error
	return results, err
}

func (repo *UserRepository) GetAllUsers() ([]models.User, error) {
	var results []models.User
	err := repo.db.Debug().Find(&results).Error
	return results, err
}

func (repo *UserRepository) DeleteUserByID(userID uint) (int64, error) {
	result := repo.db.Debug().Delete(&models.User{}, userID)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (repo *UserRepository) UpdateUserByID(user *models.User) (int64, error) {
	newUser := user
	result := repo.db.Model(models.User{}).Where("id = ?", user.ID).Updates(&newUser)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
