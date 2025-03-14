package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// URL of the ZIP file
	url := "https://github.com/dhimasan0206/adapter-templates/archive/refs/heads/main.zip"

	// Step 1: Download the ZIP file into memory
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to download file: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to download file: status code %d\n", resp.StatusCode)
		return
	}

	// Read the response body into a byte slice
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
		return
	}

	// Step 2: Create a ZIP reader from the byte slice
	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		fmt.Printf("Failed to create ZIP reader: %v\n", err)
		return
	}

	// Step 3: Iterate over the files in the ZIP archive
	for _, file := range zipReader.File {
		fmt.Printf("Reading file: %s\n", file.Name)

		// Open the file
		rc, err := file.Open()
		if err != nil {
			fmt.Printf("Failed to open file %s: %v\n", file.Name, err)
			continue
		}

		// Process the file (for example, read its content)
		content, err := io.ReadAll(rc)
		if err != nil {
			fmt.Printf("Failed to read file %s: %v\n", file.Name, err)
			rc.Close()
			continue
		}
		rc.Close()

		// Output the content or process it as needed
		fmt.Printf("Content of %s: %s\n", file.Name, string(content))
	}
}
