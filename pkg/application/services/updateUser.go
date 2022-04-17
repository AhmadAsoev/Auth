package services

import (
	"authDB/pkg/application/repository"
	"authDB/pkg/domain"
	"github.com/google/uuid"
	"net/http"
)

func UpdateUser(user domain.User, urlId string) (response domain.Response) {
	if urlId == "" {
		return domain.Response{
			Code:  http.StatusBadRequest,
			Error: "id must not be empty",
		}
	}
	id, err := uuid.Parse(urlId)
	if err != nil {
		return domain.Response{
			Code:  http.StatusBadRequest,
			Error: "Id must be uuid format",
		}
	}
	if err := user.Validate(); err != nil {
		return domain.Response{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		}
	}
	if err := repository.UpdateUser(user, id); err != nil {
		if err.Error() == "user by id did not exist" {
			return domain.Response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			}
		}
		return domain.Response{
			Code:  http.StatusInternalServerError,
			Error: "Can't update",
		}
	}
	return domain.Response{
		Code:    http.StatusOK,
		Message: "Update is successful",
	}
}
