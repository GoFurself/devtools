package main

import (
	"fmt"

	"github.com/GoFurself/devtools/pkg/userops"
)

func main() {
	userops.UserServiceFactory(userops.SQLite, userops.WithDataSourceName("user.db"))

	fmt.Println("Starting the application...")
}
