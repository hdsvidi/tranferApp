package repository

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"mncTestApp/model"
	"mncTestApp/utils"
)

type userCredentialRepoImpl struct {
	userCredDb *sqlx.DB
}

func (u *userCredentialRepoImpl) GetByUserNameAndPassword(user model.UserCredential) error {
	var isUserExist int
	err := u.userCredDb.Get(&isUserExist, "select count(id) from data_user where user_name=$1 and user_password=$2", user.GetUserName(), user.GetUserPassword())
	if err != nil {
		return errors.New(err.Error())
	}
	if isUserExist == 0 {
		return utils.DataNotFoundError()
	}
	return nil
}

func NewUserCredentialRepo(db *sqlx.DB) UserCredentialRepo {
	return &userCredentialRepoImpl{
		db,
	}
}
