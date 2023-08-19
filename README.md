# CodeAudit - CLI Tool for verifying syntax consistency of your codebases

This simple terminal tool allows you to check for syntax consistency around various items such as naming, spacing, semi-colons, etc. I wrote this tool because I realized a lot of teams
focus on things like optimization, whether to use a simple for-loop or array methods, and strict typing to solve and find bugs. I have found that sometimes bugs can come from simple things such as miss naming 
variable names, indentation issues (Python or callback hell situations), and just lack of consistency throughout codebases. This testing tool will allow you to set certain syntax preferences, dictate which project to search, and generate reports that show which files and lines failed checks, which checks were ran, and provide overall scores for the validity of your codebase.


Currently only works with Javascript files, but other languages will be available. Also only name case checking works at the moment. 

### Checks Available
|Check|Description|
|----|-----|
|Naming Convention| Checks variables and functions for correct casing (e.g. snake_case)|
|Indentation| Checks to indentation spacing for desire amount (e.g. 2 or 4)|
|Line Character Count| Checks each line in each found file to make sure character count doesn't exceed specified limit (80-100 is most readible)|

### Checks In Development:

- Spacing (between words, spaces between function definitions and variables)
- Semi-Colons

### Languages Coming Soon:

- Python
- TypeScript
- Go

## Startup Instructions  

### Basic Example:

```shell
go build
./codeAudit runchecks -j

// To pick desired output folder name

go build -o [DESIRED_NAME_OF_EXECUTABLE]
./[DESIRED_NAME_OF_EXECUTABLE] runchecks -j
```

This example will run naming convention checks on the current working directory and display the overall score for your cwd's javascript files. It will also generate a full report in your download folder.


_<small> <span>__Cool Fact:__</span> You can also install package globally by setting the PATH variable in terminal to allow using this tool anywhere. Refer to this Golang [Documentation]('https://go.dev/doc/tutorial/compile-install) on how to do so</small>_

\* _Must have Go install on local machine_

## Command Line Syntax

```shell
./codeAudit [check_type] -[flag] [path] [config_initial]=[config_value]
```

### BreakDown

| Command | Purpose | values |
|---------|---------|--------|
| Check Type | The type of check you would like to make  |  all, indent, naming, char. Default: all|
| Flag | Dictates type file extension to query for | j (javascript), p (python), t (typescript), g (golang). Default: j |
| Path | Path to desired desired directory to recursively search through for files| default ./ (cwd)|
|Config Initial| Initials determines configure option to set| i (indentation - integer), n (naming convention **), c (character count limit - integer), and r (generate report boolean) |
|Config Value| Value to assign to desired config option| c and i are integers only, see ** below for n|

<br>
\** Naming Conventions available - pascal (PascalCase), snake (snake_case), and camel (camelCase)


## Reports

This tool generates a simple failure report in the form of a a csv file called `codeAuditReport.csv`. This file after running the commands above will be found downloads folder in file system. It will list all files that had a failed validation checks and helpful information (like line number where errored), to help with debug and addressing syntax errors.

__Format__ : File Path, Line Number (where error occurred), Check Made (that failed), and issue (message clarifying error).


### Feedback Encouraged!

Please create any issues on this repo if you experience any bugs or have a feature suggestion. Looking to really grow this tool into something truly helpful for developers.