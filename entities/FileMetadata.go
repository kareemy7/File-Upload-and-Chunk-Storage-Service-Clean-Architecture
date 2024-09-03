package entities

import "time"

type FileMetadata struct {
	FileID     string    `json:"file_id"`
	FileName   string    `json:"filename"`
	ChunkSize  int       `json:"chunk_size"`
	NumChunks  int       `json:"num_chunks"`
	UploadDate time.Time `json:"upload_date"`
	ChunkIDs   []string  `json:"chunk_ids"`
}
