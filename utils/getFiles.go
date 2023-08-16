package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetCorrectFiles(root string, extension string) []string {
	files := []string{}
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.Contains(path, extension) {
			files = append(files, path)
		}

		// fmt.Println(path)
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	// fmt.Printf("files found %q", files)
	return files
}
