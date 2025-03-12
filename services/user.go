package services

import (
	"backend-server/dal"
	"backend-server/models"
	"backend-server/secrets"

	"github.com/google/uuid"
)

type IUser interface {
	SetupRepo() error
	CreateUserEntry(req *models.CreateUserRequest) (*models.CreateUserResponse, error)
}
type User struct {
	UserRepo   dal.IUser
	SecretRepo dal.ISecret
}

func NewUserServiceReqest() (IUser, error) {
	service := &User{}
	err := service.SetupRepo()
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (u *User) SetupRepo() error {
	var err error

	resp, err := dal.NewUserDalRequest()
	if err != nil {
		return err
	}
	u.UserRepo = resp

	srepo, err := dal.NewDalSecretRequest()
	if err != nil {
		return err
	}
	u.SecretRepo = srepo
	return err
}

func (u *User) CreateUserEntry(req *models.CreateUserRequest) (*models.CreateUserResponse, error) {
	resp, err := u.UserRepo.FindByConditions(&models.DBUserDetails{
		Name: req.Name,
	})
	if err != nil {
		if err.Error() != "record not found" {
			return nil, &models.ServiceResponse{
				Code:    500,
				Message: "Error while fetching the user details: " + err.Error(),
			}
		}
	}
	if resp.Name == req.Name {
		return nil, &models.ServiceResponse{
			Code:    400,
			Message: "User already exists proceed to login",
		}
	}
	secretKey := secrets.GenerateSecret()
	userId := uuid.New()
	err = u.UserRepo.Create(&models.DBUserDetails{
		UserId:    userId,
		SecretKey: secretKey,
		Name:      req.Name,
	})
	if err != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: "error while creating the user: " + err.Error(),
		}
	}

	//Create the enrty in the secret database
	err = u.SecretRepo.Create(&models.DBSecretsDetails{
		UserId: userId,
		SecretKey: secretKey,
	})

	if err != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: "error while creating the secret of the user: " + err.Error(),
		}
	}

	return &models.CreateUserResponse{
		Message: "User with the name " + req.Name + " created successfully.",
	}, nil
}

