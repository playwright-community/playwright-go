package playwright

import "strings"

func serializeHeaders(headers map[string]string) []map[string]string {
	serialized := make([]map[string]string, 0)
	for name, value := range headers {
		serialized = append(serialized, map[string]string{
			"name":  name,
			"value": value,
		})
	}
	return serialized
}

func parseHeaders(headers []interface{}) map[string]string {
	out := make(map[string]string)
	for _, header := range headers {
		entry := header.(map[string]interface{})
		out[strings.ToLower(entry["name"].(string))] = entry["value"].(string)
	}
	return out
}
