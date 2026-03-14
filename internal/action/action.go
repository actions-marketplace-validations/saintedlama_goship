package action

import (
	"fmt"
	"os"
)

// Config holds all configuration parsed from GitHub Actions inputs.
type Config struct {
	Token            string
	WorkingDirectory string
}

// Run executes the action and returns its primary output value.
func Run(cfg Config) (string, error) {
	if err := os.Chdir(cfg.WorkingDirectory); err != nil {
		return "", fmt.Errorf("change working directory to %q: %w", cfg.WorkingDirectory, err)
	}

	// TODO: implement action logic here.

	return "ok", nil
}
