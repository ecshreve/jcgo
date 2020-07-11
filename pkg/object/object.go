// Package object provides implementations of the Object interface for the
// possible JSON input value types.
package object

import (
	"strconv"

	"github.com/samsarahq/go/oops"
)

// Prefix is just a string, we redefine it so we can implement the getPrefix
// method of the Object interface once rather than for each new Object type.
type Prefix string

func (p Prefix) getPrefix() string {
	return string(p)
}

// NewPrefix returns a pointer to a Prefix for the given string.
func NewPrefix(p string) *Prefix {
	if len(p) > 0 && p[0] == '_' {
		p = p[1:]
	}
	pp := Prefix(p)
	return &pp
}

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

// FromInterface returns the Object for the given input interface and returns an
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
		return NewNumberObj(prefix, vv), nil
	case map[string]interface{}:
		obj, err := NewMapObj(prefix, vv)
		if err != nil {
			return nil, oops.Wrapf(err, "unable to create MapObj for interface: %+v", vv)
		}
		return obj, nil
	case []interface{}:
		obj, err := NewArrayObj(prefix, vv)
		if err != nil {
			return nil, oops.Wrapf(err, "unable to create ArrayObj for interface: %+v", vv)
		}
		return obj, nil
	default:
		return nil, oops.Errorf("unable to create Object from interface: %+v", input)
	}
}

// StringObj implements the Object interface for a string value.
type StringObj struct {
	*Prefix
	Val string
}

// NewStringObj returns a StringObj for the given input string.
func NewStringObj(prefix string, input string) *StringObj {
	return &StringObj{
		NewPrefix(prefix),
		input,
	}
}

// Parse returns the 2d slice of strings for the given StringObj.
func (o StringObj) Parse() ([][]string, error) {
	return [][]string{
		{string(*o.Prefix)},
		{o.Val},
	}, nil
}

// BoolObj implements the Object interface for a bool value.
type BoolObj struct {
	*Prefix
	Val bool
}

// NewBoolObj returns a BoolObj for the given input bool.
func NewBoolObj(prefix string, input bool) *BoolObj {
	return &BoolObj{
		NewPrefix(prefix),
		input,
	}
}

// Parse returns the 2d slice of strings for the given BoolObj.
func (o BoolObj) Parse() ([][]string, error) {
	return [][]string{
		{string(*o.Prefix)},
		{strconv.FormatBool(o.Val)},
	}, nil
}

// NumberObj implements the Object interface for a numeric value.
type NumberObj struct {
	*Prefix
	Val float64
}

// NewNumberObj returns a NumberObj for the given input float.
func NewNumberObj(prefix string, input float64) *NumberObj {
	return &NumberObj{
		NewPrefix(prefix),
		input,
	}
}

// Parse returns the 2d slice of strings for the given FloatObj.
//
// We do a check to see if the value is actually an integer, and treat if so.
// This check is necessary because the results of unmarshalling a JSON byte
// array into a map[string]interface{} treats all numeric values as floats. We
// want the string representation that we eventually write to a CSV file to have
// numeric values appear accurately if they're integers. i.e. "345" not "345.0".
func (o NumberObj) Parse() ([][]string, error) {
	floatVal := o.Val

	var stringVal string
	if float64(int64(floatVal)) == floatVal {
		stringVal = strconv.FormatInt(int64(floatVal), 10)
	} else {
		stringVal = strconv.FormatFloat(floatVal, 'f', -1, 64)
	}

	return [][]string{
		{string(*o.Prefix)},
		{stringVal},
	}, nil
}
