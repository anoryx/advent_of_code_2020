package aocutil

import (
	"regexp"
	"strconv"
)

// ContainsString searches a string of arrays for a string member
func ContainsString(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

// GetRegexpMap parses a regexp and returns a map of named matches
func GetRegexpMap(re *regexp.Regexp, s string) map[string]string {
	matches := re.FindStringSubmatch(s)
	matchNames := re.SubexpNames()
	matchMap := map[string]string{}
	for i, n := range matches {
		matchMap[matchNames[i]] = n
	}
	return matchMap
}

// MustAtoi converts string to int, or panics
func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
