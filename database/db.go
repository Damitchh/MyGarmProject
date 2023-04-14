package database

import (
	"MygarmProject/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	host     = os.Getenv("PGHOST")
	user     = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	dbport   = 6469
	dbname   = os.Getenv("PGDATABASE")

	db  *gorm.DB
	err error
)

func StartDB() {
	fmt.Println("connecting to database....")
	config := fmt.Sprintf("host= %s user= %s password= %s port= %d dbname= %s sslmode= disable",
		host, user, password, dbport, dbname)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	fmt.Println("connected to database")

	// Log the database object to check if it's nil
	fmt.Printf("DB object: %v\n", db)

	// Auto-migrate the models
	err = db.Debug().AutoMigrate(models.User{}, models.Comment{}, models.Photo{}, models.SocialMedia{})
	if err != nil {
		log.Fatal("error migrating to database: ", err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func PingDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	if sqlDB != nil {
		err := sqlDB.Close()
		if err != nil {
			log.Fatalf("error closing database connection: %v", err)
		}
	}
}
