package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"codeAudit/checks"
	"codeAudit/models"
	"codeAudit/utils"
)

//TODO: character count checks

func main() {
	var fileType string
	var check string
	root_directory := "./"
	defaultCharacterLimit := 100
	defaultVariableNamingConvention := "camel"
	defaultIndentationSpaces := 2
	// root_directory := "/Users/david/Programming/projects/practice-projects/git-practice-project/"

	terminalArgs := os.Args[1:]
	if len(terminalArgs) == 0 {
		fileType = ".js"
		check = "all"
	}

	if len(terminalArgs) >= 1 {
		// println("action", terminalArgs[0])
		if !utils.IsConfigOption(terminalArgs[0]) {
			check = terminalArgs[0]
		}
		fileType = ".js"
	}

	if len(terminalArgs) >= 2 {
		// println("action", terminalArgs[0], "flag", terminalArgs[1])

		if !utils.IsConfigOption(terminalArgs[1]) {
			if !utils.IsConfigOption(terminalArgs[0]) {
				check = terminalArgs[0]
			}
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
	}

	if len(terminalArgs) >= 3 {
		if !utils.IsConfigOption(terminalArgs[2]) {
			root_directory = terminalArgs[2]
			fmt.Printf("path found in terminal %s\n", root_directory)
		}
	}

	//TODO: Make it so users can do i=2, c=40, n=camel

	println("terminal arguments received")
	for a := 0; a < len(terminalArgs); a++ {
		arg := terminalArgs[a]
		if utils.IsConfigOption(arg) {
			configTuple := strings.Split(arg, "=")
			option := configTuple[0]
			value := configTuple[1]

			switch option {
			case "i":
				newIndentation, err := strconv.Atoi(value)
				if err != nil {
					fmt.Printf("Input %s wasn't able to be converted to a number to use as new number of indentation. Used default value of 2\n", value)
				} else {
					defaultIndentationSpaces = newIndentation
				}

			case "c":
				newCharacterLineLimit, err := strconv.Atoi(value)
				if err != nil {
					fmt.Printf("Input %s wasn't able to be converted to a number to use as new character limit. Used default value of 100\n", value)
				} else {
					defaultCharacterLimit = newCharacterLineLimit
				}
			case "n":
				isValidConventions := utils.IsValidNamingConvention(value)
				if !isValidConventions {
					fmt.Printf("%s isn't a valid or supported variable naming convention. See docs\n", value)
				} else {
					defaultVariableNamingConvention = value
				}
			}
		}
	}

	var report models.ConsistencyReport
	checksRan := true

	switch check {
	case "all":
		report = checks.PerformAllChecks(root_directory, fileType, defaultVariableNamingConvention, defaultIndentationSpaces, defaultCharacterLimit)
	case "indent":
		report = checks.PerformIndentationChecks(root_directory, fileType, defaultIndentationSpaces)
	case "naming":
		report = checks.PerformVariableNamingChecks(root_directory, fileType, defaultVariableNamingConvention)
	case "char":
		report = checks.PerformCharacterCountChecks(root_directory, fileType, defaultCharacterLimit)
	default:
		checksRan = false
		println("unexpected command please try again\n")
	}

	if checksRan {
		fmt.Printf("Number of issues found: %d\n", len(report.IssuesFound))
	}
}
