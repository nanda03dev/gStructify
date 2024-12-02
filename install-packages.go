package main

import (
	"fmt"
	"os/exec"
)

// RunGoModTidy runs `go mod tidy` in the specified directory
func RunGoModTidy(dir string) error {
	// Create the command to run `go mod tidy`
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = dir // Set the directory where the command should be executed

	// Run the command and capture the output or errors
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error running 'go mod tidy': %v\nOutput: %s", err, string(output))
	}

	fmt.Println("Installed go packages successfully!")
	fmt.Println(string(output))
	return nil
}
