package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Users struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"not null" `
	LastName  string `gorm:"not null" `
	Email     string `gorm:"unique;not null`
	Password  string `gorm:"not null" `
	RoleId    uint   `gorm:"not null" `
	Active    bool   `gorm:"not null;default:false"`
}

type SqlLogger struct {
	logger.Interface
}

// Trace SQL command
func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n----------------\n", sql)
}

var db *gorm.DB
var err error

func main() {
	// dsn should be env
	// dsn = "<username>:<password>@tcp(<url>:3306)/<dbname>?parseTime=true"
	dsn := "root:P@ssw0rd@tcp(127.0.0.1:3306)/admin_manage_system?parseTime=true"
	dial := mysql.Open(dsn)
	db, err = gorm.Open(dial, &gorm.Config{Logger: &SqlLogger{}})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(Users{})

	// Run this function for the first time to initial users' data.
	// CreateSampleUsers()

	app := fiber.New()
	app.Use(cors.New())
	app.Post("/user", createUser)
	app.Get("/users", getUsers)
	app.Put("/users/:id", updateUserFull)
	app.Patch("/users/:id", updateUserPartial)
	app.Delete("/users/:id", deleteUser)

	app.Listen(":3000")
}

type UsersResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"fullName"`
	RoleId    uint      `json:"roleId"`
	CreatedAt time.Time `json:"createdAt"`
	Active    bool      `json:"active"`
}

func getUsers(c *fiber.Ctx) error {
	var users []Users
	db.Find(&users)

	var usersList []UsersResponse
	for _, user := range users {
		userResponse := UsersResponse{
			ID:        user.ID,
			FullName:  user.FirstName + " " + user.LastName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			RoleId:    user.RoleId,
			Active:    user.Active,
		}
		usersList = append(usersList, userResponse)
	}

	return c.Status(fiber.StatusOK).JSON(usersList)
}

func createUser(c *fiber.Ctx) error {
	user := new(Users)
	fmt.Println(string(c.Body()))
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err,
		})
	}
	user.Password = string(hashedPassword)

	result := db.Create(user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": result.Error,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func updateUserFull(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	var user Users
	result := db.First(&user, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "User not found",
		})
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	result = db.Save(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": result.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

type UpdateUserPartialRequest struct {
	Email  *string `json:"email,omitempty"`
	RoleId *uint   `json:"roleId,omitempty"`
	Active *bool   `json:"active,omitempty"`
}

func updateUserPartial(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	var user Users
	result := db.First(&user, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "User not found",
		})
	}

	var updateData UpdateUserPartialRequest
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	fmt.Println("updateData ->", updateData)

	// Update fields based on request data
	updates := make(map[string]interface{})
	if updateData.Email != nil {
		updates["email"] = *updateData.Email
	}
	if updateData.RoleId != nil {
		updates["role_id"] = *updateData.RoleId
	}
	if updateData.Active != nil {
		updates["active"] = *updateData.Active
	}

	result = db.Model(&user).Updates(updates)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": result.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)

}

func deleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	var user Users
	result := db.First(&user, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg": "User not found",
		})
	}

	result = db.Delete(&user)
	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": result.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "the user was deleted successfully"})
}

func CreateSampleUsers() {
	users := []Users{
		{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com", Password: "abcd12345", RoleId: 1, Active: true},
		{FirstName: "Jane", LastName: "Doe", Email: "jane.doe@example.com", Password: "abcd12345", RoleId: 2, Active: true},
		{FirstName: "Bob", LastName: "Smith", Email: "bob.smith@example.com", Password: "abcd12345", RoleId: 3, Active: false},
		{FirstName: "Alice", LastName: "Johnson", Email: "alice.johnson@example.com", Password: "abcd12345", RoleId: 1, Active: true},
		{FirstName: "Tom", LastName: "Williams", Email: "tom.williams@example.com", Password: "abcd12345", RoleId: 2, Active: false},
		{FirstName: "Sarah", LastName: "Davis", Email: "sarah.davis@example.com", Password: "abcd12345", RoleId: 3, Active: true},
		{FirstName: "Michael", LastName: "Brown", Email: "michael.brown@example.com", Password: "abcd12345", RoleId: 1, Active: false},
		{FirstName: "Emily", LastName: "Wilson", Email: "emily.wilson@example.com", Password: "abcd12345", RoleId: 2, Active: true},
		{FirstName: "David", LastName: "Anderson", Email: "david.anderson@example.com", Password: "abcd12345", RoleId: 3, Active: false},
		{FirstName: "Jessica", LastName: "Thomas", Email: "jessica.thomas@example.com", Password: "abcd12345", RoleId: 1, Active: true},
		{FirstName: "Matthew", LastName: "Jackson", Email: "matthew.jackson@example.com", Password: "abcd12345", RoleId: 2, Active: true},
		{FirstName: "Ashley", LastName: "White", Email: "ashley.white@example.com", Password: "abcd12345", RoleId: 3, Active: false},
		{FirstName: "Daniel", LastName: "Harris", Email: "daniel.harris@example.com", Password: "abcd12345", RoleId: 1, Active: true},
		{FirstName: "Samantha", LastName: "Martin", Email: "samantha.martin@example.com", Password: "abcd12345", RoleId: 2, Active: true},
		{FirstName: "Christopher", LastName: "Thompson", Email: "christopher.thompson@example.com", Password: "abcd12345", RoleId: 3, Active: true},
	}

	for _, user := range users {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)

		result := db.Create(&user)
		if result.Error != nil {
			panic(result.Error)
		}
	}
}
