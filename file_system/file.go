package filesystem

import (
	"io"
	"os"
)

type FileSystem struct {
	filepath string
}

func NewFileSystem(file string) *FileSystem {
	return &FileSystem{
		filepath: file,
	}
}

func (f *FileSystem) Upload(sourceURI string, destinationPath string) (bool, string) {
	// Open source file
	sourceFile, err := os.Open(sourceURI)
	if err != nil {
		return false, "Error happened"
	}

	defer sourceFile.Close()

	// create destination file
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return false, "Error happened on destination"
	}

	defer destinationFile.Close()

	// copy the content from source to destination
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return false, "Error happened copy"
	}
	return true, "File upload successful"
}

func TestImage() bool {
	return true
}
