package service

import (
	"errors"

	"docvault-backend/internal/api"
	"docvault-backend/internal/fs"
)

type Service struct {
	fs *fs.DocFS
}

func New(fs *fs.DocFS) *Service {
	return &Service{fs: fs}
}

// ListDocs returns metadata for all documents
func (s *Service) ListDocs() ([]api.DocumentMeta, error) {
	entries, err := s.fs.List()
	if err != nil {
		return nil, err
	}

	var docs []api.DocumentMeta
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		info, err := e.Info()
		if err != nil {
			continue
		}
		docs = append(docs, api.DocumentMeta{
			Name:      e.Name(),
			UpdatedAt: info.ModTime(),
			Size:      info.Size(),
		})
	}
	return docs, nil
}

// ReadDoc returns the content of a document
func (s *Service) ReadDoc(name string) ([]byte, error) {
	return s.fs.Read(name)
}

// SaveDoc updates an existing document with new content
func (s *Service) SaveDoc(name string, content []byte) error {
	return s.fs.Write(name, content)
}

// CreateDoc creates a new document with the given name and content
func (s *Service) CreateDoc(name string, content []byte) error {
	// Check if file already exists
	if _, err := s.fs.Read(name); err == nil {
		return errors.New("document already exists")
	}
	return s.fs.Write(name, content)
}

// DeleteDoc removes a document
func (s *Service) DeleteDoc(name string) error {
	return s.fs.Delete(name)
}
