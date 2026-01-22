package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Amiibo struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func main() {
	// Read raw data
	data, err := os.ReadFile("data/amiibo_raw.json")
	if err != nil {
		log.Fatal(err)
	}

	var amiibos []Amiibo
	if err := json.Unmarshal(data, &amiibos); err != nil {
		log.Fatal(err)
	}

	os.MkdirAll("frontend/public/images", 0755)

	updatedAmiibos := []Amiibo{}

	for _, a := range amiibos {
		filename := strings.ReplaceAll(strings.ToLower(a.Name), " ", "_")
		// Clean filename
		filename = strings.ReplaceAll(filename, ".", "")
		filename = strings.ReplaceAll(filename, "&", "and")
		filename = filepath.Base(filename) + ".png" // simplification

		localPath := filepath.Join("frontend/public/images", filename)
		publicURL := "/images/" + filename

		fmt.Printf("Downloading %s to %s...\n", a.Name, localPath)

		// Download
		if err := downloadFile(a.ImageURL, localPath); err != nil {
			log.Printf("Failed to download %s: %v", a.Name, err)
			// Keep original if failed
			updatedAmiibos = append(updatedAmiibos, a)
		} else {
			// Update URL to local
			a.ImageURL = publicURL
			updatedAmiibos = append(updatedAmiibos, a)
		}
	}

	// Save updated JSON
	newData, _ := json.MarshalIndent(updatedAmiibos, "", "  ")
	os.WriteFile("data/amiibo_local.json", newData, 0644)
	fmt.Println("Done! Saved to data/amiibo_local.json")
}

func downloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
