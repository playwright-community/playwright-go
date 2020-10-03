package playwright

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"time"
)

type JSHandle struct {
	ChannelOwner
	preview string
}

func (f *JSHandle) Evaluate(expression string, options ...interface{}) (interface{}, error) {
	var arg interface{}
	forceExpression := false
	if !isFunctionBody(expression) {
		forceExpression = true
	}
	if len(options) == 1 {
		arg = options[0]
	} else if len(options) == 2 {
		arg = options[0]
		forceExpression = options[1].(bool)
	}
	result, err := f.channel.Send("evaluateExpression", map[string]interface{}{
		"expression": expression,
		"isFunction": !forceExpression,
		"arg":        serializeArgument(arg),
	})
	if err != nil {
		return nil, err
	}
	return parseResult(result), nil
}

func (f *JSHandle) EvaluateHandle(expression string, options ...interface{}) (interface{}, error) {
	var arg interface{}
	forceExpression := false
	if !isFunctionBody(expression) {
		forceExpression = true
	}
	if len(options) == 1 {
		arg = options[0]
	} else if len(options) == 2 {
		arg = options[0]
		forceExpression = options[1].(bool)
	}
	result, err := f.channel.Send("evaluateExpressionHandle", map[string]interface{}{
		"expression": expression,
		"isFunction": !forceExpression,
		"arg":        serializeArgument(arg),
	})
	if err != nil {
		return nil, err
	}
	channelOwner := fromChannel(result)
	if channelOwner == nil {
		return nil, nil
	}
	return channelOwner.(*JSHandle), nil
}

func (j *JSHandle) GetProperty(name string) (*JSHandle, error) {
	channel, err := j.channel.Send("getProperty", map[string]interface{}{
		"name": name,
	})
	if err != nil {
		return nil, err
	}
	return fromChannel(channel).(*JSHandle), nil
}

func (j *JSHandle) GetProperties() (map[string]*JSHandle, error) {
	properties, err := j.channel.Send("getPropertyList")
	if err != nil {
		return nil, err
	}
	propertiesMap := make(map[string]*JSHandle)
	for _, property := range properties.([]interface{}) {
		item := property.(map[string]interface{})
		propertiesMap[item["name"].(string)] = fromChannel(item["value"]).(*JSHandle)
	}
	return propertiesMap, nil
}

func (j *JSHandle) AsElement() *ElementHandle {
	return nil
}

func (j *JSHandle) Dispose() error {
	_, err := j.channel.Send("dispose")
	return err
}

func (j *JSHandle) String() string {
	return j.preview
}

func (j *JSHandle) JSONValue() (interface{}, error) {
	v, err := j.channel.Send("jsonValue")
	if err != nil {
		return nil, err
	}
	return parseResult(v), nil
}

func parseValue(result interface{}) interface{} {
	vMap := result.(map[string]interface{})
	if v, ok := vMap["n"]; ok {
		if math.Ceil(v.(float64))-v.(float64) == 0 {
			return int(v.(float64))
		}
		return v.(float64)
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
			return -0
		}
	}
	if v, ok := vMap["d"]; ok {
		t, _ := time.Parse(time.RFC3339, v.(string))
		return t
	}
	if v, ok := vMap["a"]; ok {
		aV := v.([]interface{})
		for i := range aV {
			aV[i] = parseValue(aV[i])
		}
		return aV
	}
	if v, ok := vMap["o"]; ok {
		aV := v.([]interface{})
		out := map[string]interface{}{}
		for key := range aV {
			entry := aV[key].(map[string]interface{})
			out[entry["k"].(string)] = parseValue(entry["v"])
		}
		return out
	}
	panic(fmt.Errorf("Unexpected value: %v", vMap))
}

func serializeValue(value interface{}, handles *[]*Channel, depth int) interface{} {
	if handle, ok := value.(*ElementHandle); ok {
		h := len(*handles)
		*handles = append(*handles, handle.channel)
		return map[string]interface{}{
			"h": h,
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
	refV := reflect.ValueOf(value)
	if refV.Kind() == reflect.Float32 || refV.Kind() == reflect.Float64 {
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
		if floatV == -0 {
			return map[string]interface{}{
				"v": "-0",
			}
		}
		if math.IsNaN(floatV) {
			return map[string]interface{}{
				"v": "NaN",
			}
		}
	}
	if refV.Kind() == reflect.Slice {
		aV := value.([]interface{})
		for i := range aV {
			aV[i] = serializeValue(aV[i], handles, depth+1)
		}
		return aV
	}
	if refV.Kind() == reflect.Map {
		out := []interface{}{}
		vM := value.(map[string]interface{})
		for key := range vM {
			out = append(out, map[string]interface{}{
				"k": key,
				"v": serializeValue(vM[key], handles, depth+1),
			})
		}
		return map[string]interface{}{
			"o": out,
		}
	}
	switch v := value.(type) {
	case time.Time:
		return map[string]interface{}{
			"d": v.Format(time.RFC3339) + "Z",
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
	return map[string]interface{}{
		"v": "undefined",
	}
}

func parseResult(result interface{}) interface{} {
	return parseValue(result)
}

func serializeArgument(arg interface{}) interface{} {
	handles := []*Channel{}
	value := serializeValue(arg, &handles, 0)
	return map[string]interface{}{
		"value":   value,
		"handles": handles,
	}
}

func newJSHandle(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *JSHandle {
	bt := &JSHandle{
		preview: initializer["preview"].(string),
	}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	bt.channel.On("previewUpdated", func(ev map[string]interface{}) {
		bt.preview = ev["preview"].(string)
	})
	return bt
}
