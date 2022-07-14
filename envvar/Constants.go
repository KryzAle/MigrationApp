package envvar

import "os"

func DBUser() string {
	user, exists := os.LookupEnv("POSTGRES_USER")
	if !exists {
		user = "postgres"
	}
	return user
}

func DBPassword() string {
	password, exists := os.LookupEnv("POSTGRES_PASSWORD")
	if !exists {
		password = "postgres"
	}
	return password
}

func DBName() string {
	dbname, exists := os.LookupEnv("POSTGRES_DB_NAME")
	if !exists {
		dbname = "projectstaff"
	}
	return dbname
}
func DBHost() string {
	host, exists := os.LookupEnv("POSTGRES_HOST")
	if !exists {
		host = "127.0.0.1"
	}
	return host
}

func DBPort() string {
	port, exists := os.LookupEnv("POSTGRES_PORT")
	if !exists {
		port = "5432"
	}
	return port
}

func ApiPort() string {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8081"
	}
	return port
}
