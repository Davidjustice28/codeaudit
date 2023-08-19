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

func runCharacterCountCheck(path string, charactersAllowed int) models.CheckResult {
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
		characterAmount := len(strings.Split(line, ""))
		// fmt.Printf("number of characters found on line %d in file %s: %d\n ", lineNumber, path, characterAmount)
		if characterAmount > charactersAllowed {
			lines_failed = append(lines_failed, lineNumber)
		} else {
			checksPassed = checksPassed + 1
		}
	}
	consistency := int(math.Round(percent.PercentOf(checksPassed, lineNumber)))

	result := models.CheckResult{File: path, ChecksMade: lineNumber, ChecksPassed: checksPassed, ConsistencyScore: consistency, LinesFailed: lines_failed}
	return result
}

func MakeCharacterCountChecks(files []string, charactersAllowed int) models.CompleteCheckResult {
	scores := []int{}
	issues := []models.IssueData{}
	for i := 0; i < len(files); i++ {
		result := runCharacterCountCheck(files[i], charactersAllowed)
		scores = append(scores, result.ConsistencyScore)
		for n := 0; n < len(result.LinesFailed); n++ {
			issueMessage := fmt.Sprintf("Line %d exceeds the specified character count limit", result.LinesFailed[n])
			issue := models.IssueData{Line: result.LinesFailed[n], File: files[i], Issue: issueMessage}
			issues = append(issues, issue)
		}

	}
	totalScore := 0
	for n := 0; n < len(scores); n++ {
		totalScore = totalScore + scores[n]
	}
	totalScore = totalScore / len(scores)
	completeCheckResult := models.CompleteCheckResult{CheckType: "character_count", FilesChecked: files, ConsistencyScores: scores, FinalConsistencyScore: totalScore, IssuesFound: issues}
	return completeCheckResult
}
