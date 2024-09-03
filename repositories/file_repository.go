package repositories

import "File-Upload-and-Chunk-Storage-Service-Clean-Architecture/entities"

type FileRepository interface {
	SaveChunk(chunkID string, chunkData []byte) error
	SaveMetadata(metadata entities.FileMetadata) error
	GetMetadata(fileID string) (entities.FileMetadata, error)
	GetChunk(chunkID string) ([]byte, error)
}
