package main

import (
	"archive/zip"
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	ignoreHelp := "relative or absolute path to a .zipignore file\nIf not provided, zipproject will look in the root directory of the project being zipped"
	ignoreFilePtr := flag.String("file", "./.zipignore", ignoreHelp)
	workDirPtr := flag.String("dir", "./", "directory to zip")
	outputPtr := flag.String("out", "project.zip", "filename to output the zip to.")

	flag.Parse()

	if *ignoreFilePtr == "./.zipignore" {
		*ignoreFilePtr = *workDirPtr + "/.zipignore"
	}

	ignoreList, err := getIgnoreList(*ignoreFilePtr)
	if err != nil {
		log.Fatal(err)
	}

	fileList, err := getFileList(*workDirPtr)
	if err != nil {
		log.Fatal(err)
	}

	zipList, err := pruneList(fileList, ignoreList)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nItems to zip:")
	for _, item := range zipList {
		fmt.Println(item)
	}
	fmt.Printf("\nOutputting to: %s\n", *outputPtr)

	err = zipFiles(*outputPtr, zipList, *workDirPtr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("All files successfully written to: %s\n", *outputPtr)
}

func zipFiles(zipName string, files []string, root string) error {
	// Create new zip file in memory
	zipFile, err := os.Create(zipName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// Create ZipWriter instance
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Add files to the archive
	for _, file := range files {
		if err = addFileToZip(zipWriter, file, root); err != nil {
			return err
		}
	}

	return nil
}

func addFileToZip(zipWriter *zip.Writer, filename string, root string) error {

	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get file info
	fileInfo, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	fileHeader, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}

	// Overwrite fileHeader.Name to preserve folder structure
	fileHeader.Name = strings.TrimPrefix(filename, root)

	// Set compression method to deflate to get better performance
	fileHeader.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(fileHeader)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, fileToZip)
	return err
}

//	getFileList gets a list of all files present in the working directory.
func getFileList(workDir string) ([]string, error) {
	fileList := make([]string, 0)

	err := filepath.Walk(workDir, func(path string, info os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	return fileList, err
}

// getIgnoreList reads the list of files to ignore from the provided location.
// Param 1: listLoc is the relative or absolute path to the ignore file.
func getIgnoreList(listLoc string) ([]string, error) {
	ignoreList := make([]string, 0)

	ignoreFile, err := os.Open(listLoc)
	if err != nil {
		return nil, err
	}
	defer ignoreFile.Close()

	scanner := bufio.NewScanner(ignoreFile)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			if line[0] != '#' {
				ignoreList = append(ignoreList, line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	ignoreList = append(ignoreList, listLoc)

	return ignoreList, nil
}

// pruneList removes all ignored files and folders from the original list.
// The returned list contains only files to zip.
func pruneList(fileList []string, ignoreList []string) ([]string, error) {
	pruneIdx := make([]int, 0)

	// TODO: Figure out how to handle ignored directories

	// Determine all the indicies of which entries to not zip
	for fileIdx, file := range fileList {
		fileInfo, err := os.Stat(file)
		if err != nil {
			return nil, err
		}

		if fileInfo.IsDir() {
			pruneIdx = append(pruneIdx, fileIdx)
			continue
		}

		for _, ignoreItem := range ignoreList {
			// Handle directories
			dirName := strings.TrimRight(ignoreItem, "/")
			if strings.Contains(file, ignoreItem) || (strings.Contains(file, dirName) && fileInfo.IsDir()) {
				pruneIdx = append(pruneIdx, fileIdx)
				continue
			}

			// Handle wildcard instances for file extensions
			if strings.HasPrefix(ignoreItem, "*") {
				suffix := strings.TrimLeft(ignoreItem, "*")
				if strings.HasSuffix(file, suffix) {
					pruneIdx = append(pruneIdx, fileIdx)
					continue
				}
			}

			// Handle standard files
			if ignoreItem == fileInfo.Name() {
				pruneIdx = append(pruneIdx, fileIdx)
				continue
			}
		}
	}

	// Create a list consisting of all files that will be zipped.
	filesToZip := make([]string, 0)
	for idx, file := range fileList {
		if !containsValue(pruneIdx, idx) {
			filesToZip = append(filesToZip, file)
		}
	}
	return filesToZip, nil
}

// Checks if an int slice contains a value
func containsValue(numSlice []int, value int) bool {
	for _, data := range numSlice {
		if data == value {
			return true
		}
	}

	return false
}
