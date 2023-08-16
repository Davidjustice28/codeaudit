package models

type IssueData struct {
	Line  int
	File  string
	Issue string
}
type CheckResult struct {
	File             string `json:"file"`
	ChecksMade       int    `json:"checks_made"`
	ChecksPassed     int    `json:"checks_passed"`
	LinesFailed      []int  `json:"lines_failed"`
	ConsistencyScore int    `json:"consistency_score"`
}

type CompleteCheckResult struct {
	CheckType             string      `json:"check_type"`
	FilesChecked          []string    `json:"files_checked"`
	ConsistencyScores     []int       `json:"consistency_scores"`
	FinalConsistencyScore int         `json:"final_consistency_score"`
	IssuesFound           []IssueData `json:"issues_found"`
}

type ConsistencyReport struct {
	Checks                   []string              `json:"checks"`
	CheckResults             []CompleteCheckResult `json:"check_results"`
	IssuesFound              []IssueData           `json:"issues_found"`
	CodeBaseConsistencyScore int                   `json:"codebase_consistency_score" `
}
