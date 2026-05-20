package main

import (
	"database/sql"
	"log"

	"rest/db/api"
	"rest/dto"
	"rest/utils"

	_ "github.com/lib/pq"
)

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Error, no se puede cargar la configuración", err)
	}
	config.DBDriver = "postgres"
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error, no se pudo conectar a la base de datos", err)
	}

	dbtx := dto.New(conn)
	server, err := api.NewServer(dbtx, config.Secret)
	if err != nil {
		log.Fatal("No se puede iniciar el servidor", err)
	}
	err = server.Start(config.ServerURL)
	if err != nil {
		log.Fatal("Error fatal", err)
	}
}
