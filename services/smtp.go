package services

import (
	"backend-server/dal"
	"backend-server/models"
	"log"
	"time"

	"github.com/google/uuid"
)

type SmtpService struct {
	SmtpRepo dal.ISmtpDal
}

type ISmtpService interface {
	SetupRepo() error
	CreateSmtpEntry(req *models.CreateSmtpEntryRequest) (*models.CreateSmtpEntryResponse, error)
	UpdateSmtpEntry(smtpId string, req *models.UpdateSmtpEntryRequest) (*models.UpdateSmtpEntryResponse, error)
	ListSmtpEntry(userId string) ([]*models.ListSmtpEntryResponse, error)
}

func NewSmtpServiceRequest() (ISmtpService, error) {
	service := &SmtpService{}
	err := service.SetupRepo()
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (s *SmtpService) SetupRepo() error {
	var err error

	resp, err := dal.NewSmtpDalRequest()
	if err != nil {
		return err
	}
	s.SmtpRepo = resp

	return err
}

func (s *SmtpService) CreateSmtpEntry(req *models.CreateSmtpEntryRequest) (*models.CreateSmtpEntryResponse, error) {

	resp, err := s.SmtpRepo.FindBy(&models.DBSMTPDetails{
		Server: req.Server,
	})

	if err != nil {
		if err.Error() != "record not found" {
			return nil, &models.ServiceResponse{
				Code:    500,
				Message: "error while adding the smtp entry " + err.Error(),
			}

		}
	}

	if resp.Server == req.Server {
		return nil, &models.ServiceResponse{
			Code:    404,
			Message: "Server already exist",
		}
	}

	err = s.SmtpRepo.Create(&models.DBSMTPDetails{
		UserId:    uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Server:    req.Server,
		Port:      req.Port,
		Username:  req.Username,
		Password:  req.Password,
	})

	if err != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: "error while adding the smtp entry " + err.Error(),
		}
	}

	return &models.CreateSmtpEntryResponse{
		Message: "User with Smtp server name " + req.Server + " added successfully.",
	}, nil
}

func (s *SmtpService) UpdateSmtpEntry(smtpId string, req *models.UpdateSmtpEntryRequest) (*models.UpdateSmtpEntryResponse, error) {

	resp, err := s.SmtpRepo.FindBy(&models.DBSMTPDetails{
		Id: uuid.MustParse(smtpId),
	})

	if err != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	if resp.Port == req.Port && resp.Server == req.Server && resp.Username == req.Username && resp.Password == req.Password {
		return nil, &models.ServiceResponse{
			Code:    409,
			Message: "Nothing to change",
		}
	}

	uerr := s.SmtpRepo.Update(resp.Id, &models.DBSMTPDetails{
		Server:   req.Server,
		Port:     req.Port,
		Password: req.Password,
		Username: req.Username,
	})

	if uerr != nil {
		return nil, &models.ServiceResponse{
			Message: "Update successfull",
		}
	}

	return &models.UpdateSmtpEntryResponse{}, nil
}

func (s *SmtpService) ListSmtpEntry(userId string) ([]*models.ListSmtpEntryResponse, error) {
	resp, err := s.SmtpRepo.GetAll(uuid.MustParse(userId))
	if err != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: "error while adding the smtp entry " + err.Error(),
		}
	}
	var smtpList []*models.ListSmtpEntryResponse
	for _, d := range resp {
		var smtp models.ListSmtpEntryResponse
		smtp.Server = d.Server
		smtp.Port = d.Port
		smtp.Username = d.Username

		smtpList = append(smtpList, &smtp)
	}
	log.Println("the resp: ", resp)
	return smtpList, nil
}
