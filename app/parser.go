package main

import (
	"fmt"
	"strconv"
	"strings"
)

func decodeRESPArray(input string) ([]string, error) {
	if !strings.HasPrefix(input, "*") {
		return nil, fmt.Errorf("Invalid RESP array format")
	}
	cleaned_input := strings.TrimPrefix(input, "*")
	split_input := strings.Split(cleaned_input, "\r\n")

	num_elements, err := strconv.Atoi(split_input[0])
	if err != nil {
		return nil, fmt.Errorf("Invalid number of elements: %v", err)
	}

	result := make([]string, 0, num_elements)

	for i := 1; i < len(split_input) && num_elements > 0; i++ {
		if strings.HasPrefix(split_input[i], "$") {
			str_len, err := strconv.Atoi(strings.TrimPrefix(split_input[i], "$"))
			if err != nil {
				return nil, fmt.Errorf("Invalid bulk string length: %v", err)
			}
			
			i++
			if i >= len(split_input) {
				return nil, fmt.Errorf("Unexpected end of input while reading bulk string")
			}
			
			if len(split_input[i]) != str_len {
				return nil, fmt.Errorf("Bulk string length mismatch: expected %d, got %d", str_len, len(split_input[i]))
			}
			result = append(result, split_input[i])
			num_elements--
		} else {
			return nil, fmt.Errorf("Expected bulk string (no other datatype supported yet), got: %s", split_input[i])
		}
	}

	return result, nil
}

func encodeBulkString(input string) string {
	return fmt.Sprintf("$%d\r\n%s\r\n", len(input), input)
}