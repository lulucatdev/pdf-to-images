package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: pdf-to-images <pdf-file-path>")
		os.Exit(1)
	}

	pdfPath := os.Args[1]
	if !strings.HasSuffix(strings.ToLower(pdfPath), ".pdf") {
		fmt.Println("Error: The provided file is not a PDF.")
		os.Exit(1)
	}

	// Get the absolute path
	absPdfPath, err := filepath.Abs(pdfPath)
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
	}

	// Get the directory and filename
	dir := filepath.Dir(absPdfPath)
	baseName := filepath.Base(absPdfPath)
	outputDirName := strings.TrimSuffix(baseName, filepath.Ext(baseName))
	outputDirPath := filepath.Join(dir, outputDirName)

	// Create the output directory
	err = os.MkdirAll(outputDirPath, 0755)
	if err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Construct the command to convert PDF to images
	cmd := exec.Command(
		"pdftoppm",
		"-jpeg",
		"-r", "300",
		"-jpegopt", "quality=100",
		absPdfPath,
		filepath.Join(outputDirPath, "page"),
	)

	// Run the command
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to convert PDF to images: %v", err)
	}

	// Rename the output files to match the page number
	files, err := filepath.Glob(filepath.Join(outputDirPath, "page-*.jpg"))
	if err != nil {
		log.Fatalf("Failed to list image files: %v", err)
	}

	for _, oldPath := range files {
		fileName := filepath.Base(oldPath)
		// Extract page number
		parts := strings.Split(fileName, "-")
		if len(parts) != 2 {
			continue
		}
		pageNum := strings.TrimSuffix(parts[1], ".jpg")
		newFileName := fmt.Sprintf("%s.jpeg", pageNum)
		newPath := filepath.Join(outputDirPath, newFileName)
		err = os.Rename(oldPath, newPath)
		if err != nil {
			log.Fatalf("Failed to rename file %s to %s: %v", oldPath, newFileName, err)
		}
	}

	fmt.Printf("Images saved in %s\n", outputDirPath)
}
