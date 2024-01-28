package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlBackend struct {
	DB *gorm.DB
}

type Backend interface {
	Initialize() (*gorm.DB, error)
}

// InitDB initializes the database connection

func (m *MysqlBackend) InitDB() (gormDB *gorm.DB, err error) {
	// Database connection parameters
	username := "root"
	password := os.Getenv("DB_PASSWORD")
	host := "localhost"
	port := "3306"
	dbName := "training"

	// Create the connection string
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)

	// Open a database connection
	conn, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return gormDB, err
	}

	db, err := gorm.Open(gormmysql.New(gormmysql.Config{
		Conn: conn,
	}))

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("database initialization error. Error: %s", err)
	}
	err = sqlDB.Ping()
	if err != nil {
		fmt.Println("Error pinging the database:", err)
		return gormDB, err
	}

	m.DB = db

	if err != nil {
		return nil, fmt.Errorf("database initialization error. Error: %s", err)
	}

	fmt.Println("Successfully connected to the database!")
	//return db, nil
	return m.DB, nil
}
