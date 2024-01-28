package playwright

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"net/url"
	"reflect"
	"runtime/debug"
	"strings"
	"time"
)

type jsHandleImpl struct {
	channelOwner
	preview string
}

func (j *jsHandleImpl) Evaluate(expression string, options ...interface{}) (interface{}, error) {
	var arg interface{}
	if len(options) == 1 {
		arg = options[0]
	}
	result, err := j.channel.Send("evaluateExpression", map[string]interface{}{
		"expression": expression,
		"arg":        serializeArgument(arg),
	})
	if err != nil {
		return nil, err
	}
	return parseResult(result), nil
}

func (j *jsHandleImpl) EvaluateHandle(expression string, options ...interface{}) (JSHandle, error) {
	var arg interface{}
	if len(options) == 1 {
		arg = options[0]
	}
	result, err := j.channel.Send("evaluateExpressionHandle", map[string]interface{}{
		"expression": expression,
		"arg":        serializeArgument(arg),
	})
	if err != nil {
		return nil, err
	}
	channelOwner := fromChannel(result)
	if channelOwner == nil {
		return nil, nil
	}
	return channelOwner.(*jsHandleImpl), nil
}

func (j *jsHandleImpl) GetProperty(name string) (JSHandle, error) {
	channel, err := j.channel.Send("getProperty", map[string]interface{}{
		"name": name,
	})
	if err != nil {
		return nil, err
	}
	return fromChannel(channel).(*jsHandleImpl), nil
}

func (j *jsHandleImpl) GetProperties() (map[string]JSHandle, error) {
	properties, err := j.channel.Send("getPropertyList")
	if err != nil {
		return nil, err
	}
	propertiesMap := make(map[string]JSHandle)
	for _, property := range properties.([]interface{}) {
		item := property.(map[string]interface{})
		propertiesMap[item["name"].(string)] = fromChannel(item["value"]).(*jsHandleImpl)
	}
	return propertiesMap, nil
}

func (j *jsHandleImpl) AsElement() ElementHandle {
	return nil
}

func (j *jsHandleImpl) Dispose() error {
	_, err := j.channel.Send("dispose")
	return err
}

func (j *jsHandleImpl) String() string {
	return j.preview
}

func (j *jsHandleImpl) JSONValue() (interface{}, error) {
	v, err := j.channel.Send("jsonValue")
	if err != nil {
		return nil, err
	}
	return parseResult(v), nil
}

func parseValue(result interface{}, refs map[float64]interface{}) interface{} {
	vMap, ok := result.(map[string]interface{})
	if !ok {
		return result
	}
	if v, ok := vMap["n"]; ok {
		if math.Ceil(v.(float64))-v.(float64) == 0 {
			return int(v.(float64))
		}
		return v.(float64)
	}

	if v, ok := vMap["u"]; ok {
		u, _ := url.Parse(v.(string))
		return u
	}

	if v, ok := vMap["bi"]; ok {
		n := new(big.Int)
		n.SetString(v.(string), 0)
		return n
	}

	if v, ok := vMap["ref"]; ok {
		if vV, ok := refs[v.(float64)]; ok {
			return vV
		}
		return nil
	}

	if v, ok := vMap["s"]; ok {
		return v.(string)
	}
	if v, ok := vMap["b"]; ok {
		return v.(bool)
	}
	if v, ok := vMap["v"]; ok {
		if v == "undefined" || v == "null" {
			return nil
		}
		if v == "NaN" {
			return math.NaN()
		}
		if v == "Infinity" {
			return math.Inf(1)
		}
		if v == "-Infinity" {
			return math.Inf(-1)
		}
		if v == "-0" {
			return math.Copysign(0, -1)
		}
		return v
	}
	if v, ok := vMap["d"]; ok {
		t, _ := time.Parse(time.RFC3339Nano, v.(string))
		return t
	}
	if v, ok := vMap["a"]; ok {
		aV := v.([]interface{})
		refs[vMap["id"].(float64)] = aV
		for i := range aV {
			aV[i] = parseValue(aV[i], refs)
		}
		return aV
	}
	if v, ok := vMap["o"]; ok {
		aV := v.([]interface{})
		out := map[string]interface{}{}
		refs[vMap["id"].(float64)] = out
		for key := range aV {
			entry := aV[key].(map[string]interface{})
			out[entry["k"].(string)] = parseValue(entry["v"], refs)
		}
		return out
	}
	panic(fmt.Errorf("Unexpected value: %v", vMap))
}

