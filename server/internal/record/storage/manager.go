package storage

import (
	"fmt"
	"io"
	"sync"
)

type StorageManager struct {
	local     *LocalStorage
	s3        *S3Storage
	mu        sync.RWMutex
}

func NewStorageManager(localPath string, s3Config *S3Config) (*StorageManager, error) {
	mgr := &StorageManager{
		local: NewLocalStorage(localPath),
	}

	if s3Config != nil {
		s3, err := NewS3Storage(s3Config)
		if err != nil {
			return nil, fmt.Errorf("init s3 storage failed: %w", err)
		}
		mgr.s3 = s3
	}

	return mgr, nil
}

func (mgr *StorageManager) Store(fileName string, reader io.Reader, useS3 bool) (*FileInfo, error) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()

	if useS3 && mgr.s3 != nil {
		return mgr.s3.Store(fileName, reader)
	}
	return mgr.local.Store(fileName, reader)
}

func (mgr *StorageManager) Delete(filePath string, storageType int) error {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()

	if storageType == 2 && mgr.s3 != nil {
		return mgr.s3.Delete(filePath)
	}
	return mgr.local.Delete(filePath)
}

func (mgr *StorageManager) Exists(filePath string, storageType int) (bool, error) {
	if storageType == 2 && mgr.s3 != nil {
		return mgr.s3.Exists(filePath)
	}
	return mgr.local.Exists(filePath)
}

func (mgr *StorageManager) GetReader(filePath string, storageType int) (io.ReadCloser, error) {
	if storageType == 2 && mgr.s3 != nil {
		return mgr.s3.GetReader(filePath)
	}
	return mgr.local.GetReader(filePath)
}

func (mgr *StorageManager) HasS3() bool {
	return mgr.s3 != nil
}
