package main

import (
	"github.com/GoFurself/devtools/pkg/userops"
)

func main() {
	userops.UserServiceFactory(userops.SQLite, userops.WithDataSourceName("user.db"))

}
