package main

import (
	"Chinnasit/adapters"
	"Chinnasit/entities"
	"Chinnasit/usecases"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db := initDatabase()
	app := fiber.New()
	app.Use(cors.New())

	userRepo := adapters.NewGormUserRepository(db)
	userUsecase := usecases.NewUserService(userRepo)
	userHandler := adapters.NewHttpUserHandler(userUsecase)

	app.Post("/user", userHandler.CreateUser)
	app.Get("/users", userHandler.GetUsers)
	app.Put("/users/:id", userHandler.UpdateUserFull)
	app.Patch("/users/:id", userHandler.UpdateUserPartial)
	app.Delete("/users/:id", userHandler.DeleteUser)

	app.Listen(":3000")
}

// Trace SQL command
type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n----------------\n", sql)
}

// Initial the database
func initDatabase() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: &SqlLogger{}})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(entities.User{})
	var count int64
	db.Model(&entities.User{}).Count(&count)
	if count == 0 {
		CreateSampleUsers(db)
	}

	return db
}

func StringPtr(s string) *string {
    return &s
}

func UintPtr(i uint) *uint {
    return &i
}

func CreateSampleUsers(db *gorm.DB) {
	users := []entities.User{
		{FirstName: StringPtr("John"), LastName: StringPtr("Doe"), Email: StringPtr("john.doe@example.com"), Password: StringPtr("abcd12345"), RoleId: UintPtr(1), Active: true},
        {FirstName: StringPtr("Jane"), LastName: StringPtr("Doe"), Email: StringPtr("jane.doe@example.com"), Password: StringPtr("abcd12345"), RoleId: UintPtr(2), Active: true},
        {FirstName: StringPtr("Bob"), LastName: StringPtr("Smith"), Email: StringPtr("bob.smith@example.com"), Password: StringPtr("abcd12345"), RoleId: UintPtr(3), Active: false},
        {FirstName: StringPtr("Alice"), LastName: StringPtr("Johnson"), Email: StringPtr("alice.johnson@example.com"), Password: StringPtr("abcd12345"), RoleId: UintPtr(1), Active: true},
        {FirstName: StringPtr("Tom"), LastName: StringPtr("Williams"), Email: StringPtr("tom.williams@example.com"), Password: StringPtr("abcd12345"), RoleId: UintPtr(2), Active: false},
	}

	for _, user := range users {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
		user.Password = StringPtr(string(hashedPassword))
		result := db.Create(&user)
		if result.Error != nil {
			panic(result.Error)
		}
	}
}
