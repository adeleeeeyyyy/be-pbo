package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"mime/multipart"
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

func SaveUpdatedFile(c *fiber.Ctx, file *multipart.FileHeader, folder string, oldFilePath string) (string, error) {
	// Pastikan folder ada
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	// Buat nama file unik
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	fullPath := filepath.Join(folder, filename)

	// Simpan file baru
	if err := c.SaveFile(file, fullPath); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	// Hapus file lama jika ada
	if oldFilePath != "" {
		_ = os.Remove(oldFilePath)
	}

	return fullPath, nil
}

