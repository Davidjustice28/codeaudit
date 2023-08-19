package checks

import (
	"bufio"
	"codeAudit/models"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/dariubs/percent"
)

func runIndentationCheck(path string, spaces int) models.CheckResult {
	checksPassed := 0
	lines_failed := []int{}
	lineNumber := 0

	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lineNumber = lineNumber + 1
		codeWithoutIndentation := strings.TrimLeft(line, " ")
		numberOfIndention := len(line) - len(codeWithoutIndentation)
		// fmt.Printf("indention found on line %d in file %s: %d spaces\n ", lineNumber, path, numberOfIndention)
		if numberOfIndention%spaces != 0 {
			lines_failed = append(lines_failed, lineNumber)
		} else {
			checksPassed = checksPassed + 1
		}
	}
	consistency := int(math.Round(percent.PercentOf(checksPassed, lineNumber)))

	result := models.CheckResult{File: path, ChecksMade: lineNumber, ChecksPassed: checksPassed, ConsistencyScore: consistency, LinesFailed: lines_failed}
	return result
}

func MakeIndentionChecks(files []string, numberOfSpaces int) models.CompleteCheckResult {
	scores := []int{}
	issues := []models.IssueData{}
	for i := 0; i < len(files); i++ {
		result1 := runIndentationCheck(files[i], numberOfSpaces)
		scores = append(scores, result1.ConsistencyScore)
		for n := 0; n < len(result1.LinesFailed); n++ {
			issueMessage := fmt.Sprintf("Line %d doesn't meet the specified indention requirements", result1.LinesFailed[n])
			issue := models.IssueData{Line: result1.LinesFailed[n], File: files[i], Issue: issueMessage}
			issues = append(issues, issue)
		}

	}
	totalScore := 0
	for n := 0; n < len(scores); n++ {
		totalScore = totalScore + scores[n]
	}
	totalScore = totalScore / len(scores)
	completeCheckResult := models.CompleteCheckResult{CheckType: "indentation", FilesChecked: files, ConsistencyScores: scores, FinalConsistencyScore: totalScore, IssuesFound: issues}
	return completeCheckResult
}
