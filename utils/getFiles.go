package utils

import (
	// "io/fs"

	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetCorrectFiles(extension string, files []string, root string) []string {
	base := root

	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() && f.Name() != strings.Split(root, "./")[1] {
			base = fmt.Sprintf("%s/%s", base, f.Name())
			// fmt.Println(base)
		}
		if filepath.Ext(path) == extension {
			fileRelativePath := fmt.Sprintf("%s/%s", base, f.Name())
			files = append(files, fileRelativePath)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return files
}
