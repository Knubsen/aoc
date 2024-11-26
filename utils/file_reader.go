package utils

import (
	"bufio"
	"os"
	"path/filepath"
)

func ReadLines(path string) ([]string, error) {
	path = getImportFilePath(path)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, scanner.Err()
}

func getImportFilePath(day string) string {
	cwd, err := os.Getwd()
	Check(err)

	importPath := filepath.Join(cwd, "days", day, "input/input.txt")
	return importPath
}
