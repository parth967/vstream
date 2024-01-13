package db

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/vstream/internal/models"

	"errors"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connectToSQLLite(dbName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func PerformMigration(db *gorm.DB) {
	// defer func() {
	// 	dbInstance, _ := db.DB()
	// 	_ = dbInstance.Close()
	// }()

	if !db.Migrator().HasTable(&models.User{}) {
		db.AutoMigrate(&models.User{})
	}

}

func DBConnect(autoMigration bool) (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading env file")
	}

	dbName := os.Getenv("LITE_DB_NAME")

	db, err := connectToSQLLite(dbName)

	if autoMigration {
		PerformMigration(db)
	}

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func ValidateUser(colName string, colValue string, expectedVal string, user *models.User, c *fiber.Ctx) (bool, error) {
	isValid := true

	dbConn, err := DBConnect(true)
	if err != nil {
		return false, errors.New("connection failed")
	}

	result := dbConn.Where("Username= ?", colValue).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return false, errors.New("record not found")
	} else if result.Error != nil {
		return false, errors.New("user table not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(expectedVal)); err != nil {
		return false, errors.New("password wrong")
	}

	return isValid, nil
}

func AddUser(username, password, permisson string, ctx *fiber.Ctx) error {

	dbConn, err := DBConnect(true)
	if err != nil {
		return errors.New("connection failed")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("password can not hashed")
	}

	var maxID int
	if err := dbConn.Model(&models.User{}).Select("COALESCE(MAX(id), 0) + 1").Scan(&maxID).Error; err != nil {
		return errors.New("something is wrong")
	}

	now := time.Now()
	newUser := models.User{
		ID:        uint8(maxID),
		Username:  username,
		Password:  string(hashedPassword),
		Access:    permisson,
		CreatedAt: now,
	}

	if err := dbConn.Create(&newUser).Error; err != nil {
		return errors.New("internal server error")
	}
	return nil
}
