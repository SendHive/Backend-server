package dal

import (
	"backend-server/external"
	"backend-server/models"
	"log"

	"github.com/google/uuid"
)

type IUser interface {
	FindBy(userId uuid.UUID) (*models.DBUserDetails, error)
}

type User struct{}

func NewUserDalRequest() (IUser, error) {
	return &User{}, nil
}

func (u *User) FindBy(userId uuid.UUID) (*models.DBUserDetails, error) {
	dbConn, err := external.GetDbConn()
	if err != nil {
		return nil, err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	defer transaction.Rollback()
	resp := &models.DBUserDetails{}
	ferr := transaction.Find(&resp, &models.DBUserDetails{
		UserId: userId,
	})
	if ferr.Error != nil {
		log.Println("the error : ", ferr)
		return nil, ferr.Error
	}
	return resp, nil
}
