package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.office.opendns.com/selthiru/Golang-Training/models"
	"gorm.io/gorm"
)

// User model
type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Define your handlers
// type HandlerFunc func(Context) error
func IndexHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Ok!")
}

// Get all users
func GetAllUsersHandler(c echo.Context, db *gorm.DB) ([]models.User, error) {
	var users []models.User // Use the User model from the imported package

	if err := db.Find(&users).Error; err != nil {
		return users, c.String(http.StatusInternalServerError, "Failed to fetch users")
	}
	return users, nil
	//return c.JSON(http.StatusOK, users)
}

// CreateUserHandler creates a new user record with JSON data from the request body
func CreateUserHandler(c echo.Context, db *gorm.DB) error {
	// Create a new User struct to hold the JSON data
	var newUser User

	// Bind the JSON data from the request body to the newUser struct
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Check if the user already exists based on the email
	var existingUser User
	if err := db.Where("id = ?", newUser.ID).First(&existingUser).Error; err == nil {
		// User with the same email already exists
		return c.String(http.StatusConflict, "User with this ID already exists")
	}

	// Insert the new user record into the database
	if err := db.Create(&newUser).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Failed to create user"+err.Error())
	}

	return c.JSON(http.StatusCreated, newUser)
}

// UpdateUserHandler updates an existing user record with JSON data from the request body
func UpdateUserHandler(c echo.Context, db *gorm.DB) error {
	// Get the user ID from the URL parameter
	userID := c.QueryParam("id")

	// Check if userID is nil or empty
	if userID == "" {
		return c.String(http.StatusBadRequest, "User ID is missing in the query parameters")
	}

	// Fetch the user record from the database by ID
	var existingUser User
	if err := db.First(&existingUser, userID).Error; err != nil {
		return c.String(http.StatusNotFound, "User not found")
	}

	// Bind the JSON data from the request body to the existingUser struct
	if err := c.Bind(&existingUser); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Update the user record in the database
	if err := db.Save(&existingUser).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Failed to update user")
	}

	return c.JSON(http.StatusOK, existingUser)
}

func DeleteUserHandler(c echo.Context, db *gorm.DB) error {
	// Get the user ID from the URL parameter
	userID := c.QueryParam("id")

	// Check if userID is nil or empty
	if userID == "" {
		return c.String(http.StatusBadRequest, "User ID is missing in the query parameters")
	}

	// Fetch the user record from the database by ID
	var existingUser User
	if err := db.First(&existingUser, userID).Error; err != nil {
		return c.String(http.StatusNotFound, "User not found")
	}

	// Delete the user record from the database
	if err := db.Delete(&existingUser).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete user")
	}

	return c.String(http.StatusOK, "User deleted successfully")
}
