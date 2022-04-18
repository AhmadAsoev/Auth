package repository

import (
	"authDB/pkg/domain"
	"github.com/google/uuid"
	"time"
)

func DeleteUser(id uuid.UUID) error {
	return DB.Model(&domain.User{}).Where(&domain.User{Id: id}).
		Update("status", domain.Delete).
		Update("deleted_at", time.Now()).Error
}
