package validation

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func ValidImage(address string) error {
	var maxSize int64 = 5 * 1024 * 1024

	fileInfo, err := os.Stat(address)
	if err != nil {
		return errors.New(fmt.Sprintf("Ошибка при получении сведений о изображении: %s", err.Error()))
	}

	if fileInfo.Size() > maxSize {
		return errors.New("Изображение должно быть меньше 5 Мб")
	}

	ext := filepath.Ext(address)
	fmt.Println(ext)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return errors.New("Файл должен иметь формат jpeg, jpg или png")
	}

	return nil
}
