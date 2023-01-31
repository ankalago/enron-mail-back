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

func ChunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for {
		if len(slice) == 0 {
			break
		}

		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}

func LimitCharacters(text string) int {
	const limit = 70000
	limitCharacters := limit
	lengthText := len(text)
	if lengthText <= limit {
		limitCharacters = lengthText
	}
	return limitCharacters
}
