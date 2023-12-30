package repository

import (
	"github.com/jacobslunga/bubblegum-server/internal/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (repo *UserRepo) FindAll() ([]model.User, error) {
	var users []model.User

	result := repo.DB.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (repo *UserRepo) FindById(id string) (*model.User, error) {
	var user model.User

	result := repo.DB.First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *UserRepo) FindByEmail(email string) (*model.User, error) {
	var user model.User

	result := repo.DB.First(&user, "email = ?", email)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *UserRepo) FindByUsername(username string) (*model.User, error) {
	var user model.User

	result := repo.DB.First(&user, "username = ?", username)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *UserRepo) Update(user *model.User) error {
	result := repo.DB.Save(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *UserRepo) Delete(user *model.User) error {
	result := repo.DB.Delete(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *UserRepo) Login(email string, password string) (*model.User, error) {
	var user model.User

	result := repo.DB.First(&user, "email = ?", email)

	if result.Error != nil {
		return nil, result.Error
	}

	if !user.CheckPassword(password) {
		return nil, nil
	}

	return &user, nil
}

func (repo *UserRepo) Register(email, username, password string) (*model.User, error) {
	user := model.User{
		Email:    email,
		Username: username,
		Password: password,
	}

	err := user.HashPassword()
	if err != nil {
		return nil, err
	}

	result := repo.DB.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
