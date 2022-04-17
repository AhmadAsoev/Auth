package repository

import (
	"authDB/pkg/domain"
	"github.com/google/uuid"
)

func GetUserByID(id uuid.UUID) (user domain.User, err error) {
	return user, DB.First(&user, id).Error
}
