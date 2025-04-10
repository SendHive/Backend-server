package services

import (
	"backend-server/dal"
	"backend-server/models"
	"log"
	"mime/multipart"

	minioDb "github.com/SendHive/Infra-Common/minio"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type IFileService interface {
	CreateFileEntry(req *models.CreateFileRequest, file *multipart.FileHeader, userId uuid.UUID) (*models.CreateFileEntryResponse, error)
	ListFiles(userId uuid.UUID) (response []*models.ListFilesResponse, err error)
}

type FileService struct {
	FileRepo    dal.IFile
	UserRepo    dal.IUser
	MinioClient *minio.Client
	IMinio      minioDb.IMinioService
}

func NewFilServiceRequest(mc *minio.Client, Im minioDb.IMinioService) (IFileService, error) {
	ser := &FileService{}
	err := ser.SetupRepo()
	if err != nil {
		return nil, err
	}
	ser.MinioClient = mc
	ser.IMinio = Im
	return ser, nil
}

func (f *FileService) SetupRepo() error {
	repo, err := dal.NewFileDalRequest()
	if err != nil {
		return err
	}
	f.FileRepo = repo

	urepo, err := dal.NewUserDalRequest()
	if err != nil {
		return err
	}
	f.UserRepo = urepo

	return nil
}

func (f *FileService) CreateFileEntry(req *models.CreateFileRequest, file *multipart.FileHeader, userId uuid.UUID) (*models.CreateFileEntryResponse, error) {
	if req.Name == "" {
		log.Println("No name countinue with the random name and saving in the minio")
	}

	userDetails, err := f.UserRepo.FindBy(userId)
	if err != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: "error while fetching the userdetails: " + err.Error(),
		}
	}

	log.Println(userDetails.Name)

	objectName, err := models.ReadCSV(file, "", userDetails.Name, f.MinioClient, f.IMinio)
	if err != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: "error while reading the csv: " + err.Error(),
		}
	}

	log.Println("the objectName for the file in the minio: ", objectName)

	ferr := f.FileRepo.Create(&models.DbFileDetails{
		Name:   objectName,
		UserId: userDetails.UserId,
	})

	if ferr != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: "error while creating  the filedetails: " + ferr.Error(),
		}
	}

	return &models.CreateFileEntryResponse{
		Message: "Created file with name: " + objectName,
	}, nil
}

func (f *FileService) ListFiles(userId uuid.UUID) (response []*models.ListFilesResponse, err error) {
	filesDetails, err := f.FileRepo.FindAll(userId)
	if err != nil {
		return nil, &models.ServiceResponse{
			Code:    500,
			Message: "error while listing files for the users: " + err.Error(),
		}
	}
	if len(filesDetails) == 0 {
		return make([]*models.ListFilesResponse, 0), nil
	}
	for _, i := range filesDetails {
		resp := &models.ListFilesResponse{}
		resp.Id = i.Id.String()
		resp.Name = i.Name
		response = append(response, resp)
	}
	return response, nil
}
