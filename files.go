package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// CreateDirsAndFiles: Creates files and directories stored in the given config in the given directory
func CreateDirsAndFiles(_startDir string, _config []Config) bool {

	var filesCreatedOk bool = true

	if len(_startDir) > 0 && len(_config) > 0 {

		if DEBUG == true {
			fmt.Println(UI_CreatingInDir, _startDir)
			fmt.Println(UI_ConfigFileData, _config)
		}

		// Loop through every config entry and generate files and directories
		for _, currentConfig := range _config {

			if DEBUG == true {
				fmt.Println(UI_ConfigFileEntry, currentConfig)
			}

			if FileDirectoryExists(_startDir) {

				if currentConfig.DirectoryOnly == true {

					// If current config just creates a directory, create a directory
					if CreateDirectory(_startDir+"/"+currentConfig.Directory) == false {

						// Failed to create directory
						filesCreatedOk = false
						fmt.Println(UI_CreateDirError, _startDir+"/"+currentConfig.Directory)
						break
					}

				} else {

					// If current config just creates directories and file, create directories and files
					if CreateFiles(_startDir, currentConfig) == false {

						// Failed to create directories and files
						filesCreatedOk = false
						fmt.Println(fmt.Sprintf(UI_CreateDirFilesError, "_startDir: ", _startDir, "_currentConfig: ", currentConfig))
						break

					}
				}

			} else {

				// No start directory
				filesCreatedOk = false
				fmt.Println(UI_DirNotFound, _startDir)
				break
			}

		} // End for

	} else {
		filesCreatedOk = false
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_startDir:", _startDir, "_config:", _config)
	}

	return filesCreatedOk

}

// FileDirectoryExists: Check the given file/directory exists
func FileDirectoryExists(_fileDirectory string) bool {

	fileDirectoryExists := false

	if len(_fileDirectory) > 0 {

		// Try to get file info
		_, err := os.Stat(_fileDirectory)

		// If any errors occur on getting file info, file does not exist
		if os.IsNotExist(err) || err != nil {
			if DEBUG == true {
				fmt.Println(UI_FileDirNotFound + _fileDirectory)
			}
			fileDirectoryExists = false
		} else {
			// File info found - file exists
			if DEBUG == true {
				fmt.Println(UI_FileDirFound + _fileDirectory)
			}
			fileDirectoryExists = true
		}
	} else {
		fmt.Println(UI_FileDirFoundError, _fileDirectory, GetFunctionName())
	}

	return fileDirectoryExists
}

