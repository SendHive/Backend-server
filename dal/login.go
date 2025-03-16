package dal

import (
	"backend-server/external"
	"backend-server/models"
)

type ILogin interface {
	Create(value *models.DbLoginDetails) error
}

type Login struct{}

func NewLoginDalRequest() (ILogin, error) {
	return &Login{}, nil
}

func (l *Login) Create(value *models.DbLoginDetails) error {
	dbConn, err := external.GetDbConn()
	if err != nil {
		return err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return transaction.Error
	}
	defer transaction.Rollback()
	customerAddition := transaction.Create(&value)
	if customerAddition.Error != nil {
		return customerAddition.Error
	}
	transaction.Commit()
	return nil
}
