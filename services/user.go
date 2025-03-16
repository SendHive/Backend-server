package services

import (
	"backend-server/dal"
	"backend-server/models"
	"backend-server/secrets"
	"log"

	"github.com/google/uuid"
)

type IUser interface {
	SetupRepo() error
	CreateUserEntry(req *models.CreateUserRequest) (*models.CreateUserResponse, error)
	GetUserQRCodeImage(userId uuid.UUID) (string, error)
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
	secretKey, QrCodeURL := secrets.GenerateSecret(req.Email)
	log.Println("The Orcode: ", QrCodeURL)

	//Generate the Hash Password
	password, err := secrets.GenerateHash(req.Password)
	if err != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: "error while generating the hash: " + err.Error(),
		}
	}
	userId := uuid.New()
	err = u.UserRepo.Create(&models.DBUserDetails{
		UserId:    userId,
		SecretKey: secretKey,
		Name:      req.Name,
		Email:     req.Email,
		TotsUrl:   QrCodeURL,
		Password:  password,
	})
	if err != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: "error while creating the user: " + err.Error(),
		}
	}

	//Create the enrty in the secret database
	err = u.SecretRepo.Create(&models.DBSecretsDetails{
		UserId:    userId,
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

func (u *User) GetUserQRCodeImage(userId uuid.UUID) (string, error) {
	userDetails, err := u.UserRepo.FindBy(userId)
	if err != nil {
		return "", &models.ServiceResponse{
			Code:    500,
			Message: "error while creating the user entry: " + err.Error(),
		}
	}
	qrcode, err := models.GenerateQRCode(userDetails.TotsUrl)
	if err != nil {
		return "", &models.ServiceResponse{
			Code:    500,
			Message: "error while genrating the qr code :" + err.Error(),
		}
	}
	return string(qrcode), nil
}
