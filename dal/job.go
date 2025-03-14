package dal

import (
	"backend-server/external"
	"backend-server/models"
	"log"
)

type IJob interface {
	Create(value *models.DBJobDetails) error
	FindBy(conditions *models.DBJobDetails) (*models.DBJobDetails, error)
}

type Job struct{}

func NewJobDalRequest() (IJob, error) {
	return &Job{}, nil
}

func (j *Job) Create(value *models.DBJobDetails) error {
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

func (j *Job) FindBy(conditions *models.DBJobDetails) (*models.DBJobDetails, error) {
	dbConn, err := external.GetDbConn()
	if err != nil {
		return nil, err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	defer transaction.Rollback()
	var resp *models.DBJobDetails
	ferr := transaction.Find(&resp, &conditions)
	if ferr.Error != nil {
		log.Println("the error while finding the job:", ferr.Error)
		return nil, ferr.Error
	}
	return resp, nil
}
