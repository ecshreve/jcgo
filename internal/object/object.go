// Package object provides the different implementations of the Object interface
// representing the different possible JSON input types.
package object

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

// FromInterface returns an Object for the given input interface{}, or an
// error if the interface is of an invalid type.
func FromInterface(prefix string, input interface{}) (Object, error) {
	switch vv := input.(type) {
	case nil:
		return NewStringObj(prefix, ""), nil
	case string:
		return NewStringObj(prefix, vv), nil
	case bool:
		return NewBoolObj(prefix, vv), nil
	case float64:
		return NewFloatObj(prefix, vv), nil
	case map[string]interface{}:
		obj, err := NewMapObj(prefix, vv)
		if err != nil {
			return nil, oops.Wrapf(err, "unable to create MapObj for interface: %+v", vv)
		}
		return obj, nil
	case []interface{}:
		obj, err := NewSliceObj(prefix, vv)
		if err != nil {
			return nil, oops.Wrapf(err, "unable to create SliceObj for interface: %+v", vv)
		}
		return obj, nil
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
		{o.Prefix},
		{o.Val},
	}, nil
}

// BoolObj implements the Object interface for a bool value.
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
		{o.Prefix},
		{strconv.FormatBool(o.Val)},
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
		{o.Prefix},
		{stringVal},
	}, nil
}

// MapObj implements the Object interface for a JSON map.
type MapObj struct {
	Prefix string
	Val    []Object
}

// NewMapObj returns a MapObj for the given input map.
func NewMapObj(prefix string, input map[string]interface{}) (*MapObj, error) {
	var vals []Object
	connector := ""

	if prefix != "" {
		connector = "_"
	}

	for k, v := range input {
		newPrefix := fmt.Sprintf("%s%s%s", prefix, connector, k)
		obj, err := FromInterface(newPrefix, v)
		if err != nil {
			return nil, oops.Wrapf(err, "unable to create MapObj")
		}
		vals = append(vals, obj)
	}

	sort.Sort(ByPrefix(vals))

	return &MapObj{
		Prefix: prefix,
		Val:    vals,
	}, nil
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

		// If this is the first item we've parsed then we can initialize ret
		// to that value and skip to parsing the next item.
		if ret == nil {
			ret = parsed
			continue
		}

		// Update the first row in ret with the new keys from the parsed value.
		ret[0] = append(ret[0], parsed[0]...)

		// If the parsed item is a simple scalar value, then set it for each
		// existing row in ret and skip to parsing the next item.
		if len(parsed) == 2 {
			for i := 1; i < len(ret); i++ {
				ret[i] = append(ret[i], parsed[1]...)
			}
			continue
		}

		// If the parsed item is not a simple scalar then update each existing
		// row, and add additional rows as needed.
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
func NewSliceObj(prefix string, input []interface{}) (*SliceObj, error) {
	var vals []Object

	for _, v := range input {
		obj, err := FromInterface(prefix, v)
		if err != nil {
			return nil, oops.Wrapf(err, "unable to create SliceObj")
		}

		vals = append(vals, obj)
	}

	// Sort the Objects in the SliceObj by their Prefix and return.
	sort.Sort(ByPrefix(vals))
	return &SliceObj{
		Prefix: prefix,
		Val:    vals,
	}, nil
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

		// If this is the first item we've parsed then set the first row in ret
		// to the first row in the parsed result.
		if ret == nil {
			ret = append(ret, parsed[0])
		}

		ret = append(ret, parsed[1:]...)
	}

	return ret, nil
}
