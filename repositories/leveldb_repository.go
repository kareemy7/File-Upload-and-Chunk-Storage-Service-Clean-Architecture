package repositories

import (
	"File-Upload-and-Chunk-Storage-Service-Clean-Architecture/entities"
	"encoding/json"

	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDBRepository struct {
	DB *leveldb.DB
}

func NewLevelDBRepository(db *leveldb.DB) *LevelDBRepository {
	return &LevelDBRepository{DB: db}
}

func (repo *LevelDBRepository) SaveChunk(chunkID string, chunkData []byte) error {
	return repo.DB.Put([]byte(chunkID), chunkData, nil)
}

func (repo *LevelDBRepository) SaveMetadata(metadata entities.FileMetadata) error {
	data, err := json.Marshal(metadata)
	if err != nil {
		return err
	}
	return repo.DB.Put([]byte("file_metadata_"+metadata.FileID), data, nil)
}

func (repo *LevelDBRepository) GetMetadata(fileID string) (entities.FileMetadata, error) {
	data, err := repo.DB.Get([]byte("file_metadata_"+fileID), nil)
	if err != nil {
		return entities.FileMetadata{}, err
	}

	var metadata entities.FileMetadata
	err = json.Unmarshal(data, &metadata)
	if err != nil {
		return entities.FileMetadata{}, err
	}

	return metadata, nil
}

func (repo *LevelDBRepository) GetChunk(chunkID string) ([]byte, error) {
	return repo.DB.Get([]byte(chunkID), nil)
}
