package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type db struct {
	// Driver string
	Host                 string
	Port                 int64
	User, Password, Name string
	Migrate              bool
}

var Database db

func (config db) GetDsn() string {
	if config.User == "" && config.Password == "" {
		return fmt.Sprintf("tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Host, config.Port, config.Name)
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.Name)
}

func init() {
	Database.Host = os.Getenv("DB_HOST")
	Database.Port, _ = strconv.ParseInt(os.Getenv("DB_PORT"), 10, 32)
	Database.User = os.Getenv("DB_USER")
	Database.Password = os.Getenv("DB_PASS")
	Database.Name = os.Getenv("DB_NAME")

	if strings.TrimSpace(Database.Name) == "" {
		Database.Name = "test-restful-product"
	}

	Database.Migrate, _ = strconv.ParseBool(os.Getenv("DB_MIGRATE"))
}
