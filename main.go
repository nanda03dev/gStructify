package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//go:embed clean-template/*
var cleanTemplate embed.FS

func main() {
	// Get current directory (where the executable is run)
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting working directory: %v", err)
		return
	}

	// Read the go.mod file to get the package name
	packageName, err := getPackageName(wd)
	if err != nil {
		fmt.Printf("Error reading go.mod: %v", err)
		return
	}

	// Accept the package name as a command-line argument
	entity := flag.String("entity", "", "Name of the package (e.g., book)")
	flag.Parse()

	var entityName = *entity

	if entityName == "" {
		fmt.Println("Entity name not specified, Please specify entity name by using this flag -entity")
		return
	}

	fmt.Println("Created layers for entity : ", entityName, " successfully!")

	// Generate the microservice using the package name
	CreateNewMS(wd, packageName, ToLowerFirst(entityName))

	// This will import all required local packages
	ImportAllPacakges(wd)

	// Run `go mod tidy` in the working directory
	err = RunGoModTidy(wd)
	if err != nil {
		fmt.Printf("Error running 'go mod tidy': %v\n", err)
	}

}

// CreateNewMS creates the microservice by copying and modifying the template files
func CreateNewMS(outputDir, packageName string, entityName string) {
	// Copy and modify template files
	err := copyDirAndModify(cleanTemplate, "clean-template", outputDir, packageName, entityName)
	if err != nil {
		panic(err)
	}
}

// copyDirAndModify copies template files and modifies them with the package name
func copyDirAndModify(efs embed.FS, srcDir, destDir, packageName string, entityName string) error {
	entries, err := efs.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(srcDir, entry.Name())
		destPath := filepath.Join(destDir, entry.Name())

		if entry.IsDir() {
			err = os.MkdirAll(destPath, 0755)
			if err != nil {
				return err
			}
			err = copyDirAndModify(efs, srcPath, destPath, packageName, entityName)
			if err != nil {
				return err
			}
		} else {
			data, err := efs.ReadFile(srcPath)
			if err != nil {
				return err
			}

			// Check if destination file exists
			if _, err := os.Stat(destPath); err == nil {
				// File exists, modify it
				err = modifyFile(destPath, entityName)
				if err != nil {
					return err
				}
			} else {

				// Rename file if it contains "template_entity"
				if strings.Contains(entry.Name(), "template_entity") {
					if entityName == "" {
						continue
					}
					destPath = strings.ReplaceAll(destPath, "template_entity", entityName)
				}

				// Replace content inside the file
				content := string(data)
				content = strings.ReplaceAll(content, "TemplateEntity", ToUpperFirst(entityName))
				content = strings.ReplaceAll(content, "templateEntity", entityName)
				content = strings.ReplaceAll(content, "github.com/nanda03dev/go-ms-template", packageName)

				err = os.WriteFile(destPath, []byte(content), 0644)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// modifyFile modifies the destination file content to include additional code
func modifyFile(filePath, entityName string) error {
	if strings.Contains(filePath, "app_module.go") {
		return ToUpdateAppModuleFile(filePath, entityName)
	}
	return nil
}

func ToUpdateAppModuleFile(filePath, entityName string) error {
	// Read the existing file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", filePath, err)
	}

	content := string(data)
	entityNameUpperFirst := ToUpperFirst(entityName)

	lineToAddInRepoType := fmt.Sprintf("%sRepository aggregates.%sRepository", entityNameUpperFirst, entityNameUpperFirst)
	repoStartKeyword := "type Repository struct {"
	endKeyword := "}"
	content = AddNewLineToExistContent(lineToAddInRepoType, content, repoStartKeyword, endKeyword, "", "")

	lineToAddInRepositories := fmt.Sprintf("%sRepository: repositories.New%sRepository(databases)", entityNameUpperFirst, entityNameUpperFirst)
	lineToAddInRepositoriesStartKeyword := "var AllRepositories = Repository{"
	content = AddNewLineToExistContent(lineToAddInRepositories, content, lineToAddInRepositoriesStartKeyword, endKeyword, "", ",\n")

	lineToAddInServiceType := fmt.Sprintf("%sService services.%sService", entityNameUpperFirst, entityNameUpperFirst)
	serviceStartKeyword := "type Service struct {"
	content = AddNewLineToExistContent(lineToAddInServiceType, content, serviceStartKeyword, endKeyword, "", "")

	lineToAddInServices := fmt.Sprintf("%sService: services.New%sService(AllRepositories.%sRepository)", entityNameUpperFirst, entityNameUpperFirst, entityNameUpperFirst)
	lineToAddInServicesStartKeyword := "var AllServices = Service{"
	content = AddNewLineToExistContent(lineToAddInServices, content, lineToAddInServicesStartKeyword, endKeyword, "", ",\n")

	lineToAddInHandlerType := fmt.Sprintf("%sHandler handlers.%sHandler", entityNameUpperFirst, entityNameUpperFirst)
	handlerStartKeyword := "type Handler struct {"
	content = AddNewLineToExistContent(lineToAddInHandlerType, content, handlerStartKeyword, endKeyword, "", "")

	lineToAddInHanlders := fmt.Sprintf("%sHandler: handlers.New%sHandler(AllServices.%sService)", entityNameUpperFirst, entityNameUpperFirst, entityNameUpperFirst)
	lineToAddInHanldersStartKeyword := "var AllHandlers = Handler{"
	content = AddNewLineToExistContent(lineToAddInHanlders, content, lineToAddInHanldersStartKeyword, endKeyword, "", ",\n")

	fmt.Printf("Modified file: %s\n", filePath)
	return WriteFileInPath(filePath, content)
}

func AddNewLineToExistContent(newline, content, startKeyword, endKeyword, addToStartOfLine, addToEndOfLine string) string {

	if strings.Contains(content, newline) {
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

// getPackageName reads the go.mod file to extract the module name (package name)
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

func WriteFileInPath(filePath, content string) error {
	// Write the modified content back to the file
	var err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing to file %s: %v", filePath, err)
	}

	fmt.Printf("Modified file: %s\n", filePath)
	return nil
}
