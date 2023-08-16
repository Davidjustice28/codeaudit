# Code Audit CLI TOOL 

Golang base cli tool that runs checks on your codebase for given file extensions. Checks if files meet your set syntax requirements and then creates a report for you in your current working directory

Currently testing with Javascript files, but will also eventually work with other programming language files (python, golang, typescript, etc).

Report generated will give a percentage (%) of files that pass each check as well as a total score (%) for files that pass checks.
Files that fail any checks will also be listed with the given lines to easy find the errors.

Currently only checks variable and function casing, but will also be checking for semi-colons, spacing, indentation, etc.