package repository

import (
	"authDB/pkg/domain"
	"errors"
	"github.com/google/uuid"
	"time"
)

func UpdateUser(user domain.User, id uuid.UUID) (err error) {
	result := DB.First(&domain.User{}, id)
	if result.Error != nil {
		return errors.New("user by id did not exist")
	}
	return DB.Where(&domain.User{Id: id}).Updates(domain.User{
		Firstname: user.Firstname,
		LastName:  user.LastName,
		Sex:       user.Sex,
		Role:      user.Role,
		BirthDate: user.BirthDate,
	}).Update("updated_at", time.Now()).Update("status", domain.Active).Error

}
