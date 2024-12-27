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

	config, configErr := getConfigFile(wd)

	// Accept the package name as a command-line argument
	entity := flag.String("entity", "", "Name of the package (e.g., book)")
	flag.Parse()

	var entityName = *entity
	if entityName == "" && configErr != nil {
		fmt.Println("Entity name not specified.")
		fmt.Println("Please specify an entity name using the '-entity' flag or add entity details in the 'gStructify.config.json' file.")
		return
	}

	config = getUpdatedConfig(entityName, config)

	fmt.Println("Created layers for entity : ", entityName, " successfully!")

	// Split the string by "/"
	parts := strings.Split(packageName, "/")
	// Get the last part
	msName = parts[len(parts)-1]

	for _, eachEntity := range config.Entities {
		// Generate the microservice using the package name
		CreateNewMS(wd, packageName, eachEntity)
	}

	// This will import all required local packages
	ImportAllPacakges(wd)

	// Run `go mod tidy` in the working directory
	err = RunGoModTidy(wd)
	if err != nil {
		fmt.Printf("Error running 'go mod tidy': %v\n", err)
	}

}

// CreateNewMS creates the microservice by copying and modifying the template files
func CreateNewMS(outputDir, packageName string, entity Entity) {
	// Copy and modify template files
	err := copyDirAndModify(cleanTemplate, "clean-template", outputDir, packageName, entity)
	if err != nil {
		panic(err)
	}
}

// copyDirAndModify recursively copies the contents of a source directory (srcDir) from the embedded file system (efs)
// to a destination directory (destDir). During the copying process, it modifies file names and contents based on
// the provided packageName and entityName.
func copyDirAndModify(efs embed.FS, srcDir, destDir, packageName string, entity Entity) error {

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
			err = copyDirAndModify(efs, srcPath, destPath, packageName, entity)
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
				err = modifyFile(destPath, entity)
				if err != nil {
					return err
				}
			} else {
				// Rename the file if it contains "template_entity" in its name
				if strings.Contains(entry.Name(), "template_entity") {
					// if entityName == "" {
					// 	continue // Skip renaming if entityName is empty
					// }
					destPath = replaceFileName(destPath, entity)
				}

				// Replace placeholders in the file content
				content := string(data)
				content = replaceEntityName(content, entity)
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
func modifyFile(filePath string, entity Entity) error {

	if strings.Contains(filePath, "repositories.go") {
		return ToUpdateRepositoriesFile(filePath, entity)
	}

	if strings.Contains(filePath, "services.go") {
		return ToUpdateServicesFile(filePath, entity)
	}
	if strings.Contains(filePath, "handlers.go") {
		return ToUpdateHandlersFile(filePath, entity)
	}

	// if strings.Contains(filePath, "app_module.go") {
	// 	return ToUpdateAppModuleFile(filePath, entityName)
	// }

	if strings.Contains(filePath, "routes.go") {
		return ToUpdateRouterFile(filePath, entity)
	}

	if strings.Contains(filePath, "response_messages.go") {
		return ToUpdateCommonResponseMessage(filePath, entity)
	}

	if strings.Contains(filePath, "entities.go") {
		return ToUpdateEntityFile(filePath, entity)
	}

	return nil
}

func ToUpdateRouterFile(filePath string, entity Entity) error {

	var newLine = `
	// TemplateEntity CRUD API'S
	templateEntityHandler := AllHandlers.TemplateEntityHandler
	templateEntityV1Routes := api.Group("/v1/templateEntity")
	templateEntityV1Routes.Post("/", templateEntityHandler.CreateTemplateEntity)
	templateEntityV1Routes.Get("/:id", templateEntityHandler.GetTemplateEntityByID)
	templateEntityV1Routes.Post("/filter", templateEntityHandler.FindTemplateEntityWithFilter)
	templateEntityV1Routes.Put("/:id", templateEntityHandler.UpdateTemplateEntityById)
	templateEntityV1Routes.Delete("/:id", templateEntityHandler.DeleteTemplateEntityById)
	`

	newLine = replaceEntityName(newLine, entity)

	// Read the existing file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", filePath, err)
	}

	content := string(data)

	handlerLine := "templateEntityHandler := AllHandlers.TemplateEntityHandler"

	if !strings.Contains(content, replaceEntityName(handlerLine, entity)) {
		startKeyword := "func InitializeRoutes(fiberApp *fiber.App) {"
		endKeyword := "}"
		content = AddNewLineToEnd(newLine, content, startKeyword, endKeyword, "", "")
	}

	return WriteFileInPath(filePath, content)
}

func ToUpdateCommonResponseMessage(filePath string, entity Entity) error {

	var newLine = `TemplateEntityNotFoundError = "templateEntity not found"`
	newLine = replaceEntityName(newLine, entity)

	// Read the existing file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", filePath, err)
	}

	content := string(data)

	startKeyword := "const ("
	endKeyword := ")"
	content = AddNewLineToEnd(newLine, content, startKeyword, endKeyword, fmt.Sprintf("\n //%s \n", ToUpperFirst(entity.EntityName)), "")

	return WriteFileInPath(filePath, content)
}

