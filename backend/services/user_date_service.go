package services

import (
    "project/models"
    "project/repository"
)

type UserDateService struct {
    UserDateRepo *repository.UserDateRepo
}

func (s *UserDateService) GetUserDates(userID int64) ([]models.UserDate, error) {
    return s.UserDateRepo.GetByUserID(userID)
}

func (s *UserDateService) CreateUserDate(ud *models.UserDate) error {
    return s.UserDateRepo.Create(ud)
}

// Update и Delete аналогично
