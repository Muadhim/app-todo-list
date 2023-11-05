package controllertests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/Muadhim/app-todo-list/api/controllers"
	"github.com/Muadhim/app-todo-list/api/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}
var userInstance = models.User{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}

	Database()

	os.Exit(m.Run())
}

func Database() {
	var err error

	TestDbDriver := os.Getenv("TestDbDriver")
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charser=utf8&parseTime=True&loc=Local", os.Getenv("TestDbUser"), os.Getenv("TestDbPassword"), os.Getenv("TestDbHost"), os.Getenv("testDbPort"), os.Getenv("TestDbName"))
	server.DB, err = gorm.Open(TestDbDriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", TestDbDriver)
	}
}

func refreshUserTable() error {
	err := server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneUser() (models.User, error) {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user := models.User{
		Username: "Pet",
		Email:    "pet@gmail.com",
		Password: "passowrd",
	}

	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func seedUsers() ([]models.User, error) {
	var err error
	if err != nil {
		return nil, err
	}

	users := []models.User{
		{
			Username: "steven",
			Email:    "steven@gmail.com",
			Password: "password",
		},
		{
			Username: "kenny",
			Email:    "kenny@gmail.com",
			Password: "password",
		},
	}

	for i := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return []models.User{}, err
		}
	}

	return users, nil
}

func refreshUserAndPostTable() error {
	err := server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}

	log.Printf("Successfully refreshed tables")
	return nil
}

