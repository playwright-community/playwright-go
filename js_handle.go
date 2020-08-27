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
		return int(v.(float64))
	}
	if v, ok := vMap["s"]; ok {
		return v
	}
	if v, ok := vMap["b"]; ok {
		return v
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

func serializeValue(value interface{}, handles *[]*JSHandle, depth int) interface{} {
	if handle, ok := value.(*JSHandle); ok {
		h := len(*handles)
		*handles = append(*handles, handle)
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
		vM := value.(map[string]interface{})
		for key := range vM {
			vM[key] = map[string]interface{}{
				"k": key,
				"v": serializeValue(vM[key], handles, depth+1),
			}
		}
		return vM
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
	handles := []*JSHandle{}
	value := serializeValue(arg, &handles, 0)
	return map[string]interface{}{
		"value":   value,
		"handles": handles,
	}
}

func newJSHandle(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *JSHandle {
	bt := &JSHandle{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
