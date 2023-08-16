package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"example.com/checks"
	"example.com/models"
	"example.com/utils"
)

func main() {
	var fileType string
	var action string

	terminalArgs := os.Args[1:]
	if len(terminalArgs) == 0 {
		fileType = ".js"
		action = "runchecks"
	}

	if len(terminalArgs) == 1 {
		// println("action", terminalArgs[0])
		action = terminalArgs[0]
		fileType = ".js"
	}

	if len(terminalArgs) == 2 {
		// println("action", terminalArgs[0], "flag", terminalArgs[1])
		action = terminalArgs[0]
		isFlag, err := regexp.MatchString("-", terminalArgs[1])

		if err != nil {
			fmt.Println("err 2")

			panic(err)
		}

		if isFlag {
			flagValue := strings.Split(terminalArgs[1], "-")[1]
			options := strings.Split(flagValue, "")
			switch options[0] {
			case "j":
				fileType = ".js"
			case "p":
				fileType = ".py"
			case "t":
				fileType = ".ts"
			case "g":
				fileType = ".go"
			default:
				fileType = ".js"
			}

		} else {
			fmt.Println("err 3")

			panic("Unknown value for flag found")
		}
	}

	if action == "runchecks" {
		func() {
			root_directory := "./"
			files := utils.GetCorrectFiles(root_directory, fileType)

			// fmt.Printf("%s files found to test\n", ".js")
			// for i := 0; i < len(files); i++ {
			// 	println(files[i])
			// }
			result := checks.MakeConventionChecks(files, "camelCase")
			// bytes, err := json.Marshal(result)
			// if err != nil {
			// 	panic(err)
			// }
			// fmt.Println(string(bytes))
			fullReport := models.ConsistencyReport{IssuesFound: result.IssuesFound, CheckResults: []models.CompleteCheckResult{result}, Checks: []string{"variable name casing"}, CodeBaseConsistencyScore: result.FinalConsistencyScore}
			// bytes, err := json.Marshal(fullReport)
			// if err != nil {
			// 	panic(err)
			// }
			// fmt.Println(string(bytes), "\nConsistency Check Score - - - > ", fullReport.CodeBaseConsistencyScore)
			fmt.Printf("\nConsistency Check Score - - - >  %d", fullReport.CodeBaseConsistencyScore)

		}()
	} else {
		println("unexpected command please try again")
	}
}
