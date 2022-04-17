package repository

import "authDB/pkg/domain"

func AddUser(user domain.User) error {
	return DB.Create(&user).Error

}
