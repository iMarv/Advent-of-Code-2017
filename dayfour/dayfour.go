package dayfour

import "sort"

// ValidPhrasesCount returns amount of correct phrases in two dimensional array
func ValidPhrasesCount(passphraseCollection [][]string, validator func([]string) bool) (valid int) {
	for _, phrase := range passphraseCollection {
		if validator(phrase) {
			valid++
		}
	}
	return valid
}

// NoDuplicatesValidator checks passphrases for duplicate entries
func NoDuplicatesValidator(phrases []string) bool {
	phr := make(map[string]bool)
	for _, phrase := range phrases {
		phr[phrase] = true
	}

	return len(phrases) == len(phr)
}

// NoAnagramsValidator checks passphrase for entries which are anagram of each other
func NoAnagramsValidator(phrases []string) bool {
	phr := make(map[string]bool)
	for _, phrase := range phrases {
		phr[orderString(phrase)] = true
	}

	return len(phrases) == len(phr)
}

func orderString(s string) string {
	var chars []rune
	for _, c := range s {
		chars = append(chars, c)
	}

	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	return string(chars)
}
