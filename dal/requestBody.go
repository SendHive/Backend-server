package dal

import (
	"backend-server/external"
	"backend-server/models"
	"fmt"

	"github.com/google/uuid"
)

type RequestBody struct{}

type IRequestBody interface {
	Create(value *models.DbRequestBody) error
	FindAll(userId uuid.UUID) (value []*models.DbRequestBody, err error)
	FindBy(conditions *models.DbRequestBody) (*models.DbRequestBody, error)
}

func NewRequestBodyDalRequest() (IRequestBody, error) {
	return &RequestBody{}, nil
}

func (r *RequestBody) Create(value *models.DbRequestBody) error {
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

func (r *RequestBody) FindAll(userId uuid.UUID) (value []*models.DbRequestBody, err error) {
	dbConn, err := external.GetDbConn()
	if err != nil {
		return nil, err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	defer transaction.Rollback()
	var response []*models.DbRequestBody
	resp := transaction.Where("user_id = ?", userId).Find(&response)
	if resp.Error != nil {
		return nil, resp.Error
	}
	if resp.RowsAffected == 0 {
		return nil, fmt.Errorf("no jobs found for user ID: %s", userId)
	}
	return response, nil
}

func (r *RequestBody) FindBy(conditions *models.DbRequestBody) (*models.DbRequestBody, error) {
	dbConn, err := external.GetDbConn()
	if err != nil {
		return nil, err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	defer transaction.Rollback()
	var resp *models.DbRequestBody
	ferr := transaction.Find(&resp, &conditions)
	if ferr.Error != nil {
		fmt.Println("the error while finding the job:", ferr.Error)
		return nil, ferr.Error
	}
	return resp, nil
}
