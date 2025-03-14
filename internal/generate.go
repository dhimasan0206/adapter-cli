package internal

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cbroglie/mustache"
)

func Generate(data map[string]string) error {
	// URL of the ZIP file
	url := "https://github.com/dhimasan0206/adapter-templates/archive/refs/heads/main.zip"

	// Step 1: Download the ZIP file into memory
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to download file: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to download file: status code %d\n", resp.StatusCode)
		return err
	}

	// Read the response body into a byte slice
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
		return err
	}

	// Step 2: Create a ZIP reader from the byte slice
	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		fmt.Printf("Failed to create ZIP reader: %v\n", err)
		return err
	}

	// Step 3: Iterate over the files in the ZIP archive
	out, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get current working directory: %v\n", err)
		return err
	}
	for _, file := range zipReader.File {
		paths := strings.Split(file.Name, "/")
		combinedPaths := append([]string{out}, paths[1:]...)
		target := filepath.Join(combinedPaths...)
		fmt.Printf("Reading file: %s\n", file.Name)

		// is directory
		if !strings.HasSuffix(target, ".mustache") {
			fmt.Println("Creating directory:", file.Name)
			err := os.MkdirAll(target, 0777)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				return err
			}
			fmt.Println("Directory created:", target)
			continue
		}

		// Open the file
		rc, err := file.Open()
		if err != nil {
			fmt.Printf("Failed to open file %s: %v\n", file.Name, err)
			return err
		}

		// Process the file (for example, read its content)
		content, err := io.ReadAll(rc)
		if err != nil {
			fmt.Printf("Failed to read file %s: %v\n", file.Name, err)
			rc.Close()
			return err
		}
		rc.Close()

		// Output the content or process it as needed
		// fmt.Printf("Content of %s: %s\n", file.Name, string(content))

		result, err := mustache.Render(string(content), data)
		if err != nil {
			fmt.Println("Error rendering template:", err)
			return err
		}
		newFile := strings.Replace(target, ".mustache", "", 1)
		err = os.WriteFile(newFile, []byte(result), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return err
		}
		fmt.Println(newFile + " has been generated successfully.")
	}
	return nil
}
