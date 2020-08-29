package playwright

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
