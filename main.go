// fakefiles project main.go
package main

import (
	"fmt"
	"os"
)

func main() {

	var configData []Config = nil
	var configOK = false
	var rootDirOK = false
	var filesCreatedOk = false
	exitCode := 0

	if DEBUG == true {
		fmt.Println(UI_Arguments, os.Args)
	}

	if os.Args != nil {

		if len(os.Args) == 1 {
			// No user arguments given - display help
			fmt.Println(UI_Help)
		}
		if len(os.Args) == 2 {
			if IsStringHelpArgument(os.Args[1]) {
				// User has given help argument - display help
				fmt.Println(UI_Help)
			} else {
				// User has given only one argument that is not a help argument - display error
				exitCode = -1
				fmt.Println(UI_InvalidArgs)
			}
		}
		if len(os.Args) == 3 {

			if DEBUG == true {
				fmt.Println(os.Args)
			}

			// Check root dir exists
			if FileDirectoryExists(os.Args[1]) == true {
				rootDirOK = true
			} else {
				// Root dir does not exist
				fmt.Println(UI_DirNotFound, os.Args[1])
			}

			// Read config file
			configData = ReadConfigFile(os.Args[2])

			if configData != nil && len(configData) > 0 {

				if DEBUG == true {
					for _, currentConfigLine := range configData {
						fmt.Println(currentConfigLine)
					}
				}

				configOK = true

			} else {
				// Config file not found
				fmt.Println(UI_ConfigFileNotFound, os.Args[2])
			}

			// If config read in and root directory exists, create the files
			if configOK == true && rootDirOK == true {
				filesCreatedOk = CreateDirsAndFiles(os.Args[1], configData)
				if filesCreatedOk == false {
					exitCode = -1
				}
			} else {
				exitCode = -1
			}

		}
		if len(os.Args) > 3 {
			// Too many arguments - display error
			exitCode = -1
			fmt.Println(UI_InvalidArgs)
		}

	} else {
		// No arguments
		exitCode = -1
		fmt.Println(UI_NoParametersGiven)
	}

	os.Exit(exitCode)
}

// IsStringHelpArgument: Returns true if given string is a help argument, false if it is not
func IsStringHelpArgument(_theString string) bool {

	isHelpArgument := false

	if len(_theString) > 0 {

		switch _theString {
		case "?":
			isHelpArgument = true
		case "/?":
			isHelpArgument = true
		case "-?":
			isHelpArgument = true
		case "--?":
			isHelpArgument = true
		case "h":
			isHelpArgument = true
		case "/h":
			isHelpArgument = true
		case "-h":
			isHelpArgument = true
		case "--h":
			isHelpArgument = true
		case "help":
			isHelpArgument = true
		case "/help":
			isHelpArgument = true
		case "-help":
			isHelpArgument = true
		case "--help":
			isHelpArgument = true
		}

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_theString:", _theString)
	}

	return isHelpArgument
}
