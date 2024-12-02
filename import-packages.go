package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func ImportAllPacakges(wd string) {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".go" {
			cmd := exec.Command("goimports", "-w", path)
			if err := cmd.Run(); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Imported all local packages successfully!")
}
