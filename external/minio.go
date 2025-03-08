package external

import (
	"log"

	minioDb "github.com/SendHive/Infra-Common/minio"
	"github.com/minio/minio-go/v7"
)

func ConnectMinio() (*minio.Client, minioDb.IMinioService, error) {
	dbI, err := minioDb.NewMinioRequest()
	if err != nil {
		return nil, nil, err
	}
	conn, err := dbI.MinioConnect()
	if err != nil {
		return nil, nil, err
	}
	return conn, dbI, nil
}

func UploadFile(dbI minioDb.IMinioService, conn *minio.Client, file string, objectName string, bucketName string) error {
	err := dbI.CreateBucket(conn, bucketName)
	if err != nil {
		log.Println("the error while creating the bucket: ", err)
		return err
	}
	perr := dbI.PutObject(conn, bucketName, file, objectName)
	if perr != nil {
		return perr
	}
	return nil
}
