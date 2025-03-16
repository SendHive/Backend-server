package services

import (
	"backend-server/dal"
	"backend-server/models"
	"backend-server/secrets"
)

type ILoginService interface {
	SetupRepo() error
	CreateLoginEntry(req *models.CreateLoginRequest) (resp *models.CreateLoginResponse, err error)
}

type LoginService struct {
	LoginRepo dal.ILogin
	UserRepo  dal.IUser
}

func NewLoginService() (ILoginService, error) {
	ser := &LoginService{}
	err := ser.SetupRepo()
	if err != nil {
		return nil, err
	}
	return ser, nil
}

func (l *LoginService) SetupRepo() error {
	var err error
	login, err := dal.NewLoginDalRequest()
	if err != nil {
		return err
	}
	user, err := dal.NewUserDalRequest()
	if err != nil {
		return err
	}
	l.LoginRepo = login
	l.UserRepo = user
	return nil
}

func (l *LoginService) CreateLoginEntry(req *models.CreateLoginRequest) (resp *models.CreateLoginResponse, err error) {
	userDetails, err := l.UserRepo.FindByConditions(&models.DBUserDetails{
		Email: req.Email,
	})
	if err != nil {
		if err.Error() == "record not found" {
			return nil, &models.ServiceResponse{
				Code:    404,
				Message: "user with name doesnot exist",
			}
		} else {
			return nil, &models.ServiceResponse{
				Code:    500,
				Message: "error while finding the userdetails: " + err.Error(),
			}
		}
	}

	//Camparing the password
	check, err := secrets.ComparePassword(req.Password, userDetails.Password)
	if err != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: "error while comparing the password: " + err.Error(),
		}
	}
	if !check {
		return nil, &models.ServiceResponse{
			Code:    404,
			Message: "Password doesn't match",
		}
	}

	err = l.LoginRepo.Create(&models.DbLoginDetails{
		UserId:  userDetails.UserId,
		IsLogin: true,
	})
	if err != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: "error while creating a entry in the login repo: " + err.Error(),
		}
	}
	return &models.CreateLoginResponse{
		Message: "User LoggedIn Successfully",
	}, nil
}
