package services

import (
	"authDB/pkg/application/repository"
	"authDB/pkg/domain"
	"github.com/google/uuid"
	"net/http"
)

func AddUser(user domain.User) (response domain.Response) {
	if err := user.Validate(); err != nil {
		return domain.Response{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		}
	}

	id := uuid.New()
	user.Id = id

	if err := repository.AddUser(user); err != nil {
		return domain.Response{
			Code:  http.StatusInternalServerError,
			Error: "Can't create user into db",
		}
	}

	return domain.Response{
		Code:    http.StatusOK,
		Message: "User created",
	}
}
