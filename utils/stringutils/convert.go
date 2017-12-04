package stringutils

func StringToChars(str string) []byte {
	result := make([]byte, len(str))

	for idx := 0; idx < len(str); idx++ {
		result[idx] = str[idx]
	}

	return result
}

func CharsToStrings(chars []byte) []string {
	results := make([]string, len(chars))
	for idx, val := range chars {
		results[idx] = string(val)
	}
	return results
}

