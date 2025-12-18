package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	purehtml "github.com/purescraps/purehtml/go"
)

func main() {
	configPath := flag.String("config", "", "path to purehtml config file (YAML)")
	flag.Parse()

	if *configPath == "" {
		log.Fatal("config flag is required")
	}

	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("input HTML path or URL is required")
	}
	input := args[0]

	// Load config
	config, err := loadConfig(*configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Get HTML content
	html, err := getHTML(input)
	if err != nil {
		log.Fatalf("failed to get HTML: %v", err)
	}

	// Extract data
	result, err := purehtml.ExtractFromString(purehtml.DefaultBackend, html, config, input)
	if err != nil {
		log.Fatalf("failed to extract: %v", err)
	}

	// Output result as JSON
	output, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal result: %v", err)
	}
	fmt.Println(string(output))
}

func loadConfig(path string) (purehtml.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	factory := purehtml.NewConfigFactory()
	config, err := factory.FromYAML(string(data))
	if err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}

	return config, nil
}

func getHTML(input string) (string, error) {
	// Check if input is a URL
	if u, err := url.Parse(input); err == nil && (u.Scheme == "http" || u.Scheme == "https") {
		return fetchURL(input)
	}

	// Otherwise treat as file path
	data, err := os.ReadFile(input)
	if err != nil {
		return "", fmt.Errorf("reading file: %w", err)
	}
	return string(data), nil
}

func fetchURL(targetURL string) (string, error) {
	resp, err := http.Get(targetURL)
	if err != nil {
		return "", fmt.Errorf("fetching URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response body: %w", err)
	}

	return string(body), nil
}
