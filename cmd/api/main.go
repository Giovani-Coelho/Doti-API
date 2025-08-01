package main

import (
	"fmt"
	"net/http"

	"github.com/Giovani-Coelho/Doti-API/config"
	"github.com/Giovani-Coelho/Doti-API/config/logger"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/http/middleware"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/http/server"
	database "github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/db"
)

var PORT = config.Env.PORT

func main() {
	logger.Info("About to start application")

	logger.Info("Init the database")
	conn, err := database.NewPostgresDB()

	if err != nil {
		logger.Info("Error initializing database")
		panic(err)
	}

	defer conn.Close()

	router := server.Routes(conn)

	fmt.Printf("Server is running on port :%d", PORT)

	http.ListenAndServe(
		fmt.Sprintf(":%d", PORT),
		middleware.CorsConfig(router),
	)
}
