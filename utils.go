package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// Reads the go.mod file to extract the module name (package name)
func getPackageName(dir string) (string, error) {
	goModPath := filepath.Join(dir, "go.mod")

	data, err := os.ReadFile(goModPath)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "module") {
			// Extract the module name (package name)
			parts := strings.Fields(line)
			if len(parts) > 1 {
				return parts[1], nil
			}
		}
	}

	return "", fmt.Errorf("module name not found in go.mod")
}

// This function Add new line to the content at specfic index
func AddNewLineToStart(newline, content, startKeyword, endKeyword, addToStartOfLine, addToEndOfLine string) string {
	if strings.Contains(normalizeWhitespace(content), normalizeWhitespace(newline)) {
		return content // Line already exists, no modification needed
	}

	// Find the type Repository struct block and locate the closing bracket
	if idx := strings.Index(content, startKeyword); idx != -1 {
		// Find the closing curly bracket for the struct
		start := idx + len(startKeyword)
		// end := strings.Index(content[start:], endKeyword)

		content = content[:start] + "\n" + addToStartOfLine + newline + addToEndOfLine + "\n" + content[start+1:]
	}

	return content
}

func AddNewLineToEnd(newline, content, startKeyword, endKeyword, addToStartOfLine, addToEndOfLine string) string {
	if strings.Contains(normalizeWhitespace(content), normalizeWhitespace(newline)) {
		return content // Line already exists, no modification needed
	}

	// Find the type Repository struct block and locate the closing bracket
	if idx := strings.Index(content, startKeyword); idx != -1 {
		// Find the closing curly bracket for the struct
		start := idx + len(startKeyword)
		end := strings.Index(content[start:], endKeyword)
		if end != -1 {
			// Insert the new repository line before the closing curly bracket
			insertPos := start + end
			content = content[:insertPos] + addToStartOfLine + newline + addToEndOfLine + content[insertPos:]
		}
	}

	return content
}

// ToUpperFirst converts the first letter of a string to uppercase
func ToUpperFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

// ToLowerFirst converts the first letter of a string to uppercase
func ToLowerFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(string(s[0])) + s[1:]
}

// ReplaceAll template entity to given entity name
func replaceEntityName(content, entityName string) string {
	entityNameUpperFirst := ToUpperFirst(entityName)
	content = strings.ReplaceAll(content, "TemplateEntity", entityNameUpperFirst)
	content = strings.ReplaceAll(content, "templateEntity", entityName)
	content = strings.ReplaceAll(content, "templateEntity", entityName)
	content = strings.ReplaceAll(content, "ms-name", msName)
	content = strings.ReplaceAll(content, "EPOCH", GetEpoch())

	return content
}

func replaceFileName(pathName string, entityName string) string {
	entity_name := CamelToSnake(entityName)
	pathName = strings.ReplaceAll(pathName, "template_entity", entity_name)
	pathName = strings.ReplaceAll(pathName, "EPOCH", GetEpoch())
	return pathName
}

// CamelToSnake converts a CamelCase string to snake_case
func CamelToSnake(input string) string {
	// Insert an underscore before any uppercase letter followed by a lowercase letter or a digit
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	snake := re.ReplaceAllString(input, `${1}_${2}`)

	// Convert the entire string to lowercase
	return strings.ToLower(snake)
}

// Write file content in  given file path
func WriteFileInPath(filePath, content string) error {
	// Write the modified content back to the file
	var err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing to file %s: %v", filePath, err)
	}

	return nil
}

func GetEpoch() string {
	// Get the current time
	currentTime := time.Now()

	// Format the time as YYYYMMDDHHMMSS
	formattedTime := currentTime.Format("20060102150405")

	return formattedTime
}

func normalizeWhitespace(input string) string {
	// Replace multiple spaces with a single space
	return strings.Join(strings.Fields(input), " ")
}
