package env

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func New(filename string) (map[string]string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	path := wd + string(os.PathSeparator) + filename

	e := map[string]string{}

	if err := load(path, e); err != nil {
		return nil, err
	}

	return e, nil
}

// load will open a single env file and fill the given env variable
// It will return error if cannot open the file or parse it
func load(path string, e map[string]string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if key, value, err := parse(line); err != nil {
			return err
		} else if key != "" {
			e[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// parse will extract key and value from the given line.
// It will return error if the line is invalid.
func parse(line string) (string, string, error) {
	line = strings.TrimSpace(line)

	if len(line) == 0 {
		return "", "", nil
	}

	if line[0] == '#' {
		return "", "", nil
	}

	s := strings.Index(line, "=")
	if s == -1 {
		return "", "", errors.New("Invalid line: " + line)
	}

	k := strings.TrimSpace(line[:s])
	v := strings.TrimSpace(line[s+1:])

	return k, v, nil
}
