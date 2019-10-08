// Package env is a simple package to read environment variable files and load them into the OS variables.
// It parses env files and returns their key/values as a string map.
// It also loads or overloads them into the OS variables.
package env

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

// Load will read the given env file and return its variables as a map, it also load variables into the OS variables.
// It won't overwrite the existing variables in the OS.
func Load(filename string) (map[string]string, error) {
	return load(filename, false)
}

// Overload will read the given env file and return its variables as a map, it also load variables into the OS variables.
// It will overwrite the existing variables in the OS.
func Overload(filename string) (map[string]string, error) {
	return load(filename, true)
}

// load will read the given env file and return its variables as a map, it also load variables into the OS variables.
// The overwrite parameter determine that it should overwrite OS variables or not.
func load(filename string, overwrite bool) (map[string]string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	path := wd + string(os.PathSeparator) + filename
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	variables, err := read(file)
	if err != nil {
		return nil, err
	}

	for k, v := range variables {
		if overwrite == false && os.Getenv(k) != "" {
			continue
		}

		if err := os.Setenv(k, v); err != nil {
			return nil, err
		}
	}

	return variables, nil
}

// read will open the given file and return the file content as a map.
func read(file *os.File) (map[string]string, error) {
	e := map[string]string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if key, value, err := parse(scanner.Text()); err != nil {
			return nil, err
		} else if key != "" {
			e[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return e, nil
}

// parse will extract env key/values from the given line
func parse(line string) (string, string, error) {
	ln := strings.TrimSpace(line)

	if len(ln) == 0 {
		return "", "", nil
	}

	if ln[0] == '#' {
		return "", "", nil
	}

	s := strings.Index(ln, "=")
	if s == -1 {
		return "", "", errors.New("Invalid ln: " + ln)
	}

	k := strings.TrimSpace(ln[:s])
	v := strings.TrimSpace(ln[s+1:])

	return k, v, nil
}
