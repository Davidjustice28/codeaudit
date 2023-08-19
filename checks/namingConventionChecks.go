package checks

import (
	"bufio"
	"fmt"
	"os"

	"math"
	"regexp"
	"strings"

	"codeAudit/models"

	"github.com/dariubs/percent"
)

func RunNamingConventionCheck(filePath string, style string) models.CheckResult {
	var rg *regexp.Regexp
	camelCaseRegex := regexp.MustCompile("^[a-z0-9]+([A-Z].*)*$")
	snake_case_regex := regexp.MustCompile(`^[a-z0-9](\_[a-z0-9])*$`)
	pascalCaseRegex := regexp.MustCompile(`^[A-Z][a-z0-9]*([A-Z].*)*$`)
	if style == "camel" {
		rg = camelCaseRegex
	}
	if style == "snake" {
		rg = snake_case_regex
	}

	if style == "pascal" {
		rg = pascalCaseRegex
	}
	f, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	varibleCount := 0
	checksPassed := 0
	lines_failed := []int{}
	lineNumber := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lineNumber = lineNumber + 1
		line_items := strings.Split(line, " ")
		for i := 0; i < len(line_items); i++ {
			if line_items[i] == "const" || line_items[i] == "let" {
				varibleCount = varibleCount + 1
				passesChecks := rg.MatchString(line_items[i+1])
				if passesChecks {
					checksPassed = checksPassed + 1
				} else {
					lines_failed = append(lines_failed, lineNumber)
				}
			}

			if line_items[i] == "function" {
				functionName := strings.Split(line_items[i+1], "(")[0]
				varibleCount = varibleCount + 1
				passesChecks := rg.MatchString(functionName)
				if passesChecks {
					checksPassed = checksPassed + 1
				} else {
					lines_failed = append(lines_failed, lineNumber)
				}
			}
		}
	}
	consistency := int(math.Round(percent.PercentOf(checksPassed, varibleCount)))
	for a := 0; a < len(lines_failed); a++ {
	}
	result := models.CheckResult{File: filePath, ChecksMade: varibleCount, ChecksPassed: checksPassed, ConsistencyScore: consistency, LinesFailed: lines_failed}
	return result
}

func MakeNamingConventionChecks(files []string, variableNamingConvention string) models.CompleteCheckResult {
	scores := []int{}
	issues := []models.IssueData{}
	for i := 0; i < len(files); i++ {
		result1 := RunNamingConventionCheck(files[i], variableNamingConvention)
		scores = append(scores, result1.ConsistencyScore)
		for n := 0; n < len(result1.LinesFailed); n++ {
			issueMessage := fmt.Sprintf("Variable or function found on line %d doesn't meet the specified naming convention", result1.LinesFailed[n])
			issue := models.IssueData{Line: result1.LinesFailed[n], File: files[i], Issue: issueMessage}
			issues = append(issues, issue)
		}

	}
	totalScore := 0
	for n := 0; n < len(scores); n++ {
		totalScore = totalScore + scores[n]
	}
	totalScore = totalScore / len(scores)
	completeCheckResult := models.CompleteCheckResult{CheckType: "variable_naming", FilesChecked: files, ConsistencyScores: scores, FinalConsistencyScore: totalScore, IssuesFound: issues}
	return completeCheckResult
}
