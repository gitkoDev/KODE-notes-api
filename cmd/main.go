package main

import (
	"github.com/gitkoDev/KODE-test-task/cmd/api"
	"github.com/gitkoDev/KODE-test-task/db"
	"github.com/gitkoDev/KODE-test-task/internal/handler"
	"github.com/gitkoDev/KODE-test-task/internal/repository"
	"github.com/gitkoDev/KODE-test-task/internal/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// Config stage
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatal("error loading env file:", err)
	}

	// DB  connection stage
	db, err := db.PostgresConnection()
	if err != nil {
		logrus.Fatal("error connecting to database:", err)
	}
	logrus.Println("db connected")
	defer db.Close()

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	// Server connection stage
	server := new(api.APIServer)
	if err := server.Run("8080", handler.RegisterRoutes()); err != nil {
		logrus.Errorln("error running server:", err)
	}

	// port := ":8080"
	// server := api.NewAPIServer(port, db)
	// logrus.Printf("server running on port %s", port)
	// if err := server.Run(); err != nil {
	// 	logrus.Fatal("error connecting to server:", err)
	// }
}
