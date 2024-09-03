package controllers

import (
	"File-Upload-and-Chunk-Storage-Service-Clean-Architecture/usecases"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	UseCase *usecases.FileUseCase
}

func NewFileController(uc *usecases.FileUseCase) *FileController {
	return &FileController{UseCase: uc}
}

func (ctrl *FileController) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file_")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer f.Close()

	fileData, err := io.ReadAll(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	metadata, err := ctrl.UseCase.UploadFile(fileData, file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "file_id": metadata.FileID})
}

func (ctrl *FileController) DownloadFile(c *gin.Context) {
	fileID := c.Param("file_id")

	fileData, metadata, err := ctrl.UseCase.DownloadFile(fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=\""+metadata.FileName+"\"")
	c.Data(http.StatusOK, "application/octet-stream", fileData)
}
