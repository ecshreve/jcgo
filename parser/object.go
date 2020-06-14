package parser

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/samsarahq/go/oops"
)

// Object is representation of a JSON object.
type Object interface {
	getPrefix() string
	Parse() ([][]string, error)
}

// ByPrefix implements the sort.Interface on the Prefix field.
type ByPrefix []Object

func (objs ByPrefix) Len() int           { return len(objs) }
func (objs ByPrefix) Less(i, j int) bool { return objs[i].getPrefix() < objs[j].getPrefix() }
func (objs ByPrefix) Swap(i, j int)      { objs[i], objs[j] = objs[j], objs[i] }

// ObjectFromInterface returns an Object for the given input interface{}.
func ObjectFromInterface(prefix string, input interface{}) (Object, error) {
	switch vv := input.(type) {
	case string:
		return NewStringObj(prefix, vv), nil
	case bool:
		return NewBoolObj(prefix, vv), nil
	case float64:
		return NewFloatObj(prefix, vv), nil
	case map[string]interface{}:
		return NewMapObj(prefix, vv), nil
	case []interface{}:
		return NewSliceObj(prefix, vv), nil
	default:
		return nil, oops.Errorf("unable to create Object from interface: %+v", input)
	}
}

// StringObj implements the Object interface for a string value.
type StringObj struct {
	Prefix string
	Val    string
}

// NewStringObj returns a StringObj for the given input string.
func NewStringObj(prefix string, input string) *StringObj {
	return &StringObj{
		Prefix: prefix,
		Val:    input,
	}
}

func (o StringObj) getPrefix() string {
	return o.Prefix
}

// Parse returns the 2d slice of strings for the given StringObj.
func (o StringObj) Parse() ([][]string, error) {
	return [][]string{
		[]string{o.Prefix},
		[]string{o.Val},
	}, nil
}

// BoolObj implements the Object interface for a float value.
type BoolObj struct {
	Prefix string
	Val    bool
}

// NewBoolObj returns a BoolObj for the given input bool.
func NewBoolObj(prefix string, input bool) *BoolObj {
	return &BoolObj{
		Prefix: prefix,
		Val:    input,
	}
}

func (o BoolObj) getPrefix() string {
	return o.Prefix
}

// Parse returns the 2d slice of strings for the given BoolObj.
func (o BoolObj) Parse() ([][]string, error) {
	return [][]string{
		[]string{o.Prefix},
		[]string{strconv.FormatBool(o.Val)},
	}, nil
}

// FloatObj implements the Object interface for a float value.
type FloatObj struct {
	Prefix string
	Val    float64
}

// NewFloatObj returns a FloatObj for the given input float.
func NewFloatObj(prefix string, input float64) *FloatObj {
	return &FloatObj{
		Prefix: prefix,
		Val:    input,
	}
}

func (o FloatObj) getPrefix() string {
	return o.Prefix
}

// Parse returns the 2d slice of strings for the given FloatObj.
func (o FloatObj) Parse() ([][]string, error) {
	floatVal := o.Val

	var stringVal string
	if float64(int64(floatVal)) == floatVal {
		stringVal = strconv.FormatInt(int64(floatVal), 10)
	} else {
		stringVal = strconv.FormatFloat(floatVal, 'f', -1, 64)
	}

	return [][]string{
		[]string{o.Prefix},
		[]string{stringVal},
	}, nil
}

// MapObj implements the Object interface for a JSON map.
type MapObj struct {
	Prefix string
	Val    []Object
}

// NewMapObj returns a MapObj for the given input map.
func NewMapObj(prefix string, input map[string]interface{}) *MapObj {
	var vals []Object
	connector := ""

	if prefix != "" {
		connector = "_"
	}

	for k, v := range input {
		newPrefix := fmt.Sprintf("%s%s%s", prefix, connector, k)
		obj, _ := ObjectFromInterface(newPrefix, v)
		vals = append(vals, obj)
	}

	sort.Sort(ByPrefix(vals))

	return &MapObj{
		Prefix: prefix,
		Val:    vals,
	}
}

func (o MapObj) getPrefix() string {
	return o.Prefix
}

// Parse returns the 2d slice of strings for the given MapObj.
func (o MapObj) Parse() ([][]string, error) {
	var ret [][]string

	for _, item := range o.Val {
		parsed, err := item.Parse()
		if err != nil {
			return nil, oops.Wrapf(err, "unable to parse item: %+v", item)
		}

		if ret == nil {
			ret = parsed
			continue
		}

		ret[0] = append(ret[0], parsed[0]...)
		if len(parsed) == 2 {
			for i := 1; i < len(ret); i++ {
				ret[i] = append(ret[i], parsed[1]...)
			}
			continue
		}

		lastRow := ret[len(ret)-1]
		for i := 1; i < len(parsed); i++ {
			if i == len(ret) {
				ret = append(ret, lastRow)
			}
			ret[i] = append(ret[i], parsed[i]...)
		}
	}

	return ret, nil
}

// SliceObj implements the Object interface for a JSON array.
type SliceObj struct {
	Prefix string
	Val    []Object
}

// NewSliceObj returns a SliceObj for the given input slice.
func NewSliceObj(prefix string, input []interface{}) *SliceObj {
	var vals []Object

	for _, v := range input {
		obj, _ := ObjectFromInterface(prefix, v)
		vals = append(vals, obj)
	}

	sort.Sort(ByPrefix(vals))

	return &SliceObj{
		Prefix: prefix,
		Val:    vals,
	}
}

func (o SliceObj) getPrefix() string {
	return o.Prefix
}

// Parse returns the 2d slice of strings for the given SliceObj.
func (o SliceObj) Parse() ([][]string, error) {
	var ret [][]string

	for _, item := range o.Val {
		parsed, err := item.Parse()
		if err != nil {
			return nil, oops.Wrapf(err, "unable to parse item: %+v", item)
		}

		if len(ret) == 0 {
			ret = append(ret, parsed[0])
		}

		ret = append(ret, parsed[1:]...)
	}

	return ret, nil
}
