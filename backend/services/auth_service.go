package services

import (
    "errors"
    "project/models"
    "project/repository"
    "project/utils"
)

type AuthService struct {
    UserRepo *repository.UserRepo
}

func (s *AuthService) Authenticate(login, password string) (*models.User, error) {
    user, err := s.UserRepo.GetByLogin(login)
    if err != nil {
        return nil, errors.New("user not found")
    }
    if !utils.CheckPasswordHash(password, user.Password) {
        return nil, errors.New("invalid password")
    }
    return user, nil
}