// CreateFiles: Create the files in the given config in the given directory
func CreateFiles(_startDir string, _config Config) bool {

	var filesCreated bool = false
	var numberOfFiles int64 = 0
	var counter int64 = 0
	var fileNamesCounter int = 0
	var fileName string
	var fileSizeMB int64 = 0
	var fullFilePath string
	var fileSizeBytes int64 = 0
	var fileCreated bool = false

	if len(_startDir) > 0 && len(_config.Directory) > 0 {

		if FileDirectoryExists(_startDir+"/"+_config.Directory) == false {

			// Create directory if it does not already exist
			if CreateDirectory(_startDir+"/"+_config.Directory) == false {
				fmt.Println(UI_CreateDirError, _startDir+"/"+_config.Directory)
			}

		} else {
			// Directory already exists
			if DEBUG == true {
				fmt.Println(UI_FileDirFound + _startDir + "/" + _config.Directory)
			}
		}

		// For the config number of files range - get a random number in this range
		numberOfFiles = GetRandomNumber(_config.NumberOfFiles.Lowest, _config.NumberOfFiles.Highest)

		if DEBUG == true {
			fmt.Println(UI_NumberOfFilesToCreate + strconv.Itoa(int(numberOfFiles)))
		}

		if numberOfFiles > 0 && numberOfFiles >= _config.NumberOfFiles.Lowest && numberOfFiles <= _config.NumberOfFiles.Highest {

			// For the random number of files
			for counter = 0; counter != numberOfFiles; counter++ {

				if len(_config.FileNames) > 0 && len(_config.FileExtensions) > 0 {

					if len(_config.FileNames) == len(_config.FileExtensions) {

						// For the config number of file names and file name extensions
						for fileNamesCounter = 0; fileNamesCounter != len(_config.FileNames); fileNamesCounter++ {

							// For the config file size range, get a random number in this range
							fileSizeMB = GetRandomNumber(_config.FileSizes.Lowest, _config.FileSizes.Highest)

							if DEBUG == true {
								fmt.Println(UI_FileSize + strconv.Itoa(int(fileSizeMB)))
							}

							if fileSizeMB > 0 && fileSizeMB >= _config.FileSizes.Lowest && fileSizeMB <= _config.FileSizes.Highest {

								// Build the file name using the config file name and file extension
								fileName = _config.FileNames[fileNamesCounter] + "-" + strconv.Itoa(int(counter)) + "." + _config.FileExtensions[fileNamesCounter]

								if len(fileName) > 0 {

									// Build the full directory path using the working directory, config directory and file name
									fullFilePath = _startDir + "/" + _config.Directory + "/" + fileName

									if len(fullFilePath) > 0 {

										if DEBUG == true {
											fmt.Println(UI_FilePath + fullFilePath)
										}

										// Get file size in bytes
										fileSizeBytes = ConvertMegabytesToBytes(fileSizeMB)

										if fileSizeBytes != -1 && fileSizeBytes > 0 {

											// Add a random amount to file size
											fileSizeBytes += fileSizeBytes + GetRandomNumber(0, MB_IN_BYTES)

											if fileSizeBytes != -1 && fileSizeBytes > 0 {

												// Create the file
												fileCreated = CreateFile(fullFilePath, fileSizeBytes)

												if fileCreated == true {
													// If one file created, set files created to true
													filesCreated = true
												} else {
													fmt.Println(UI_FileCreateError + fullFilePath)
												}

											} else {
												fmt.Println(UI_InvalidFileSizeError + strconv.Itoa(int(fileSizeBytes)))
											}

										} else {
											fmt.Println(UI_InvalidFileSizeError + strconv.Itoa(int(fileSizeBytes)))
										}

									} else {
										fmt.Println(UI_FilePathError)
									}

								} else {
									fmt.Println(UI_FileError)
								}

							} else {
								fmt.Println(UI_InvalidFileSizeError + strconv.Itoa(int(fileSizeMB)))
							}

						} // End file names loop

					} else {
						fmt.Println(fmt.Sprintf(UI_FileNamesCountMismatchError, strconv.Itoa(len(_config.FileNames)), strconv.Itoa(len(_config.FileExtensions))))
					}

				} else {
					fmt.Println(UI_NoFileNameOrFileExtensionsError)
				}

			} // End file number loop

		} else {
			fmt.Println(UI_InvalidNumberOfFilesError)
		}

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_startDir:", _startDir, "_config:", _config)
	}

	return filesCreated
}

// CreateFile: Creates a file with the given file name and file size
func CreateFile(_fileName string, _fileSize int64) bool {

	var fileCreated bool = true
	var randomData []byte
	var bytesWritten int64
	var bytesToWrite int64
	var writeBuffer *bufio.Writer

	if len(_fileName) > 0 && _fileSize > 0 {

		fmt.Print(UI_CreatingFile + _fileName)

		// Create the file
		file, err := os.Create(_fileName)
		if err == nil {

			// Fill file with random data
			bytesToWrite = GetRandomNumber(0, MB_IN_BYTES)

			if bytesToWrite > 0 {

				writeBuffer = bufio.NewWriter(file)
				writeBuffer = bufio.NewWriterSize(writeBuffer, int(bytesToWrite))

				for bytesWritten = 0; bytesWritten < _fileSize; bytesWritten += bytesToWrite {

					randomData = GetRandomData(bytesToWrite)

					if randomData != nil && int64(len(randomData)) == bytesToWrite {

						bytesWritten, err := writeBuffer.Write(randomData)

						if err != nil || bytesWritten != len(randomData) {
							fileCreated = false
							break
						}

						randomData = nil
					}
				} // end for

				writeBuffer.Flush()
				writeBuffer = nil

				if fileCreated == false {
					fmt.Println(UI_FileCreateError, _fileName, _fileSize, err.Error())

				} else {
					// File created
					fmt.Println(UI_Done)
				}
			}

		} else {
			if DEBUG == true {
				fmt.Println(UI_FileCreateError, _fileName, _fileSize)
			}
			fileCreated = false
		}

		file.Sync()
		file.Close()

	} else {
		fileCreated = false
		fmt.Println(UI_FileCreateError, _fileName, _fileSize, GetFunctionName())
	}

	return fileCreated
}

// CreateDirectory: Creates a directory with the given path
func CreateDirectory(_dirName string) bool {

	var dirCreated bool = true

	if len(_dirName) > 0 {

		err := os.Mkdir(_dirName, 0755)
		if err != nil {
			dirCreated = false
			fmt.Println(UI_CreateDirError, _dirName, err.Error())
		} else {
			if DEBUG == true {
				fmt.Println(UI_CreatedDirectory + _dirName)
			}
		}

	} else {
		dirCreated = false
		fmt.Println(UI_CreateDirError, _dirName, GetFunctionName())
	}

	return dirCreated
}