func serializeValue(value interface{}, handles *[]*channel, depth int) interface{} {
	if handle, ok := value.(*elementHandleImpl); ok {
		h := len(*handles)
		*handles = append(*handles, handle.channel)
		return map[string]interface{}{
			"h": h,
		}
	}
	if handle, ok := value.(*jsHandleImpl); ok {
		h := len(*handles)
		*handles = append(*handles, handle.channel)
		return map[string]interface{}{
			"h": h,
		}
	}
	if u, ok := value.(*url.URL); ok {
		return map[string]interface{}{
			"u": u.String(),
		}
	}

	if depth > 100 {
		panic(errors.New("Maximum argument depth exceeded"))
	}
	if value == nil {
		return map[string]interface{}{
			"v": "undefined",
		}
	}
	if n, ok := value.(*big.Int); ok {
		return map[string]interface{}{
			"bi": n.String(),
		}
	}

	switch v := value.(type) {
	case time.Time:
		return map[string]interface{}{
			"d": v.Format(time.RFC3339Nano),
		}
	case int:
		return map[string]interface{}{
			"n": v,
		}
	case string:
		return map[string]interface{}{
			"s": v,
		}
	case bool:
		return map[string]interface{}{
			"b": v,
		}
	}

	refV := reflect.ValueOf(value)

	switch refV.Kind() {
	case reflect.Float32, reflect.Float64:
		floatV := refV.Float()
		if math.IsInf(floatV, 1) {
			return map[string]interface{}{
				"v": "Infinity",
			}
		}
		if math.IsInf(floatV, -1) {
			return map[string]interface{}{
				"v": "-Infinity",
			}
		}
		// https://github.com/golang/go/issues/2196
		if floatV == math.Copysign(0, -1) {
			return map[string]interface{}{
				"v": "-0",
			}
		}
		if math.IsNaN(floatV) {
			return map[string]interface{}{
				"v": "NaN",
			}
		}
		return map[string]interface{}{
			"n": floatV,
		}
	case reflect.Slice:
		aV := make([]interface{}, refV.Len())
		for i := 0; i < refV.Len(); i++ {
			aV[i] = serializeValue(refV.Index(i).Interface(), handles, depth+1)
		}
		return map[string]interface{}{
			"a": aV,
		}
	case reflect.Map:
		out := []interface{}{}
		vM := value.(map[string]interface{})
		for key := range vM {
			v := serializeValue(vM[key], handles, depth+1)
			// had key, so convert "undefined" to "null"
			if reflect.DeepEqual(v, map[string]interface{}{
				"v": "undefined",
			}) {
				v = map[string]interface{}{
					"v": "null",
				}
			}
			out = append(out, map[string]interface{}{
				"k": key,
				"v": v,
			})
		}
		return map[string]interface{}{
			"o": out,
		}
	}
	return map[string]interface{}{
		"v": "undefined",
	}
}

func parseResult(result interface{}) interface{} {
	return parseValue(result, map[float64]interface{}{})
}

func serializeArgument(arg interface{}) interface{} {
	handles := []*channel{}
	value := serializeValue(arg, &handles, 0)
	return map[string]interface{}{
		"value":   value,
		"handles": handles,
	}
}

func serializeError(err error) map[string]interface{} {
	stack := strings.Split(string(debug.Stack()), "\n")
	return map[string]interface{}{
		"error": &Error{
			Name:    "Playwright for Go Error",
			Message: err.Error(),
			Stack:   strings.Join(stack[:len(stack)-5], "\n"),
		},
	}
}

func newJSHandle(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *jsHandleImpl {
	bt := &jsHandleImpl{
		preview: initializer["preview"].(string),
	}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	bt.channel.On("previewUpdated", func(ev map[string]interface{}) {
		bt.preview = ev["preview"].(string)
	})
	return bt
}
