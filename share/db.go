package share

import (
	"apisix-backend/config"
	"apisix-backend/structs"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	AppConfig := config.AppConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		AppConfig.DBUser,
		AppConfig.DBPassword,
		AppConfig.DBHost,
		AppConfig.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	DB = db
}

func MigrateDB() {
	models := []interface{}{
		&structs.ProductsEntity{},
		&structs.MyUsersEntity{},
	}

	if err := DB.AutoMigrate(models...); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migrated successfully")

	createInitialUser()
}

func createInitialUser() {
	const username = "myuser"
	const password = "mypassword"

	var existing structs.MyUsersEntity
	err := DB.Where("username = ?", username).First(&existing).Error
	if err == nil {
		log.Println("User already exists, skipping creation.")
		return
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		log.Fatal("Error checking existing user:", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Error hashing password:", err)
	}

	newUser := structs.MyUsersEntity{
		Username: username,
		Password: string(hashedPassword),
	}

	if err := DB.Create(&newUser).Error; err != nil {
		log.Fatal("Failed to create user:", err)
	}

	log.Println("Initial user created.")
}
