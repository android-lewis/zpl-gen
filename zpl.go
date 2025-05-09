package zpl

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
	"unicode"
)

// GenerateDetailMap takes a struct as input and returns a map
// where the keys are placeholder strings in the format <<FieldName>>
// and the values are the corresponding string values from the struct fields.
func GenerateDetailMap(details interface{}) map[string]string {
	detailsMap := make(map[string]string)
	v := reflect.ValueOf(details)
	t := reflect.TypeOf(details)

	for i := 0; i < t.NumField(); i++ {
		fieldName := t.Field(i).Name
		placeholderName := "<<" + fieldName + ">>"
		detailsMap[placeholderName] = v.Field(i).String()
	}

	return detailsMap
}

// GenerateLabelFile reads a ZPL label template file from disk and replaces all placeholder tokens using the provided map.
func GenerateLabelFile(filename string, detailsMap map[string]string) (string, error) {

	file, err := os.Open(filename)

	if err != nil {
		return "", fmt.Errorf("cannot open file at path %s", filename)
	}

	defer file.Close()
	var output bytes.Buffer
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		result := replacePlaceHolders(line, detailsMap)
		output.WriteString(result + "\n")
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return output.String(), nil
}

// GenerateLabelString processes a raw ZPL label string in memory, replacing placeholder tokens using the provided map.
func GenerateLabelString(file string, detailsMap map[string]string) (string, error) {
	var output string

	split := strings.Split(file, "\n")

	for _, line := range split {
		result := replacePlaceHolders(line, detailsMap)
		output += result + "\n"
	}

	return output, nil
}

// This helper function replaces placeholder tokens in a single line using the provided map.
func replacePlaceHolders(line string, detailsMap map[string]string) string {
	placeholderRegex := regexp.MustCompile(`<<[a-zA-Z0-9_]+>>`)

	result := placeholderRegex.ReplaceAllStringFunc(line, func(match string) string {
		if replacement, exists := detailsMap[match]; exists {
			return replacement
		}
		return match // Leave unchanged if no matching value
	})

	return result
}

// Removes all non-printable Unicode characters from a string.
func cleanString(text string) string {
	text = strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, text)

	return text
}
