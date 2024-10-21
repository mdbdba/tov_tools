package helpers

import "fmt"

// MapStringIntToString converts a map of integers keyed by strings to a
// comma-delimited string.
func MapStringIntToString(src map[string]int) (tgt string) {
	tgt = "["
	firstLoop := true
	keys := make([]string, 0, len(src))
	for k := range src {
		keys = append(keys, k)
	}
	for _, k := range keys {
		joinChr := ", "
		if firstLoop {
			joinChr = ""
			firstLoop = false
		}
		tgt = fmt.Sprintf("%s%s\"%s\": %d", tgt, joinChr, k, src[k])
	}
	tgt = fmt.Sprintf("%s]", tgt)
	return
}
