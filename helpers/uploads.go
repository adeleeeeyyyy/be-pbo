package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"github.com/gofiber/fiber/v2"
)

// SaveUploadedFile menyimpan file ke folder tertentu dan mengembalikan path-nya.
func SaveUploadedFile(c *fiber.Ctx, fieldName, folder string) (string, error) {
	file, err := c.FormFile(fieldName)
	if err != nil {
		return "", err
	}

	if file == nil {
		return "", fmt.Errorf("file is empty")
	}

	// buat folder jika belum ada
	os.MkdirAll(folder, os.ModePerm)

	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
	path := filepath.Join(folder, filename)

	if err := c.SaveFile(file, path); err != nil {
		return "", err
	}

	return path, nil
}

func UpdateUploadedFile(c *fiber.Ctx, fieldName, uploadDir, oldPath string) (string, error) {
    file, err := c.FormFile(fieldName)
    if err != nil || file == nil {
        return oldPath, nil // tidak upload baru → pakai image lama
    }

    // hapus file lama jika ada
    if oldPath != "" {
        os.Remove(oldPath)
    }

    // upload baru
    newPath, err := SaveUploadedFile(c, fieldName, uploadDir)
    if err != nil {
        return oldPath, err
    }

    return newPath, nil
}
