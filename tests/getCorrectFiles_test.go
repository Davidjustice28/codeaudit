package tests

import (
	"codeAudit/utils"
	"fmt"
	"reflect"
	"testing"
)

func TestGetCorrectFilesReturnsExpectedFiles(t *testing.T) {
	files := utils.GetCorrectFiles("./mock_directory", ".js")
	testFiles := []string{"mock_directory/index.js", "mock_directory/main.js"}
	filesMatch := reflect.DeepEqual(files, testFiles)
	if len(files) != len(testFiles) {
		t.Errorf("Incorrect number of files returned. Expected %d but received %d", 2, len(files))
	}
	if !filesMatch {

		expectedFiles := "index.js, main.js"
		returnedFiles := ""
		for _, file := range files {
			returnedFiles = fmt.Sprintf("%s, %s", returnedFiles, file)
		}
		t.Errorf("The util didn't return the correct files expect. Expected %s but received %s", expectedFiles, returnedFiles)
	}
}
