package handlers

import (
	"errors"
	"strconv"
)

func ParseIntQueryParam(paramStr string, paramName string, defaultValue int) (int, error) {
	if paramStr == "" {
		return defaultValue, nil
	}

	parsed, err := strconv.Atoi(paramStr)
	if err != nil {
		return 0, errors.New("Invalid " + paramName)
	}
	return parsed, nil
}
