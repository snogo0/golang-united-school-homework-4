package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
// For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	if input == "" {
		return "", fmt.Errorf("Nothing to parse: %w", errorEmptyInput)
	}
	runes := make([]rune, 0)
	for _, r := range input {
		switch {
		case unicode.IsSpace(r):
			// skip appending
		case strings.ContainsAny(string(r), "+-"):
			if len(runes) > 0 {
				runes = append(runes, rune(" "[0]))
			}
			runes = append(runes, r)
		default:
			runes = append(runes, r)
		}
	}
	operands := strings.Split(string(runes), " ")
	if len(operands) != 2 {
		return "", fmt.Errorf("Wrong number of operands: %w", errorNotTwoOperands)
	}
	sum := 0
	for _, op := range operands {
		val, err := strconv.Atoi(string(op))
		if err != nil {
			return "", fmt.Errorf("Conversation error: %w", err)
		}
		sum += val
	}
	return strconv.Itoa(sum), nil
}
