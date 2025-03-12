package dal

import (
	"backend-server/external"
	"backend-server/models"
)

type Secret struct{}

type ISecret interface {
	Create(value *models.DBSecretsDetails) error
}

func NewDalSecretRequest() (ISecret, error) {
	return &Secret{}, nil
}

func (s *Secret) Create(value *models.DBSecretsDetails) error {
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
