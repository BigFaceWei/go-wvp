package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	basePath string
}

func NewLocalStorage(basePath string) *LocalStorage {
	os.MkdirAll(basePath, 0755)
	return &LocalStorage{basePath: basePath}
}

func (s *LocalStorage) Store(fileName string, reader io.Reader) (*FileInfo, error) {
	filePath := filepath.Join(s.basePath, fileName)
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("create directory failed: %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("create file failed: %w", err)
	}
	defer file.Close()

	written, err := io.Copy(file, reader)
	if err != nil {
		return nil, fmt.Errorf("write file failed: %w", err)
	}

	return &FileInfo{
		FileName:    fileName,
		FilePath:    filePath,
		FileSize:    written,
		StorageType: 1,
	}, nil
}

func (s *LocalStorage) Delete(filePath string) error {
	return os.Remove(filepath.Join(s.basePath, filePath))
}

func (s *LocalStorage) Exists(filePath string) (bool, error) {
	_, err := os.Stat(filepath.Join(s.basePath, filePath))
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (s *LocalStorage) GetReader(filePath string) (io.ReadCloser, error) {
	return os.Open(filepath.Join(s.basePath, filePath))
}

func (s *LocalStorage) GetType() int {
	return 1
}
