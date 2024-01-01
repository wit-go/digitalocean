package digitalocean

import 	(
	"encoding/json"
)

// formatJSON takes an unformatted JSON string and returns a formatted version.
func FormatJSON(unformattedJSON string) (string, error) {
	var jsonData interface{}

	// Decode the JSON string into an interface
	err := json.Unmarshal([]byte(unformattedJSON), &jsonData)
	if err != nil {
		return "", err
	}

	// Re-encode the JSON with indentation for formatting
	formattedJSON, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		return "", err
	}

	return string(formattedJSON), nil
}
