package main

import (
	"fmt"
	"os"
	"strings"
)

// modifyFile modifies the destination file content to include additional code
func modifyFile(filePath string, entity Entity) error {

	if strings.Contains(filePath, "repository.go") {
		return ToUpdateRepositoriesFile(filePath, entity)
	}

	if strings.Contains(filePath, "service.go") {
		return ToUpdateServicesFile(filePath, entity)
	}
	if strings.Contains(filePath, "handler.go") {
		return ToUpdateHandlersFile(filePath, entity)
	}

	// if strings.Contains(filePath, "app_module.go") {
	// 	return ToUpdateAppModuleFile(filePath, entityName)
	// }

	if strings.Contains(filePath, "route.go") {
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

	lineToAddInRepoType := fmt.Sprintf("%sRepository aggregate.%sRepository", entityNameUpperFirst, entityNameUpperFirst)
	repoStartKeyword := "type Repository struct {"
	endKeyword := "}"
	content = AddNewLineToStart(lineToAddInRepoType, content, repoStartKeyword, endKeyword, "", "")

	lineToAddInRepositories := fmt.Sprintf("%sRepository: repository.New%sRepository()", entityNameUpperFirst, entityNameUpperFirst)
	lineToAddInRepositoriesStartKeyword := "var AllRepositories = Repository{"
	content = AddNewLineToStart(lineToAddInRepositories, content, lineToAddInRepositoriesStartKeyword, endKeyword, "", ",")

	lineToAddInServiceType := fmt.Sprintf("%sService service.%sService", entityNameUpperFirst, entityNameUpperFirst)
	serviceStartKeyword := "type Service struct {"
	content = AddNewLineToStart(lineToAddInServiceType, content, serviceStartKeyword, endKeyword, "", "")

	lineToAddInServices := fmt.Sprintf("%sService: service.New%sService(Allrepository.%sRepository)", entityNameUpperFirst, entityNameUpperFirst, entityNameUpperFirst)
	lineToAddInServicesStartKeyword := "var AllServices = Service{"
	content = AddNewLineToStart(lineToAddInServices, content, lineToAddInServicesStartKeyword, endKeyword, "", ",")

	lineToAddInHandlerType := fmt.Sprintf("%sHandler handler.%sHandler", entityNameUpperFirst, entityNameUpperFirst)
	handlerStartKeyword := "type Handler struct {"
	content = AddNewLineToStart(lineToAddInHandlerType, content, handlerStartKeyword, endKeyword, "", "")

	lineToAddInHanlders := fmt.Sprintf("%sHandler: handler.New%sHandler(Allservice.%sService)", entityNameUpperFirst, entityNameUpperFirst, entityNameUpperFirst)
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

	lineToAddInServices := fmt.Sprintf("%sService: New%sService(AllRepository.%sRepository)", entityNameUpperFirst, entityNameUpperFirst, entityNameUpperFirst)
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
