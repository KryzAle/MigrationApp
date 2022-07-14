package utils

import (
	"fmt"
	"testfb/envvar"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBCredentials struct {
	Dsn string
}

func GetDatabaseCredentials() DBCredentials {
	user := envvar.DBUser()
	password := envvar.DBPassword()
	name := envvar.DBName()
	host := envvar.DBHost()
	port := envvar.DBPort()

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		user, password, name, host, port)

	return DBCredentials{Dsn: dsn}
}

func GetConnection() *gorm.DB {
	credentials := GetDatabaseCredentials()
	connection, connErr := gorm.Open(postgres.Open(credentials.Dsn), &gorm.Config{})
	if connErr != nil {
		panic("can't connect to db")
	}

	return connection
}
