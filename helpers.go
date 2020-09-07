package playwright

import (
	"reflect"
	"regexp"
	"strings"
	"sync"

	"github.com/danwakefield/fnmatch"
)

type (
	routeHandler = func(*Route, *Request)
)

func skipFieldSerialization(val reflect.Value) bool {
	typ := val.Type()
	return (typ.Kind() == reflect.Ptr ||
		typ.Kind() == reflect.Interface ||
		typ.Kind() == reflect.Map ||
		typ.Kind() == reflect.Slice) && val.IsNil()
}

func transformStructValues(in interface{}) interface{} {
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if _, ok := in.(*Channel); ok {
		return in
	}
	if v.Kind() == reflect.Map || v.Kind() == reflect.Struct {
		return transformStructIntoMapIfNeeded(in)
	}
	if v.Kind() == reflect.Slice {
		outSlice := []interface{}{}
		for i := 0; i < v.Len(); i++ {
			if !skipFieldSerialization(v.Index(i)) {
				outSlice = append(outSlice, transformStructValues(v.Index(i).Interface()))
			}
		}
		return outSlice
	}
	return in
}

func transformStructIntoMapIfNeeded(inStruct interface{}) map[string]interface{} {
	out := make(map[string]interface{})
	v := reflect.ValueOf(inStruct)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	typ := v.Type()
	if v.Kind() == reflect.Struct {
		// Merge into the base map by the JSON struct tag
		for i := 0; i < v.NumField(); i++ {
			fi := typ.Field(i)
			// Skip the values when the field is a pointer (like *string) and nil.
			if !skipFieldSerialization(v.Field(i)) {
				// We use the JSON struct fields for getting the original names
				// out of the field.
				tagv := fi.Tag.Get("json")
				key := strings.Split(tagv, ",")[0]
				if key == "" {
					key = fi.Name
				}
				out[key] = transformStructValues(v.Field(i).Interface())
			}
		}
	} else if v.Kind() == reflect.Map {
		// Merge into the base map
		for _, key := range v.MapKeys() {
			if !skipFieldSerialization(v.MapIndex(key)) {
				out[key.String()] = transformStructValues(v.MapIndex(key).Interface())
			}
		}
	}
	return out
}

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
		base = transformStructIntoMapIfNeeded(options[0])
		option = options[1]
	}
	v := reflect.ValueOf(option)
	if v.Kind() == reflect.Slice {
		if v.Len() == 0 {
			return base
		}
		option = v.Index(0).Interface()
	}

	if option == nil {
		return base
	}
	v = reflect.ValueOf(option)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	optionMap := transformStructIntoMapIfNeeded(v.Interface())
	for key, value := range optionMap {
		base[key] = value
	}
	return base
}

func remapValue(inMapValue reflect.Value, outStructValue reflect.Value) {
	switch outStructValue.Type().Kind() {
	case reflect.Bool:
		outStructValue.SetBool(inMapValue.Bool())
	case reflect.String:
		outStructValue.SetString(inMapValue.String())
	case reflect.Float64:
		outStructValue.SetFloat(inMapValue.Float())
	case reflect.Int:
		outStructValue.SetInt(int64(inMapValue.Float()))
	case reflect.Map:
		for _, key := range inMapValue.MapKeys() {
			remapMapToStruct(inMapValue.MapIndex(key).Interface(), inMapValue.Interface())
		}
	case reflect.Struct:
		structTyp := outStructValue.Type()
		for i := 0; i < outStructValue.NumField(); i++ {
			fi := structTyp.Field(i)
			key := strings.Split(fi.Tag.Get("json"), ",")[0]
			structField := outStructValue.Field(i)
			structFieldDeref := outStructValue.Field(i)
			if structField.Type().Kind() == reflect.Ptr {
				structField.Set(reflect.New(structField.Type().Elem()))
				structFieldDeref = structField.Elem()
			}
			for _, e := range inMapValue.MapKeys() {
				if key == e.String() {
					value := inMapValue.MapIndex(e)
					remapValue(value.Elem(), structFieldDeref)
				}
			}
		}
	default:
		panic(inMapValue.Interface())
	}
}

func remapMapToStruct(inputMap interface{}, outStruct interface{}) {
	remapValue(reflect.ValueOf(inputMap), reflect.ValueOf(outStruct).Elem())
}

func isFunctionBody(expression string) bool {
	expression = strings.TrimSpace(expression)
	return strings.HasPrefix(expression, "function") ||
		strings.HasPrefix(expression, "async ") ||
		strings.Contains(expression, "=> ")
}

type urlMatcher struct {
	urlOrPredicate interface{}
}

func newURLMatcher(urlOrPredicate interface{}) *urlMatcher {
	return &urlMatcher{
		urlOrPredicate: urlOrPredicate,
	}
}

func (u *urlMatcher) Match(url string) bool {
	switch v := u.urlOrPredicate.(type) {
	case *regexp.Regexp:
		return v.MatchString(url)
	case string:
		return fnmatch.Match(v, url, 0)
	}
	if reflect.TypeOf(u.urlOrPredicate).Kind() == reflect.Func {
		function := reflect.ValueOf(u.urlOrPredicate)
		result := function.Call([]reflect.Value{reflect.ValueOf(url)})
		return result[0].Bool()
	}
	panic(u.urlOrPredicate)
}

type routeHandlerEntry struct {
	matcher *urlMatcher
	handler routeHandler
}

func newRouteHandlerEntry(matcher *urlMatcher, handler routeHandler) *routeHandlerEntry {
	return &routeHandlerEntry{
		matcher: matcher,
		handler: handler,
	}
}

type safeStringSet struct {
	sync.Mutex
	v []string
}

func (s *safeStringSet) Has(expected string) bool {
	s.Lock()
	defer s.Unlock()
	for _, v := range s.v {
		if v == expected {
			return true
		}
	}
	return false
}

func (s *safeStringSet) Add(v string) {
	if s.Has(v) {
		return
	}
	s.Lock()
	s.v = append(s.v, v)
	s.Unlock()
}

func (s *safeStringSet) Remove(remove string) {
	s.Lock()
	defer s.Unlock()
	newSlice := make([]string, 0)
	for _, v := range s.v {
		if v != remove {
			newSlice = append(newSlice, v)
		}
	}
	if len(s.v) != len(newSlice) {
		s.v = newSlice
	}
}

func newSafeStringSet(v []string) *safeStringSet {
	return &safeStringSet{
		v: v,
	}
}

const DEFAULT_TIMEOUT = 30 * 1000

type timeoutSettings struct {
	parent            *timeoutSettings
	timeout           int
	navigationTimeout int
}

func (t *timeoutSettings) SetTimeout(timeout int) {
	t.timeout = timeout
}

func (t *timeoutSettings) Timeout() int {
	if t.timeout != 0 {
		return t.timeout
	}
	if t.parent != nil {
		return t.parent.Timeout()
	}
	return DEFAULT_TIMEOUT
}

func (t *timeoutSettings) SetNavigationTimeout(navigationTimeout int) {
	t.navigationTimeout = navigationTimeout
}

func (t *timeoutSettings) NavigationTimeout() int {
	if t.navigationTimeout != 0 {
		return t.navigationTimeout
	}
	if t.parent != nil {
		return t.parent.NavigationTimeout()
	}
	return DEFAULT_TIMEOUT
}

func newTimeoutSettings(parent *timeoutSettings) *timeoutSettings {
	return &timeoutSettings{
		parent:            parent,
		timeout:           DEFAULT_TIMEOUT,
		navigationTimeout: DEFAULT_TIMEOUT,
	}
}
