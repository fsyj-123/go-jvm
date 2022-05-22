package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) (*DirEntry, error) {
	// convert path to absolute path
	absDir, err := filepath.Abs(path)
	if err != nil {
		// use error instead of panic
		return nil, err
	}
	return &DirEntry{absDir: absDir}, nil
}
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	classFilePath := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(classFilePath)
	return data, self, err
}
func (self *DirEntry) String() string {
	return self.absDir
}
