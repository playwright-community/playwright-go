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

type rawHeaders struct {
	headersArray map[string]string
	headersMap   map[string][]string
}

func (r *rawHeaders) Get(name string) string {
	values := r.GetAll(name)
	sep := ", "
	if strings.ToLower(name) == "set-cookie" {
		sep = "; "
	}
	return strings.Join(values, sep)
}
func (r *rawHeaders) GetAll(name string) []string {
	return r.headersMap[strings.ToLower(name)]
}
func (r *rawHeaders) Headers() map[string]string {
	out := make(map[string]string)
	for key, value := range r.headersArray {
		out[strings.ToLower(key)] = strings.ToLower(value)
	}
	return out
}

func (r *rawHeaders) HeadersArray() map[string]string {
	return r.headersArray
}
func newRawHeaders(headers []interface{}) *rawHeaders {
	r := &rawHeaders{}
	r.headersArray = parseHeaders(headers)
	r.headersMap = make(map[string][]string)
	for _, header := range headers {
		entry := header.(map[string]interface{})
		name := strings.ToLower(entry["name"].(string))
		if _, ok := r.headersMap[name]; !ok {
			r.headersMap[name] = make([]string, 0)
			r.headersMap[name] = append(r.headersMap[name], entry["value"].(string))
		}
	}
	return r
}
