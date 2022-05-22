package classpath

import (
	"fmt"
	"os"
	"strings"
)

const pathSeparator = string(os.PathListSeparator)

// Entry entry is the interface of classpath jar
//
// readClass readClass is a method of readClass according to the variable of className， separated by '/'
// the byte is the binary data of class file
// String just like the toString of java
type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	// the concrete class
	if strings.Contains(path, pathSeparator) {
		return newCompositeEntry(path)
	}
	// use * to match all the jar in the directory
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	// the zip file
	if isCompressPackage(path) {
		entry, err := newZipEntry(path)
		if err != nil {
			return nil
		}
		return entry
	}
	entry, err := newDirEntry(path)
	if err != nil {
		// TODO: handle or throw error
	}
	return entry
}

func isCompressPackage(path string) bool {
	return strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".zip")
}
