package storage

import (
	"errors"
	"io"
)

type FileInfo struct {
	FileName   string
	FilePath   string
	FileSize   int64
	StorageType int
}

type Storage interface {
	Store(fileName string, reader io.Reader) (*FileInfo, error)
	Delete(filePath string) error
	Exists(filePath string) (bool, error)
	GetReader(filePath string) (io.ReadCloser, error)
	GetType() int
}
