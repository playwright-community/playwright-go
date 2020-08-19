package playwright

import (
	"reflect"
	"strings"
)

// transformOptions handles the parameter data transformation
func transformOptions(options ...interface{}) map[string]interface{} {
	var base map[string]interface{}
	var option interface{}
	// Case 1: No options are given
	if len(options) == 0 {
		return make(map[string]interface{})
	}
	if len(options) == 1 {
		// Case 2: a single value (either struct or map) is given.
		base = make(map[string]interface{})
		option = options[0]
	} else if len(options) == 2 {
		// Case 3: two values are given. The first one needs to be a map and the
		// second one can be a struct or map. It will be then get merged into the first
		// base map.
		base = options[0].(map[string]interface{})
		val := reflect.ValueOf(options[1])
		if val.Kind() == reflect.Slice {
			if val.Len() == 0 {
				return base
			}
			option = val.Index(0).Interface()
		}
	}
	if option == nil {
		return base
	}

	v := reflect.ValueOf(option)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	typ := v.Type()
	if v.Kind() == reflect.Struct {
		// Merge into the base map by the JSON struct tag
		for i := 0; i < v.NumField(); i++ {
			fi := typ.Field(i)
			// Skip the values when the field is a pointer (like *string) and nil.
			if !(fi.Type.Kind() == reflect.Ptr && v.Field(i).IsNil()) {
				// We use the JSON struct fields for getting the original names
				// out of the field.
				tagv := fi.Tag.Get("json")
				key := strings.Split(tagv, ",")[0]
				if key == "" {
					key = fi.Name
				}
				base[key] = v.Field(i).Interface()
			}
		}
	} else if v.Kind() == reflect.Map {
		// Merge into the base map
		for key, value := range option.(map[string]interface{}) {
			base[key] = value
		}
	}
	return base
}
