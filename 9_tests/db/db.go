package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Golang_MIPT/9_tests/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "transactionsdb"
)

var DB *gorm.DB

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Создаём таблицу users
	createExtensionsChkPass := `
		CREATE EXTENSION IF NOT EXISTS "pgcrypto";
	`
	_, err = db.Exec(createExtensionsChkPass)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Extension 'chkpass' created successfully!")

	createUsersTableQuery := `
		CREATE TABLE IF NOT EXISTS users
		(
		u_id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		email VARCHAR(100) NOT NULL,
		u_password text NOT NULL, 
		unique(email)
		)
	`
	createTransactionsTableQuery := `
		CREATE TABLE IF NOT EXISTS transactions
		(
		id SERIAL PRIMARY KEY,
		u_id int REFERENCES users(u_id), 
		amount NUMERIC, 
		t_currency varchar(3), 
		t_type varchar(15), 
		t_category varchar(50), 
		t_date varchar(15), 
		description VARCHAR(100) NOT NULL

	)`
	_, err = db.Exec(createUsersTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table 'users' created successfully!")
	_, err = db.Exec(createTransactionsTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table 'transactions' created successfully!")
}

func Connect() {
	var err error
	//	dsn := "host=localhost user=user password=password dbname=transactionsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", host, port, user, password, dbname)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database\n", err)
		os.Exit(2)
	}
	log.Println("Connection Opened to Database")
}

func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.Transaction{})
	log.Println("Database Migration Completed")
}
