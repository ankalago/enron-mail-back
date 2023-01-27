package utils

import "strings"

func RemoveNullValue(slice []interface{}) []interface{} {
	var output []interface{}
	for _, element := range slice {
		if element != nil {
			output = append(output, element)
		}
	}
	return output
}

func ValidateEmptyText(text string) string {
	if text != "" {
		return text
	}
	return "\n"
}

func FormatText(text string) string {
	txt := strings.Split(text, ": ")
	return txt[len(txt)-1]
}
