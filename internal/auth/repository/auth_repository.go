package repository

import (
	"log"

	"github.com/sohibjon7731/ecommerce_backend/config"
	"github.com/sohibjon7731/ecommerce_backend/internal/auth/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository() *AuthRepository {
	db, err := gorm.Open(postgres.Open(config.GetDBDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database" + err.Error())
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("failed to migrate user model: " + err.Error())
	}

	return &AuthRepository{DB: db}

}

func (r *AuthRepository) CreateUser(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *AuthRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("email= ?", email).First(&user).Error
	return &user, err
}

func (r *AuthRepository) ExistUserEmail(email string) (bool, error) {

	var count int64
	err := r.DB.Model(&model.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil

}