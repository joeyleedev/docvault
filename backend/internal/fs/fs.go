package fs

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type DocFS struct {
	root string
}

func New(root string) *DocFS {
	return &DocFS{root: root}
}

func (d *DocFS) safePath(name string) (string, error) {
	if strings.Contains(name, "..") {
		return "", errors.New("invalid file name")
	}
	if !strings.HasSuffix(name, ".md") {
		name += ".md"
	}
	return filepath.Join(d.root, name), nil
}

func (d *DocFS) List() ([]os.DirEntry, error) {
	return os.ReadDir(d.root)
}

func (d *DocFS) Read(name string) ([]byte, error) {
	path, err := d.safePath(name)
	if err != nil {
		return nil, err
	}
	return os.ReadFile(path)
}

func (d *DocFS) Write(name string, content []byte) error {
	path, err := d.safePath(name)
	if err != nil {
		return err
	}
	return os.WriteFile(path, content, 0644)
}

func (d *DocFS) Delete(name string) error {
	path, err := d.safePath(name)
	if err != nil {
		return err
	}
	return os.Remove(path)
}
