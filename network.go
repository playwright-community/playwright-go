package playwright

import "strings"

func parseHeaders(headers []interface{}) map[string]string {
	out := make(map[string]string)
	for _, header := range headers {
		entry := header.(map[string]interface{})
		out[strings.ToLower(entry["name"].(string))] = entry["value"].(string)
	}
	return out
}
