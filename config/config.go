package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func Getenv() Config {
	e := godotenv.Load()

	if e != nil {
		fmt.Print(e)
	}

	dbUser := os.Getenv("db_user")
	dbPassword := os.Getenv("db_password")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	port := os.Getenv("port")

	if dbUser == "" {
		fmt.Println("db_user is not set")
	}
	if dbPassword == "" {
		fmt.Println("db_password is not set")
	}
	if dbName == "" {
		fmt.Println("db_name is not set")
	}
	if dbHost == "" {
		fmt.Println("db_host is not set")
	}
	if dbPort == "" {
		dbPort = "3306"
	}

	if port == "" {
		port = "9000"
	}

	return Config{
		Port:       port,
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
	}
}
