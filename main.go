package main

import (
	"GolangNortwindRestApi/database"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	databaseConnection := database.InitDB()
	//Logica
	defer databaseConnection.Close()
	fmt.Println(databaseConnection)
}
