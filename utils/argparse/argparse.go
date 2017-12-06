package argparse

import (
	"os"
	"strconv"
	"errors"
	"fmt"
)

func ValidateLength(expected int) error {
	if len(os.Args) < expected {
		return errors.New(fmt.Sprintf("To little arguments. Expected: %d.", expected))
	}
	return nil
}

func getValue(position int) (string, error) {
	if len(os.Args) <= position {
		return "", errors.New("Too little arguments.")
	}
	return os.Args[position], nil
}

func ReadDecimalInt(position int) (int, error) {
	value, err := getValue(position)
	if err != nil {
		return 0, err
	}
	parsed, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return 0, err
	}
	return int(parsed), nil
}

func ReadDecimalOrDefault(position int, default_ int) int {
	value, err := ReadDecimalInt(position)
	if err != nil {
		return default_
	}
	return value
}

func ReadString(position int) (string, error) {
	value, err := getValue(position)
	if err != nil {
		return "", err
	}
	return value, nil
}

func ReadStringOrDefault(position int, default_ string) string {
	value, err := ReadString(position)
	if err != nil {
		return default_
	}
	return value
}