package pkg

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func Env() error {
	// Read dotenv
	content, err := os.ReadFile(".env")
	if err != nil {
		return errors.New("cannot read .env file")
	}

	// Split the content into line
	lines := bufio.NewScanner(strings.NewReader(string(content)))
	for lines.Scan() {
		line := lines.Text()
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		// Split the line into key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return errors.New("invalid line in .env file: " + line)
		}

		key := parts[0]
		value := parts[1]

		// Set env variable
		err1 := os.Setenv(key, value)
		if err1 != nil {
			return errors.New("cannot set environment variable: " + key + " = " + value)
		}
	}

	return lines.Err()
}
