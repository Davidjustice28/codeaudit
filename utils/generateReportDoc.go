package utils

import (
	"codeAudit/models"
	"encoding/csv"
	"fmt"
	"os"
)

func GenerateFailureReport(report models.ConsistencyReport) bool {
	reportGenerated := true
	file, err := os.Create("CodeAuditReport.csv")
	if err != nil {
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
		reportGenerated = false
		return reportGenerated
	}
	writeError := w.WriteAll(rows)
	if writeError != nil {
		reportGenerated = false
		return reportGenerated
	}
	return reportGenerated

}
