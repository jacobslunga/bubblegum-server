package service

import (
	"errors"

	"github.com/jacobslunga/bubblegum-server/internal/model"
	"github.com/jacobslunga/bubblegum-server/internal/repository"
)

var (
	ErrEmailUserExists    = errors.New("There is already someone with this email")
	ErrUsernameUserExists = errors.New("There is already someone with this username")
	ErrUserNotFound       = errors.New("User not found")
)

type UserService struct {
	repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) GetUserById(userId string) (*model.User, error) {
	return s.repo.FindById(userId)
}

func (s *UserService) LoginUser(email string, password string) (*model.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (s *UserService) CreateUser(email, username, password string) (*model.User, error) {
	emailUser, _ := s.repo.FindByEmail(email)
	usernameUser, _ := s.repo.FindByUsername(username)
	if emailUser != nil {
		return nil, ErrEmailUserExists
	}
	if usernameUser != nil {
		return nil, ErrUsernameUserExists
	}

	return s.repo.Register(email, username, password)
}

func (s *UserService) UpdateUser(user *model.User) error {
	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(user *model.User) error {
	return s.repo.Delete(user)
}
