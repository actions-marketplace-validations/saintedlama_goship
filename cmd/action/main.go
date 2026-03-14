package main

import (
	"fmt"
	"log"
	"os"

	"github.com/saintedlama/goship/internal/action"
)

func main() {
	cfg := action.Config{
		Token:            getInput("GITHUB_TOKEN"),
		WorkingDirectory: getInputWithDefault("WORKING_DIRECTORY", "."),
	}

	result, err := action.Run(cfg)
	if err != nil {
		writeError(err.Error())
		os.Exit(1)
	}

	if err := setOutput("result", result); err != nil {
		log.Fatalf("failed to set output: %v", err)
	}
}

// getInput reads a GitHub Actions input from the environment.
// For docker actions the env mapping is declared in action.yml under runs.env.
func getInput(key string) string {
	return os.Getenv(key)
}

func getInputWithDefault(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

// setOutput writes an output to $GITHUB_OUTPUT (file-based protocol).
func setOutput(key, value string) error {
	outputFile := os.Getenv("GITHUB_OUTPUT")
	if outputFile == "" {
		// Fallback for local development / testing.
		fmt.Printf("::set-output name=%s::%s\n", key, value)
		return nil
	}
	f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return fmt.Errorf("open GITHUB_OUTPUT: %w", err)
	}
	defer f.Close()
	_, err = fmt.Fprintf(f, "%s=%s\n", key, value)
	return err
}

// writeError emits a GitHub Actions error annotation to stderr.
func writeError(msg string) {
	fmt.Fprintf(os.Stderr, "::error::%s\n", msg)
}
