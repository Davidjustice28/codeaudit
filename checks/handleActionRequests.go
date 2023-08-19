package checks

import (
	"codeAudit/models"
	"codeAudit/utils"
	"fmt"
)

func PerformAllChecks(root_directory string, fileType string, name_convention string, indentation int, char_count int) models.ConsistencyReport {
	files := utils.GetCorrectFiles(root_directory, fileType)

	// FIXME: make naming convention style an optional terminal arg not hard coded
	result1 := MakeNamingConventionChecks(files, name_convention)
	// FIXME: make spaces an optional terminal arg not hard coded
	result2 := MakeIndentionChecks(files, indentation)
	// FIXME: make spaces an optional terminal arg not hard coded
	result3 := MakeCharacterCountChecks(files, char_count)

	results := []models.CompleteCheckResult{result1, result2, result3}
	checksMade := []string{"Variable Name Casing", "Indentation", "Character Count Limit Per Line"}
	issues := []models.IssueData{}
	totalScore := (result1.FinalConsistencyScore + result2.FinalConsistencyScore + result3.FinalConsistencyScore) / 3

	issues = append(issues, result1.IssuesFound...)
	issues = append(issues, result2.IssuesFound...)
	issues = append(issues, result3.IssuesFound...)

	// FIXME: use full report object to create reports
	fullReport := models.ConsistencyReport{IssuesFound: issues, CheckResults: results, Checks: checksMade, CodeBaseConsistencyScore: totalScore}
	for r := 0; r < len(results); r++ {
		resultEntry := results[r]
		fmt.Printf("Check \"%s\" score: %d%%\n", resultEntry.CheckType, resultEntry.FinalConsistencyScore)
	}
	fmt.Printf("\nTotal Syntax Consistency Score - - - >  %d%%\n", fullReport.CodeBaseConsistencyScore)
	return fullReport

}

func PerformCharacterCountChecks(root_directory string, fileType string, char_count int) models.ConsistencyReport {
	files := utils.GetCorrectFiles(root_directory, fileType)

	// FIXME: make spaces an optional terminal arg not hard coded
	result1 := MakeCharacterCountChecks(files, char_count)

	results := []models.CompleteCheckResult{result1}
	checksMade := []string{"Character Count Limit Per Line"}
	issues := []models.IssueData{}
	issues = append(issues, result1.IssuesFound...)
	totalScore := result1.FinalConsistencyScore

	// FIXME: use full report object to create reports
	fullReport := models.ConsistencyReport{IssuesFound: issues, CheckResults: results, Checks: checksMade, CodeBaseConsistencyScore: totalScore}
	for r := 0; r < len(results); r++ {
		resultEntry := results[r]
		fmt.Printf("Check \"%s\" score: %d%%\n", resultEntry.CheckType, resultEntry.FinalConsistencyScore)
	}
	fmt.Printf("\nTotal Syntax Consistency Score - - - >  %d%%\n", fullReport.CodeBaseConsistencyScore)
	return fullReport
}

func PerformVariableNamingChecks(root_directory string, fileType string, name_convention string) models.ConsistencyReport {
	files := utils.GetCorrectFiles(root_directory, fileType)

	// FIXME: make naming convention style an optional terminal arg not hard coded
	result1 := MakeNamingConventionChecks(files, name_convention)

	results := []models.CompleteCheckResult{result1}
	checksMade := []string{"Variable Name Casing"}
	issues := []models.IssueData{}
	issues = append(issues, result1.IssuesFound...)
	totalScore := result1.FinalConsistencyScore

	// FIXME: use full report object to create reports
	fullReport := models.ConsistencyReport{IssuesFound: issues, CheckResults: results, Checks: checksMade, CodeBaseConsistencyScore: totalScore}
	for r := 0; r < len(results); r++ {
		resultEntry := results[r]
		fmt.Printf("Check \"%s\" score: %d%%\n", resultEntry.CheckType, resultEntry.FinalConsistencyScore)
	}
	fmt.Printf("\nTotal Syntax Consistency Score - - - >  %d%%\n", fullReport.CodeBaseConsistencyScore)
	return fullReport
}

func PerformIndentationChecks(root_directory string, fileType string, indentation int) models.ConsistencyReport {
	files := utils.GetCorrectFiles(root_directory, fileType)

	// FIXME: make spaces an optional terminal arg not hard coded
	result1 := MakeIndentionChecks(files, indentation)

	results := []models.CompleteCheckResult{result1}
	checksMade := []string{"Indentation"}
	issues := []models.IssueData{}
	issues = append(issues, result1.IssuesFound...)
	totalScore := result1.FinalConsistencyScore

	// FIXME: use full report object to create reports
	fullReport := models.ConsistencyReport{IssuesFound: issues, CheckResults: results, Checks: checksMade, CodeBaseConsistencyScore: totalScore}
	for r := 0; r < len(results); r++ {
		resultEntry := results[r]
		fmt.Printf("Check \"%s\" score: %d%%\n", resultEntry.CheckType, resultEntry.FinalConsistencyScore)
	}
	fmt.Printf("\nTotal Syntax Consistency Score - - - >  %d%%\n", fullReport.CodeBaseConsistencyScore)
	return fullReport
}
