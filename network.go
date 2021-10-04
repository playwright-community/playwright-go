package playwright

import (
	"strings"
)

type rawHeaders struct {
	headersArray []map[string]string
	headersMap   map[string]map[string]bool
}

func (r *rawHeaders) Get(name string) string {
	values := r.GetAll(name)
	if len(values) == 0 {
		return ""
	}
	sep := ", "
	if strings.ToLower(name) == "set-cookie" {
		sep = "\n"
	}
	return strings.Join(values, sep)
}
func (r *rawHeaders) GetAll(name string) []string {
	name = strings.ToLower(name)
	out := make([]string, 0)
	for value := range r.headersMap[name] {
		out = append(out, value)
	}
	return out
}
func (r *rawHeaders) Headers() map[string]string {
	out := make(map[string]string)
	for key := range r.headersMap {
		out[key] = r.Get(key)
	}
	return out
}

func (r *rawHeaders) HeadersArray() []map[string]string {
	return r.headersArray
}
func newRawHeaders(headers interface{}) *rawHeaders {
	r := &rawHeaders{}
	r.headersArray = make([]map[string]string, 0)
	for _, header := range headers.([]interface{}) {
		entry := header.(map[string]interface{})
		r.headersArray = append(r.headersArray, map[string]string{
			strings.ToLower(entry["name"].(string)): entry["value"].(string),
		})
	}
	r.headersMap = make(map[string]map[string]bool)
	for _, header := range headers.([]interface{}) {
		entry := header.(map[string]interface{})
		name := strings.ToLower(entry["name"].(string))
		if _, ok := r.headersMap[name]; !ok {
			r.headersMap[name] = make(map[string]bool)
		}
		r.headersMap[name][entry["value"].(string)] = true
	}
	return r
}
