package models

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"regexp"
	"strings"
	"time"

	minioDb "github.com/SendHive/Infra-Common/minio"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type ServiceResponse struct {
	Code    int
	Message string
	Data    interface{} 
}

func (e *ServiceResponse) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Data: %+v", e.Code, e.Message, e.Data)
}

type TaskBody struct {
	TaskId uuid.UUID
	Name   string
}

const (
	STATUS_PENDING     = "PENDING"
	STATUS_IN_PROGRESS = "INPROGRESS"
	STATUS_COMPLETE    = "COMPLETED"
	CHARSET            = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func ReadCSV(file *multipart.FileHeader, filterDomain, username string, mc *minio.Client, Im minioDb.IMinioService) ([]string, error) {
	f, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	var emails []string

	for i, row := range records {
		if i == 0 {
			continue
		}
		if len(row) > 0 {
			email := strings.TrimSpace(row[0])
			if emailRegex.MatchString(email) && (filterDomain == "" || strings.HasSuffix(email, "@"+filterDomain)) {
				emails = append(emails, email)
			}
		}
	}

	if len(emails) == 0 {
		return nil, fmt.Errorf("no valid emails found in the file")
	}

	bucketName := username
	err = Im.CreateBucket(mc, bucketName)
	if err != nil {
		return nil, fmt.Errorf("failed to check bucket existence: %w", err)
	}

	objectName := fmt.Sprintf("%s-%s.csv", username, RandomString(4))

	err = Im.PutObject(mc, bucketName, file.Filename, objectName)
	if err != nil {
		return nil, fmt.Errorf("failed to upload file to MinIO: %w", err)
	}

	log.Printf("File uploaded to MinIO bucket: %s, object: %s", bucketName, objectName)

	return emails, nil
}

func RandomString(leng int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, leng)

	for i := range b {
		b[i] = CHARSET[r.Intn(len(CHARSET))]
	}

	return string(b)
}
