package service

import (
	"go-pos-app/internal/model"
	"go-pos-app/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) GetUsers() ([]model.User, error){
	return s.Repo.GetUsers()
}

func (s *UserService) CreateUsers(user model.User) error {
	
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil{
		return err
	}

	user.Password = string(hashedPassword)

	return s.Repo.CreateUsers(user)
}

func (s *UserService) UpdateUser(user model.User) error{
	return s.Repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int)error {
	return s.Repo.DeleteUser(id)
}