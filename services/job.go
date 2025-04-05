package services

import (
	"backend-server/dal"
	"backend-server/external"
	"backend-server/models"
	"encoding/json"
	"log"
	"mime/multipart"
	"time"

	minioDb "github.com/SendHive/Infra-Common/minio"
	"github.com/SendHive/Infra-Common/queue"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/rabbitmq/amqp091-go"
)

type IJobService interface {
	SetupRepo() error
	CreateJobEntry(req *models.CreateJobRequest, userId uuid.UUID, file *multipart.FileHeader) (*models.CreateJobResponse, error)
	UploadFiletoQueue(uuid.UUID, string) error
	ListJobEntry(userId uuid.UUID) ([]*models.ListJobEntryResponse, error)
}

type JobService struct {
	JobRepo     dal.IJob
	UserRepo    dal.IUser
	Queue       amqp091.Queue
	QConnn      *amqp091.Connection
	MinioClient *minio.Client
	IMinio      minioDb.IMinioService
}

func NewJobServiceRequest(queue amqp091.Queue, qConn *amqp091.Connection, mc *minio.Client, Im minioDb.IMinioService) (IJobService, error) {
	ser := &JobService{}

	err := ser.SetupRepo()
	if err != nil {
		return nil, err
	}
	ser.QConnn = qConn
	ser.Queue = queue
	ser.MinioClient = mc
	ser.IMinio = Im
	return ser, nil
}

func (job *JobService) SetupRepo() error {
	var err error
	repo, err := dal.NewJobDalRequest()
	if err != nil {
		return err
	}
	job.JobRepo = repo

	urepo, err := dal.NewUserDalRequest()
	if err != nil {
		return err
	}
	job.UserRepo = urepo

	return nil
}

func (job *JobService) CreateJobEntry(req *models.CreateJobRequest, userId uuid.UUID, file *multipart.FileHeader) (*models.CreateJobResponse, error) {

	jobDetails, err := job.JobRepo.FindBy(&models.DBJobDetails{
		Name: req.Name,
	})
	if err != nil {
		if err.Error() != "record not found" {
			return nil, &models.ServiceResponse{
				Code:    500,
				Message: "error while adding the job entry " + err.Error(),
			}

		}
	}

	if jobDetails.Name == req.Name {
		return nil, &models.ServiceResponse{
			Code:    404,
			Message: "The task with this name already exists",
		}
	}

	if file.Size == 0 {
		return nil, &models.ServiceResponse{
			Code:    404,
			Message: "Please check file, the contents of the file are empty.",
		}
	}

	
	userDetails, err := job.UserRepo.FindBy(userId)
	if err != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: "error while fetching the userdetails: " + err.Error(),
		}
	}

	objectName, err := models.ReadCSV(file, "", userDetails.Name, job.MinioClient, job.IMinio)
	if err != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: "error while reading the csv: " + err.Error(),
		}
	}

	taskId := uuid.New()
	jerr := job.JobRepo.Create(&models.DBJobDetails{
		Name:       req.Name,
		UserId:     userId,
		TaskId:     taskId,
		CreatedAt:  time.Now(),
		Status:     models.STATUS_PENDING,
		ObjectName: objectName,
	})

	if jerr != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: "error while creating the job: " + jerr.Error(),
		}
	}

	errChan := make(chan error, 1)
	go func() {
		defer close(errChan)
		qerr := job.UploadFiletoQueue(taskId, req.Name)
		if qerr != nil {
			errChan <- qerr
		}
	}()
	select {
	case err := <-errChan:
		if err != nil {
			return nil, &models.ServiceResponse{
				Code:    500,
				Message: "error while publishing the message to queue: " + err.Error(),
			}
		}
	default:
	}

	return &models.CreateJobResponse{
		Message: "Job has started successfully",
	}, nil
}

func (job *JobService) UploadFiletoQueue(taskId uuid.UUID, name string) error {
	Iq, err := queue.NewQueueRequest()
	if err != nil {
		log.Println("the error while creating the queue instance: ", err)
		return &models.ServiceResponse{
			Code:    500,
			Message: "error while creating the queue request: " + err.Error(),
		}
	}

	task := models.TaskBody{
		TaskId: taskId,
		Name:   name,
	}

	body, err := json.Marshal(task)

	if err != nil {
		return &models.ServiceResponse{
			Code:    500,
			Message: "error while converting the task info in the string: " + err.Error(),
		}
	}

	Qerr := external.PublishMessage(job.Queue, Iq, job.QConnn, string(body))
	if Qerr != nil {
		return &models.ServiceResponse{
			Code:    500,
			Message: Qerr.Error(),
		}
	}
	return nil
}

func (job *JobService) ListJobEntry(userId uuid.UUID) ([]*models.ListJobEntryResponse, error) {
	userDetails, err := job.UserRepo.FindBy(userId)
	if err != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: "error while finding the user: " + err.Error(),
		}
	}
	jobDetails, err := job.JobRepo.FindAll(userDetails.UserId)
	if err != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: "error while finding the job: " + err.Error(),
		}
	}
	var resp []*models.ListJobEntryResponse
	for _, i := range jobDetails {
		temp := &models.ListJobEntryResponse{}
		temp.Name = i.Name
		temp.Status = i.Status
		temp.Type = i.Type
		resp = append(resp, temp)
	}
	log.Println("the list: ", resp)
	return resp, nil
}
