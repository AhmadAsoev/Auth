package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID `json:"id" gorm:"column:id"`
	Firstname string    `json:"first_name" gorm:"column:first_name"`
	LastName  string    `json:"last_name" gorm:"column:last_name"`
	Sex       Sex       `json:"sex" gorm:"column:sex"`
	BirthDate *string   `json:"birth_date" gorm:"column:birth_date"`
	Role      Role      `json:"role" gorm:"column:role"`
	Status    *Status   `json:"status" gorm:"column:status; default:active"`
	CreatedAt *string   `json:"created_at" gorm:"column:created_at; default:now()"`
	UpdatedAt *string   `json:"updated_at,omitempty" gorm:"column:updated_at"`
	DeletedAT *string   `json:"deleted_at,omitempty" gorm:"column:deleted_at"`
}

func (User) TableName() string {
	return "users"
}

func (u User) Validate() error {
	if len(u.Firstname) > 32 {
		return errors.New("firstName must not be more then 32")
	}
	if len(u.LastName) > 32 {
		return errors.New(" lastName must not be more then 32 ")
	}
	if err := u.Sex.Validate(); err != nil {
		return err
	}

	if err := u.Role.Validate(); err != nil {
		return err
	}

	if u.Status != nil {
		if err := u.Status.Validate(); err != nil {
			return err
		}
	}

	if u.BirthDate != nil {
		if _, err := time.Parse("02.01.2006", *u.BirthDate); err != nil {
			return errors.New("could not parse from string to time")
		}
	}

	if u.CreatedAt != nil {
		if _, err := time.Parse("02.01.2006T15:04:05", *u.CreatedAt); err != nil {
			return errors.New("cant parse createdAt")
		}
	}

	if u.UpdatedAt != nil {
		if _, err := time.Parse("02.01.2006T15:04:05", *u.UpdatedAt); err != nil {
			return errors.New("can't parse updatedAt")
		}
	}

	if u.DeletedAT != nil {
		if _, err := time.Parse("02.01.2006T15:04:05", *u.DeletedAT); err != nil {
			return errors.New("can't parse deletedAt")
		}
	}

	return nil
}

const (
	Male     = "male"
	Female   = "female"
	Admin    = "admin"
	Customer = "customer"
	Support  = "support"
	Active   = "active"
	Block    = "block"
	Delete   = "delete"
)

type Status string

func (s Status) Validate() error {
	switch s {
	case Active, Block, Delete:
		return nil
	default:
		return errors.New("must be active, block, delete")
	}
}

type Role string

func (r Role) Validate() error {
	switch r {
	case Admin, Customer, Support:
		return nil
	default:
		return errors.New("must be admin, customer, support")
	}
}

type Sex string

func (s Sex) Validate() error {
	switch s {
	case Male, Female:
		return nil
	default:
		return errors.New("must be male or female")
	}
}
