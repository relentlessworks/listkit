package model

import "regexp"

// matchRegex is a helper that uses Go's regexp package.
func matchRegex(item, pattern string) (bool, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false, err
	}
	return re.MatchString(item), nil
}

// CompileRegex compiles a regex pattern, returning an error if invalid.
func CompileRegex(pattern string) (*regexp.Regexp, error) {
	return regexp.Compile(pattern)
}

// FilterRegex filters items using a compiled regex.
func FilterRegex(items []string, re *regexp.Regexp) []string {
	var result []string
	for _, item := range items {
		if re.MatchString(item) {
			result = append(result, item)
		}
	}
	return result
}

// FindAll returns all regex matches for each item.
func FindAll(items []string, pattern string) []string {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return []string{}
	}
	var result []string
	for _, item := range items {
		matches := re.FindAllString(item, -1)
		result = append(result, matches...)
	}
	return result
}

// ReplaceRegex replaces regex matches with the replacement string.
func ReplaceRegex(items []string, pattern, replacement string) []string {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return items
	}
	result := make([]string, len(items))
	for i, item := range items {
		result[i] = re.ReplaceAllString(item, replacement)
	}
	return result
}
