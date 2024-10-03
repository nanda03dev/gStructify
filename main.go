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
var EntityName string

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

	if *entity != "" {
		EntityName = *entity
		fmt.Println("creating layers for entity : ", EntityName)
		return
	}

	// Generate the microservice using the package name
	CreateNewMS(wd, packageName)
	ImportAllPacakges(wd)

}

// CreateNewMS creates the microservice by copying and modifying the template files
func CreateNewMS(outputDir, packageName string) {
	// Copy and modify template files
	err := copyDirAndModify(cleanTemplate, "clean-template", outputDir, packageName)
	if err != nil {
		panic(err)
	}

	println("Files copied and modified successfully!")
}

// copyDirAndModify copies template files and modifies them with the package name
func copyDirAndModify(efs embed.FS, srcDir, destDir, packageName string) error {
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
			err = copyDirAndModify(efs, srcPath, destPath, packageName)
			if err != nil {
				return err
			}
		} else {
			data, err := efs.ReadFile(srcPath)
			if err != nil {
				return err
			}

			// Rename file if it contains "template_entity"
			if strings.Contains(entry.Name(), "template_entity") {
				if EntityName == "" {
					continue
				}
				destPath = strings.ReplaceAll(destPath, "template_entity", "user")
			}

			// Replace content inside the file
			content := string(data)
			content = strings.ReplaceAll(content, "TemplateEntity", "User")
			content = strings.ReplaceAll(content, "templateEntity", "user")
			content = strings.ReplaceAll(content, "github.com/nanda03dev/go-ms-template", packageName)

			err = os.WriteFile(destPath, []byte(content), 0644)
			if err != nil {
				return err
			}
		}
	}

	return nil
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
