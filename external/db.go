package external

import (
	"backend-server/models"
	"log"

	infraDb "github.com/SendHive/Infra-Common/db"
	"gorm.io/gorm"
)

func ConnectDB() error {
	IdbConn, err := infraDb.NewDbRequest()
	if err != nil {
		log.Println("error while connecting to the database: ", err)
		return err
	}

	dbConn, err := IdbConn.InitDB()
	if err != nil {
		log.Println("error while getting the database instance: ", err)
		return err
	}

	//Migrating the tables if not done
	dbConn.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	err = dbConn.AutoMigrate(&models.DBSMTPDetails{}, &models.DBJobDetails{})
	if err != nil {
		log.Println("error while migrating the database instance: ", err)
		return err
	}
	log.Println("Database migration successfull....")
	return nil
}

func GetDbConn() (*gorm.DB, error) {
	IdbConn, err := infraDb.NewDbRequest()
	if err != nil {
		log.Println("error while connecting to the database: ", err)
		return nil, err
	}

	dbConn, err := IdbConn.InitDB()
	if err != nil {
		log.Println("error while getting the database instance: ", err)
		return nil, err
	}
	return dbConn, nil
}
