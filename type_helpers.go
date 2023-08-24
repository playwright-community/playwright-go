package playwright

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string {
	return &v
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool {
	return &v
}

// Int is a helper routine that allocates a new int32 value
// to store v and returns a pointer to it.
func Int(v int) *int {
	return &v
}

// Float is a helper routine that allocates a new float64 value
// to store v and returns a pointer to it.
func Float(v float64) *float64 {
	return &v
}

// Null will be used in certain scenarios where a strict nil pointer
// check is not possible
func Null() interface{} {
	return "PW_NULL"
}

// StringSlice is a helper routine that allocates a new StringSlice value
// to store v and returns a pointer to it.
func StringSlice(v ...string) *[]string {
	var o []string
	o = append(o, v...)
	return &o
}

// IntSlice is a helper routine that allocates a new IntSlice value
// to store v and returns a pointer to it.
func IntSlice(v ...int) *[]int {
	var o []int
	o = append(o, v...)
	return &o
}
