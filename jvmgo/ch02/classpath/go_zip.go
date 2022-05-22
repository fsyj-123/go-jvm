package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(absPath string) (*ZipEntry, error) {
	path, err := filepath.Abs(absPath)
	if err != nil {
		return nil, err
	}
	return &ZipEntry{path}, nil
}

func (e *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	reader, err := zip.OpenReader(e.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if file.Name == className {
			readCloser, err := file.Open()
			if err != nil {
				return nil, nil, err
			}
			defer readCloser.Close()
			data, err := ioutil.ReadAll(readCloser)
			if err != nil {
				return nil, nil, err
			}
			return data, e, nil
		}
	}
	return nil, nil, errors.New("class not found")
}

func (e *ZipEntry) String() string {
	return e.absPath
}
