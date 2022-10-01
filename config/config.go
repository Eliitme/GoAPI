package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
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
		dbUser = "root"
	}
	if dbPassword == "" {
		dbPassword = ""
	}
	if dbName == "" {
		dbName = "go_api"
	}
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "3306"
	}

	if port == "" {
		port = "8080"
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
