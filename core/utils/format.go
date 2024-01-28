package utils

import (
	"encoding/json"
	"path"
	"regexp"
)

func FormatPath(paths ...string) (r string) {

	for i := range paths {
		r = path.Join(r, paths[i])
	}

	return
}

func ValidEthAddress(addr string) bool {

	re := regexp.MustCompile(`(?i)(0x)?[0-9a-f]{40}`) // (?i) case-insensitive, (0x)? optional hex prefix
	return re.MatchString(addr)
}

func MarshalJSON(v interface{}) string {

	// Encode any struct to JSON
	bytes, err := json.Marshal(v)
	CheckError(err, WarningMode)

	return string(bytes)
}

func UnmarshalJSON(data string, v interface{}) {

	// String to bytes slice
	bytes := []byte(data)

	// Decode JSON to any struct
	err := json.Unmarshal(bytes, v)
	CheckError(err, WarningMode)
}
