package main

import (
	"File-Upload-and-Chunk-Storage-Service-Clean-Architecture/controllers"
	"File-Upload-and-Chunk-Storage-Service-Clean-Architecture/infrastructure"
	"File-Upload-and-Chunk-Storage-Service-Clean-Architecture/repositories"
	"File-Upload-and-Chunk-Storage-Service-Clean-Architecture/usecases"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db, err := infrastructure.NewDatabase("./data")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Initialize repository with the database
	fileRepo := repositories.NewLevelDBRepository(db.DB)

	// Initialize use case with the repository
	fileUseCase := usecases.NewFileUseCase(fileRepo)

	// Initialize controller with the use case
	fileController := controllers.NewFileController(fileUseCase)

	// Initialize Gin router
	r := gin.Default()

	// Define routes
	r.POST("/upload", fileController.UploadFile)
	r.GET("/download/:file_id", fileController.DownloadFile)

	// Gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		db.Close()
	}()

	r.Run() // Start the server
}
