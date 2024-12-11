package filesage

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

const uploadDir = "uploads"

// should probably belong in client code when starting the server instead
func ensureDirExists() {
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir("uploads", 0755)
	}
}

func SaveFile(file multipart.File, filename string) error {
	ensureDirExists()

	ext := filepath.Ext(filename)
	nameWithoutExt := strings.TrimSuffix(filename, ext)
	newFilename := nameWithoutExt + uuid.New().String() + ext

	dst, err := os.Create(filepath.Join(uploadDir, newFilename))
	if err != nil {
		return err
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if _, err = dst.Write(fileBytes); err != nil {
		return err
	}

	return nil
}