func ToUpdateAppModuleFile(filePath string, entity Entity) error {
	// Read the existing file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", filePath, err)
	}

	content := string(data)
	entityNameUpperFirst := ToUpperFirst(entity.EntityName)

	lineToAddInRepoType := fmt.Sprintf("%sRepository aggregates.%sRepository", entityNameUpperFirst, entityNameUpperFirst)
	repoStartKeyword := "type Repository struct {"
	endKeyword := "}"
	content = AddNewLineToStart(lineToAddInRepoType, content, repoStartKeyword, endKeyword, "", "")

	lineToAddInRepositories := fmt.Sprintf("%sRepository: repositories.New%sRepository()", entityNameUpperFirst, entityNameUpperFirst)
	lineToAddInRepositoriesStartKeyword := "var AllRepositories = Repository{"
	content = AddNewLineToStart(lineToAddInRepositories, content, lineToAddInRepositoriesStartKeyword, endKeyword, "", ",")

	lineToAddInServiceType := fmt.Sprintf("%sService services.%sService", entityNameUpperFirst, entityNameUpperFirst)
	serviceStartKeyword := "type Service struct {"
	content = AddNewLineToStart(lineToAddInServiceType, content, serviceStartKeyword, endKeyword, "", "")

	lineToAddInServices := fmt.Sprintf("%sService: services.New%sService(AllRepositories.%sRepository)", entityNameUpperFirst, entityNameUpperFirst, entityNameUpperFirst)
	lineToAddInServicesStartKeyword := "var AllServices = Service{"
	content = AddNewLineToStart(lineToAddInServices, content, lineToAddInServicesStartKeyword, endKeyword, "", ",")

	lineToAddInHandlerType := fmt.Sprintf("%sHandler handlers.%sHandler", entityNameUpperFirst, entityNameUpperFirst)
	handlerStartKeyword := "type Handler struct {"
	content = AddNewLineToStart(lineToAddInHandlerType, content, handlerStartKeyword, endKeyword, "", "")

	lineToAddInHanlders := fmt.Sprintf("%sHandler: handlers.New%sHandler(AllServices.%sService)", entityNameUpperFirst, entityNameUpperFirst, entityNameUpperFirst)
	lineToAddInHanldersStartKeyword := "var AllHandlers = Handler{"
	content = AddNewLineToStart(lineToAddInHanlders, content, lineToAddInHanldersStartKeyword, endKeyword, "", ",")

	return WriteFileInPath(filePath, content)
}

func ToUpdateEntityFile(filePath string, entity Entity) error {
	// Read the existing file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", filePath, err)
	}

	content := string(data)
	entityNameUpperFirst := ToUpperFirst(entity.EntityName)

	lineToAddInRepoType := fmt.Sprintf("&%s{},", entityNameUpperFirst)
	repoStartKeyword := "var Entities = []interface{}{"
	endKeyword := "}"
	content = AddNewLineToStart(lineToAddInRepoType, content, repoStartKeyword, endKeyword, "", "")

	return WriteFileInPath(filePath, content)
}

func ToUpdateRepositoriesFile(filePath string, entity Entity) error {
	// Read the existing file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", filePath, err)
	}

	content := string(data)
	entityNameUpperFirst := ToUpperFirst(entity.EntityName)

	lineToAddInRepoType := fmt.Sprintf("%sRepository %sRepository", entityNameUpperFirst, entityNameUpperFirst)
	repoStartKeyword := "type Repositories struct {"
	endKeyword := "}"
	content = AddNewLineToStart(lineToAddInRepoType, content, repoStartKeyword, endKeyword, "", "")

	lineToAddInRepositories := fmt.Sprintf("%sRepository: New%sRepository(databases)", entityNameUpperFirst, entityNameUpperFirst)
	lineToAddInRepositoriesStartKeyword := "allRepositories = &Repositories{"
	content = AddNewLineToStart(lineToAddInRepositories, content, lineToAddInRepositoriesStartKeyword, endKeyword, "", ",")

	return WriteFileInPath(filePath, content)
}

func ToUpdateServicesFile(filePath string, entity Entity) error {
	// Read the existing file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", filePath, err)
	}

	content := string(data)
	entityNameUpperFirst := ToUpperFirst(entity.EntityName)
	endKeyword := "}"

	lineToAddInServiceType := fmt.Sprintf("%sService %sService", entityNameUpperFirst, entityNameUpperFirst)
	serviceStartKeyword := "type Services struct {"
	content = AddNewLineToStart(lineToAddInServiceType, content, serviceStartKeyword, endKeyword, "", "")

	lineToAddInServices := fmt.Sprintf("%sService: New%sService(AllRepositories.%sRepository)", entityNameUpperFirst, entityNameUpperFirst, entityNameUpperFirst)
	lineToAddInServicesStartKeyword := "allServices = &Services{"
	content = AddNewLineToStart(lineToAddInServices, content, lineToAddInServicesStartKeyword, endKeyword, "", ",")

	return WriteFileInPath(filePath, content)
}

func ToUpdateHandlersFile(filePath string, entity Entity) error {
	// Read the existing file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", filePath, err)
	}

	content := string(data)
	entityNameUpperFirst := ToUpperFirst(entity.EntityName)
	endKeyword := "}"

	lineToAddInHandlerType := fmt.Sprintf("%sHandler %sHandler", entityNameUpperFirst, entityNameUpperFirst)

	handlerStartKeyword := "type Handlers struct {"
	content = AddNewLineToStart(lineToAddInHandlerType, content, handlerStartKeyword, endKeyword, "", "")

	lineToAddInHanlders := fmt.Sprintf("%sHandler: New%sHandler(AllServices.%sService)", entityNameUpperFirst, entityNameUpperFirst, entityNameUpperFirst)
	lineToAddInHanldersStartKeyword := "allHandlers = &Handlers{"
	content = AddNewLineToStart(lineToAddInHanlders, content, lineToAddInHanldersStartKeyword, endKeyword, "", ",")

	return WriteFileInPath(filePath, content)
}
