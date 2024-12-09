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
var msName string

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
		fmt.Printf("go.mod file does not exist in the current directory %v", err)
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

	// Split the string by "/"
	parts := strings.Split(packageName, "/")
	// Get the last part
	msName = parts[len(parts)-1]

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
	fmt.Println("packageName ", packageName)

	err := copyDirAndModify(cleanTemplate, "clean-template", outputDir, packageName, entityName)
	if err != nil {
		panic(err)
	}
}

// copyDirAndModify recursively copies the contents of a source directory (srcDir) from the embedded file system (efs)
// to a destination directory (destDir). During the copying process, it modifies file names and contents based on
// the provided packageName and entityName.
func copyDirAndModify(efs embed.FS, srcDir, destDir, packageName string, entityName string) error {

	// Read the list of entries (files and directories) in the source directory
	entries, err := efs.ReadDir(srcDir)
	if err != nil {
		return err // Return the error if the directory can't be read
	}

	for _, entry := range entries {
		// Construct the source and destination paths for the current entry
		srcPath := filepath.Join(srcDir, entry.Name())
		destPath := filepath.Join(destDir, entry.Name())

		if entry.IsDir() {
			// If the entry is a directory, create the corresponding directory in the destination
			err = os.MkdirAll(destPath, 0755)
			if err != nil {
				fmt.Println("Failed to create the destination folder due to an error")
				return err
			}

			// Recursively process the subdirectory
			err = copyDirAndModify(efs, srcPath, destPath, packageName, entityName)
			if err != nil {
				return err
			}
		} else {
			// If the entry is a file, read its content
			data, err := efs.ReadFile(srcPath)
			if err != nil {
				fmt.Println("Failed to read the source file due to an error")
				return err
			}

			// Check if the destination file already exists
			if _, err := os.Stat(destPath); err == nil {
				// If the file exists, modify its content
				err = modifyFile(destPath, entityName)
				if err != nil {
					return err
				}
			} else {
				// Rename the file if it contains "template_entity" in its name
				if strings.Contains(entry.Name(), "template_entity") {
					// if entityName == "" {
					// 	continue // Skip renaming if entityName is empty
					// }
					destPath = strings.ReplaceAll(destPath, "template_entity", entityName)
				}

				// Replace placeholders in the file content
				content := string(data)
				content = replaceEntityName(content, entityName)
				content = strings.ReplaceAll(content, "github.com/nanda03dev/go-ms-template", packageName)

				// Write the modified content to the destination file
				err = os.WriteFile(destPath, []byte(content), 0644)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil // Return nil if the operation completes successfully
}

// modifyFile modifies the destination file content to include additional code
func modifyFile(filePath, entityName string) error {
	if strings.Contains(filePath, "app_module.go") {
		return ToUpdateAppModuleFile(filePath, entityName)
	}

	if strings.Contains(filePath, "routes.go") {
		return ToUpdateRouterFile(filePath, entityName)
	}

	if strings.Contains(filePath, "response-messages.go") {
		return ToUpdateCommonResponseMessage(filePath, entityName)
	}

	return nil
}

func ToUpdateRouterFile(filePath, entityName string) error {

	var newLine = `
	// TemplateEntity CRUD API'S
	templateEntityHandler := appModule.Handler.TemplateEntityHandler
	templateEntityV1Routes := api.Group("/v1/templateEntity")
	templateEntityV1Routes.Post("/", templateEntityHandler.CreateTemplateEntity)
	templateEntityV1Routes.Get("/:id", templateEntityHandler.GetTemplateEntityByID)
	templateEntityV1Routes.Put("/:id", templateEntityHandler.UpdateTemplateEntityById)
	templateEntityV1Routes.Delete("/:id", templateEntityHandler.DeleteTemplateEntityById)
	`

	newLine = replaceEntityName(newLine, entityName)

	// Read the existing file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", filePath, err)
	}

	content := string(data)

	handlerLine := "templateEntityHandler := appModule.Handler.TemplateEntityHandler"

	if !strings.Contains(content, replaceEntityName(handlerLine, entityName)) {
		startKeyword := "func InitializeRoutes(app *fiber.App) {"
		endKeyword := "}"
		content = AddNewLineToExistContent(newLine, content, startKeyword, endKeyword, "\n", "\n")
	}

	return WriteFileInPath(filePath, content)
}

func ToUpdateCommonResponseMessage(filePath, entityName string) error {

	var newLine = `TemplateEntityNotFoundError = "templateEntity not found"`
	newLine = replaceEntityName(newLine, entityName)

	// Read the existing file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", filePath, err)
	}

	content := string(data)

	startKeyword := "const ("
	endKeyword := ")"
	content = AddNewLineToExistContent(newLine, content, startKeyword, endKeyword, fmt.Sprintf("\n //%s \n", ToUpperFirst(entityName)), "")

	fmt.Printf("Modified file: %s\n", filePath)
	return WriteFileInPath(filePath, content)
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

	return WriteFileInPath(filePath, content)
}
