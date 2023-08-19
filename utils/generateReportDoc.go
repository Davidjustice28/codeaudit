package utils

import (
	"codeAudit/models"
	"encoding/csv"
	"fmt"
	"os"
	"syscall"
)

func GenerateFailureReport(report models.ConsistencyReport, downloads_path string) bool {
	reportGenerated := true
	if chdirError := syscall.Chdir(downloads_path); chdirError != nil {
		fmt.Printf("change directory error. downloads path used %s\n%s\n", downloads_path, chdirError)
		reportGenerated = false
		return reportGenerated
	}
	file, err := os.Create("CodeAuditReport.csv")
	if err != nil {
		fmt.Println("csv create error", err)
		reportGenerated = false
		return reportGenerated
	}
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()

	rows := [][]string{}
	for _, issue := range report.IssuesFound {
		lineNumber := fmt.Sprintf("%d", issue.Line)
		row := []string{issue.File, lineNumber, issue.Issue}
		rows = append(rows, row)
	}
	if titleError := w.Write([]string{"File", "Line number", "Bug Found"}); titleError != nil {
		fmt.Println("write csv title error", titleError)
		reportGenerated = false
		return reportGenerated
	}
	writeError := w.WriteAll(rows)
	if writeError != nil {
		fmt.Println("write rows in csv error", writeError)
		reportGenerated = false
		return reportGenerated
	}
	return reportGenerated

}
