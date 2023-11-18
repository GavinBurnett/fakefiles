package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config file data
type Config struct {
	Directory      string      // Directory Name
	DirectoryOnly  bool        // Flag to create a directory with no files
	NumberOfFiles  NumberRange // Number of files to generate - number is random value between set lower and upper values
	FileSizes      NumberRange // Size of files to generate - size is random value between set lower and upper values
	FileNames      []string    // Names of files to generate
	FileExtensions []string    // Extensions of files to generate
}

// Number range data
type NumberRange struct {
	Lowest  int64
	Highest int64
}

// Config file strings
const LINE_COMMENT = "#"
const LINE_SPLIT = ":"
const LINE_CONFIG_ENTRY = "="
const DIR_CONFIG = "Dir"
const FILES_CONFIG = "Files"
const SIZES_CONFIG = "Sizes"
const NAMES_CONFIG = "Names"
const EXTENSIONS_CONFIG = "Extensions"

const RANGE = "-"
const SEPARATOR = ","

// ReadConfigFile: Read config file in
func ReadConfigFile(_configFile string) []Config {

	var configData []Config = nil

	if DEBUG == true {
		fmt.Println(UI_ReadingConfigFile, _configFile)
	}

	if len(_configFile) > 0 {

		if FileDirectoryExists(_configFile) {

			configFile, err := os.Open(_configFile)

			if err == nil {

				if DEBUG == true {
					fmt.Println(UI_FileFound + _configFile)
				}

				// Read every line in the config file and store the data in Config struct
				configFileReader := bufio.NewScanner(configFile)

				for configFileReader.Scan() {

					if len(configFileReader.Text()) > 0 {

						// Process current line in config file

						if strings.HasPrefix(configFileReader.Text(), LINE_COMMENT) {
							// Do not process any commented out lines
							if DEBUG == true {
								fmt.Println(_configFile + UI_SkippingLine + configFileReader.Text())
							}
						} else {

							var configLine Config

							if strings.Contains(configFileReader.Text(), LINE_SPLIT) {

								// Current line has directory and file data
								configLine.DirectoryOnly = false

								currentLine := strings.Split(configFileReader.Text(), LINE_SPLIT)

								for _, currentLineBlock := range currentLine {

									if (len(currentLineBlock)) > 0 {

										// Get directory name
										if strings.Contains(currentLineBlock, DIR_CONFIG+LINE_CONFIG_ENTRY) {
											configLine.Directory = GetDirectory(currentLineBlock)
										}

										// Get number of files range
										if strings.Contains(currentLineBlock, FILES_CONFIG+LINE_CONFIG_ENTRY) {
											configLine.NumberOfFiles.Lowest, configLine.NumberOfFiles.Highest = GetNumberRange(currentLineBlock)
										}

										// Get file sizes range
										if strings.Contains(currentLineBlock, SIZES_CONFIG+LINE_CONFIG_ENTRY) {
											configLine.FileSizes.Lowest, configLine.FileSizes.Highest = GetNumberRange(currentLineBlock)
										}

										// Get list of file names
										if strings.Contains(currentLineBlock, NAMES_CONFIG+LINE_CONFIG_ENTRY) {
											configLine.FileNames = GetStringList(currentLineBlock)
										}

										// Get list of extensions
										if strings.Contains(currentLineBlock, EXTENSIONS_CONFIG+LINE_CONFIG_ENTRY) {
											configLine.FileExtensions = GetStringList(currentLineBlock)
										}

									}

								} // end for

							} else {

								// Current line has directory only
								configLine.DirectoryOnly = true
								configLine.Directory = GetDirectory(configFileReader.Text())
							}

							if DEBUG == true {
								fmt.Println(configLine)
							}

							// Add current line data to config data
							configData = append(configData, configLine)

						}
					} else {
						// Line is empty
					}

				} // end for

			} else {
				fmt.Println(UI_ReadConfigFileError, GetFunctionName(), _configFile, err)
			}

			configFile.Close()

		} else {
			fmt.Println(UI_ReadConfigFileError, GetFunctionName(), _configFile, UI_FileNotFound)
		}

	} else {
		fmt.Println(UI_ReadConfigFileError, GetFunctionName(), _configFile, UI_FileNotFound)
	}

	return configData
}

// GetDirectory: Gets the directory path from the given string
func GetDirectory(_directoryString string) string {

	var directory string = ""

	if len(_directoryString) > 0 {

		directory = strings.Split(_directoryString, LINE_CONFIG_ENTRY)[1]

		if len(directory) > 0 {
			if DEBUG == true {
				fmt.Println(UI_FileDirFound, _directoryString)
			}
		} else {
			fmt.Println(UI_FileDirNotFound, _directoryString)
		}

	} else {
		fmt.Println(UI_FileDirNotFound, _directoryString)
	}

	return directory
}

// GetNumberRange: Gets the lowest and highest values from the passed in string
func GetNumberRange(_numberOfFilesString string) (int64, int64) {

	var lowest int64 = 0
	var highest int64 = 0

	if len(_numberOfFilesString) > 0 {

		parseRange := strings.Split(_numberOfFilesString, LINE_CONFIG_ENTRY)[1]

		if len(parseRange) > 0 {

			lowAndHigh := strings.Split(parseRange, RANGE)

			if len(lowAndHigh) > 1 {

				lowest = ParseStringToInt64(lowAndHigh[0])
				highest = ParseStringToInt64(lowAndHigh[1])

			} else {
				fmt.Println(UI_ConfigFileReadError, GetFunctionName(), _numberOfFilesString)
			}

		} else {
			fmt.Println(UI_ConfigFileReadError, GetFunctionName(), _numberOfFilesString)
		}

	} else {
		fmt.Println(UI_ConfigFileReadError, GetFunctionName(), _numberOfFilesString)
	}

	return lowest, highest
}

// GetStringList: Takes comma delimited string list and returns all entries in a string array
func GetStringList(_stringList string) []string {

	var stringList []string

	if len(_stringList) > 0 {

		parseStringList := strings.Split(_stringList, LINE_CONFIG_ENTRY)[1]

		if len(parseStringList) > 0 {

			stringList = strings.Split(parseStringList, SEPARATOR)

			if len(stringList) > 0 {
				if DEBUG == true {
					fmt.Println(UI_ProcessedStringList, _stringList)
				}
			} else {
				fmt.Println(UI_ConfigFileReadError, GetFunctionName(), _stringList)
			}

		} else {
			fmt.Println(UI_ConfigFileReadError, GetFunctionName(), _stringList)
		}

	} else {
		fmt.Println(UI_ConfigFileReadError, GetFunctionName(), _stringList)
	}

	return stringList
}

// ParseStringToInt64: Converts given string into an int64
func ParseStringToInt64(_inputString string) int64 {

	var outputInt64 int64 = -1

	if len(_inputString) > 0 {

		parsedInt, parseErr := strconv.ParseInt(_inputString, 10, 64)

		if parseErr == nil {
			outputInt64 = parsedInt
		} else {
			fmt.Println(UI_ParseError, _inputString)
		}

	}

	return outputInt64
}
