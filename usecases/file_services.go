package usecases

import (
	"File-Upload-and-Chunk-Storage-Service-Clean-Architecture/entities"
	"File-Upload-and-Chunk-Storage-Service-Clean-Architecture/repositories"
	"time"

	"github.com/google/uuid"
)

type FileUseCase struct {
	Repo repositories.FileRepository
}

func NewFileUseCase(repo repositories.FileRepository) *FileUseCase {
	return &FileUseCase{Repo: repo}
}

func (uc *FileUseCase) UploadFile(fileData []byte, fileName string) (entities.FileMetadata, error) {
	chunkSize := 1024 * 1024
	numChunks := len(fileData) / chunkSize
	if len(fileData)%chunkSize != 0 {
		numChunks++
	}

	chunkIDs := make([]string, 0, numChunks)
	for i := 0; i < numChunks; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(fileData) {
			end = len(fileData)
		}

		chunkID := uuid.New().String()
		err := uc.Repo.SaveChunk(chunkID, fileData[start:end])
		if err != nil {
			return entities.FileMetadata{}, err
		}

		chunkIDs = append(chunkIDs, chunkID)
	}

	fileMetadata := entities.FileMetadata{
		FileID:     uuid.New().String(),
		FileName:   fileName,
		ChunkSize:  chunkSize,
		NumChunks:  numChunks,
		UploadDate: time.Now(),
		ChunkIDs:   chunkIDs,
	}

	err := uc.Repo.SaveMetadata(fileMetadata)
	if err != nil {
		return entities.FileMetadata{}, err
	}

	return fileMetadata, nil
}

func (uc *FileUseCase) DownloadFile(fileID string) ([]byte, entities.FileMetadata, error) {
	metadata, err := uc.Repo.GetMetadata(fileID)
	if err != nil {
		return nil, entities.FileMetadata{}, err
	}

	var fileData []byte
	for _, chunkID := range metadata.ChunkIDs {
		chunkData, err := uc.Repo.GetChunk(chunkID)
		if err != nil {
			return nil, entities.FileMetadata{}, err
		}
		fileData = append(fileData, chunkData...)
	}

	return fileData, metadata, nil
}
