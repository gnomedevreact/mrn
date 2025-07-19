package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: generate <module_name> <output_dir>")
		return
	}

	module := os.Args[1]
	outputDir := os.Args[2]
	targetDir := filepath.Join(outputDir, module)

	err := os.MkdirAll(targetDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to create directory: %v\n", err)
		return
	}

	files := []string{"handlers.go", "models.go", "services.go"}
	for _, file := range files {
		path := filepath.Join(targetDir, file)
		content := fmt.Sprintf("package %s\n\n", module)
		err := os.WriteFile(path, []byte(content), 0644)
		if err != nil {
			fmt.Printf("Failed to create file %s: %v\n", file, err)
			return
		}
	}

	fmt.Printf("Module '%s' generated in %s\n", module, targetDir)
}
