package services

import (
	"authDB/pkg/application/repository"
	"authDB/pkg/domain"
	"github.com/google/uuid"
	"net/http"
)

func DeleteUser(urlId string) (response domain.Response) {
	if len(urlId) == 0 {
		return domain.Response{
			Code:  http.StatusBadRequest,
			Error: "Id must not be empty",
		}
	}
	id, err := uuid.Parse(urlId)
	if err != nil {
		return domain.Response{
			Code:  http.StatusBadRequest,
			Error: "Format id not uuid",
		}
	}

	user, err := repository.GetUserByID(id)
	if err != nil {
		return domain.Response{
			Code:  http.StatusInternalServerError,
			Error: "Can't get by ID",
		}
	}
	if *user.Status == domain.Delete {
		return domain.Response{
			Code:  http.StatusBadRequest,
			Error: "This account is already deleted",
		}
	}

	if err := repository.DeleteUser(id); err != nil {
		return domain.Response{
			Code:  http.StatusInternalServerError,
			Error: "Can't delete user by id",
		}
	}

	return domain.Response{
		Code:    http.StatusOK,
		Message: "successful deleted",
	}
}
