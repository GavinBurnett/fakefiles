package main

const (
	UI_Arguments                       = `Arguments: `
	UI_InvalidArgs                     = `ERROR: Invalid arguments`
	UI_ReadingConfigFile               = `Reading config file: `
	UI_FileFound                       = `File found: `
	UI_FileDirFound                    = `File/Directory found: `
	UI_FileDirFoundError               = `File/Directory check error: `
	UI_ReadConfigFileError             = `ERROR: Reading config file: `
	UI_FileNotFound                    = `ERROR: File not found`
	UI_FileDirNotFound                 = `ERROR: File/Directory not found: `
	UI_DirNotFound                     = `ERROR: Directory not found: `
	UI_SkippingLine                    = ` Skipping Line: `
	UI_ConfigFileReadError             = `ERROR: Config file read error - line: `
	UI_CreateDirError                  = `ERROR: Directory create error: `
	UI_CreateDirFilesError             = `ERROR: Directory/files create error: `
	UI_NoDirGiven                      = `ERROR: No directory specified`
	UI_NoParametersGiven               = `ERROR: No parameters specified`
	UI_CreatingFile                    = `Creating file: `
	UI_ParameterInvalid                = `ERROR: Invalid parameter: %s , Parameters: `
	UI_RandomDataError                 = `ERROR: Random data not generated`
	UI_FileCreateError                 = `ERROR: Failed to create file: `
	UI_NumberOfFilesToCreate           = `Number of files to create: `
	UI_FileSize                        = `File size: `
	UI_FilePath                        = `File name and path: `
	UI_FilePathError                   = `ERROR: File name and path invalid.`
	UI_FileError                       = `ERROR: File name invalid.`
	UI_InvalidFileSizeError            = `ERROR: Invalid file size: `
	UI_FileNamesCountMismatchError     = `ERROR: File names count %s does not match file extension count %s .`
	UI_NoFileNameOrFileExtensionsError = `ERROR: No file names or file extensions.`
	UI_InvalidNumberOfFilesError       = `ERROR: Invalid number of files.`
	UI_ProcessedStringList             = `Processed string list: `
	UI_Done                            = ` - Done.`
	UI_ParseError                      = `ERROR: Parse failed: `
	UI_CreatingInDir                   = `Creating files in directory: `
	UI_ConfigFileData                  = `Config file data: `
	UI_ConfigFileEntry                 = `Config file entry: `
	UI_CreatedDirectory                = `Created directory: `
	UI_ConfigFileNotFound              = `ERROR: Config file not found.`

	UI_Help = `fakefiles v1.0 by gburnett@outlook.com

Creates files filled with random data.

Arguments: 
	fakefiles directory FakeFilesConfig.txt

	directory - directory to create files in
	FakeFilesConfig.txt - config file setting number and type of files to create`
)
