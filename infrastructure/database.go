package infrastructure

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

type Database struct {
	DB *leveldb.DB
}

func NewDatabase(path string) (*Database, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}
	return &Database{DB: db}, nil
}

func (d *Database) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
}
