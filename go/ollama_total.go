package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// getOllamaTotal calculates the total storage used by all Ollama models
func getOllamaTotal() (string, error) {
	// Run the ollama list command
	cmd := exec.Command("ollama", "list")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute 'ollama list': %w", err)
	}

	// Convert to string and split by newline
	outputStr := string(output)
	lines := strings.Split(strings.TrimSpace(outputStr), "\n")

	// Skip the header line
	if len(lines) <= 1 {
		return "0 B", nil
	}
	lines = lines[1:]

	// Regular expression to match size values like "2.7 GB", "17 GB", "800 MB", etc.
	sizePattern := regexp.MustCompile(`(\d+(?:\.\d+)?) (MB|GB|TB)`)

	var totalBytes float64

	// Process each line
	for _, line := range lines {
		matches := sizePattern.FindStringSubmatch(line)
		if len(matches) == 3 {
			sizeValue, err := strconv.ParseFloat(matches[1], 64)
			if err != nil {
				continue
			}

			sizeUnit := matches[2]

			// Convert to bytes for consistent calculation
			switch sizeUnit {
			case "MB":
				totalBytes += sizeValue * 1000000
			case "GB":
				totalBytes += sizeValue * 1000000000
			case "TB":
				totalBytes += sizeValue * 1000000000000
			}
		}
	}

	// Convert back to the most appropriate unit with one decimal place
	if totalBytes >= 1000000000000 {
		return fmt.Sprintf("%.1f TB", totalBytes/1000000000000), nil
	} else if totalBytes >= 1000000000 {
		return fmt.Sprintf("%.1f GB", totalBytes/1000000000), nil
	} else if totalBytes >= 1000000 {
		return fmt.Sprintf("%.1f MB", totalBytes/1000000), nil
	}

	return fmt.Sprintf("%.0f B", totalBytes), nil
}

func main() {
	total, err := getOllamaTotal()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Println(total)
}
