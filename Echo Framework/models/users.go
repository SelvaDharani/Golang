package models

// User model
type User struct {
	ID    uint `gorm:"primary_key"`
	Name  string
	Email string
}

func (User) TableName() string {
	return "users" // Specify the custom table name here
}
