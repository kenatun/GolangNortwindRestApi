package database

import "database/sql"

func InitDB() *sql.DB {
	connectionString := "root:root@tcp(localhost:3307)/northwind"
	databaseConnection, err := sql.Open("mysql", connectionString)
	//se ejecutara cuando termina de ejecutar el main para cerrar
	if err != nil {
		panic(err.Error()) //Error handling, manejo de errores
	}
	return databaseConnection
}
