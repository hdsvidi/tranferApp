package repository

import "mncTestApp/model"

type UserCredentialRepo interface {
	GetByUserNameAndPassword(user model.UserCredential) error
}
