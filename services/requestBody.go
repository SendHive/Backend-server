package services

import (
	"backend-server/dal"
	"backend-server/models"
	"log"

	"github.com/google/uuid"
)

type RequestBody struct {
	ReqRepo  dal.IRequestBody
	UserRepo dal.IUser
}

type IRequestBody interface {
	CreateRequestEntry(req *models.CreateRequestBodyRequest, userId uuid.UUID) (resp *models.CreateFileEntryResponse, err error)
	ListAllRequestEntry(userId uuid.UUID) (response []*models.ListRequestBodyResponse, err error)
	FindRequestEntry(reqId uuid.UUID, userId uuid.UUID) (response *models.ListRequestBodyResponse, err error)
	UpdateRequestEntry(reqId uuid.UUID, userId uuid.UUID, req *models.UpdateRequestBodyEntry) (response *models.UpdateRequestEntry, err error)
}

func NewRequestBodyServiceRequest() (IRequestBody, error) {
	ser := &RequestBody{}
	err := ser.SetupRepo()
	if err != nil {
		return nil, err
	}
	return ser, nil
}

func (r *RequestBody) SetupRepo() error {
	repo, err := dal.NewRequestBodyDalRequest()
	if err != nil {
		return err
	}
	r.ReqRepo = repo
	urepo, err := dal.NewUserDalRequest()
	if err != nil {
		return err
	}
	r.UserRepo = urepo

	return nil
}

func (r *RequestBody) CheckUser(userId uuid.UUID) error {
	resp, err := r.UserRepo.FindBy(userId)
	if err != nil {
		if err.Error() != "record not found" {
			return &models.ServiceResponse{
				Code:    500,
				Message: "error while finding the user while creating the request " + err.Error(),
			}

		} else {
			return &models.ServiceResponse{
				Code:    404,
				Message: "error while creating a request: " + err.Error(),
			}
		}
	}
	if resp.UserId != userId {
		return &models.ServiceResponse{
			Code:    404,
			Message: "User not found",
		}
	}
	return nil
}

func (r *RequestBody) CreateRequestEntry(req *models.CreateRequestBodyRequest, userId uuid.UUID) (*models.CreateFileEntryResponse, error) {
	uErr := r.CheckUser(userId)
	if uErr != nil {
		return nil, uErr
	}
	rErr := r.ReqRepo.Create(&models.DbRequestBody{
		Name:        req.Name,
		RequestBody: req.Promo_Text,
		UserId:      userId,
	})
	if rErr != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: "error while registering the data in the database: " + rErr.Error(),
		}
	}
	return &models.CreateFileEntryResponse{
		Message: "added the requestBody successfully",
	}, nil
}

func (r *RequestBody) ListAllRequestEntry(userId uuid.UUID) (response []*models.ListRequestBodyResponse, err error) {
	uErr := r.CheckUser(userId)
	if uErr != nil {
		return nil, uErr
	}
	resp, err := r.ReqRepo.FindAll(userId)
	if err != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: "error while finding the request for the given user:" + err.Error(),
		}
	}
	for _, i := range resp {
		temp := &models.ListRequestBodyResponse{}
		temp.Name = i.Name
		temp.Promo_Text = i.RequestBody
		response = append(response, temp)
	}
	return response, nil
}

func (r *RequestBody) FindRequestEntry(reqId uuid.UUID, userId uuid.UUID) (response *models.ListRequestBodyResponse, err error) {
	uErr := r.CheckUser(userId)
	if uErr != nil {
		return nil, uErr
	}
	resp, err := r.ReqRepo.FindBy(&models.DbRequestBody{
		Id: reqId,
	})
	if err != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: "error while getting the requestBody info: " + err.Error(),
		}
	}
	return &models.ListRequestBodyResponse{
		Name:       resp.Name,
		Promo_Text: resp.RequestBody,
	}, nil
}

func (r *RequestBody) UpdateRequestEntry(reqId uuid.UUID, userId uuid.UUID, req *models.UpdateRequestBodyEntry) (response *models.UpdateRequestEntry, err error) {
	uErr := r.CheckUser(userId)
	if uErr != nil {
		return nil, uErr
	}
	resp, err := r.ReqRepo.FindBy(&models.DbRequestBody{
		Id: reqId,
	})
	if err != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: "error while fetching the promo text details: " + err.Error(),
		}
	}
	log.Println("the promo details: ", resp.RequestBody)
	log.Println("the requestBody: ", req)
	Uresponse, err := r.ReqRepo.Update(&models.DbRequestBody{
		Id:          resp.Id,
		Name:        req.Name,
		UserId:      userId,
		RequestBody: req.Promo_Text,
	})
	if err != nil {
		return nil, &models.ServiceResponse{
			Code: 500,
			Message: "error while updating the promo text : "+ err.Error(),
		}
	}
	return &models.UpdateRequestEntry{
		Name:       Uresponse.Name,
		Promo_Text: Uresponse.RequestBody,
	}, nil
}
