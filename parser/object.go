package parser

import (
	"fmt"
	"sort"
)

// Object is representation of a JSON object.
type Object interface {
	getPrefix() string
	Parse() [][]string
}

// ByPrefix implements the sort.Interface on the Prefix field.
type ByPrefix []Object

func (objs ByPrefix) Len() int           { return len(objs) }
func (objs ByPrefix) Less(i, j int) bool { return objs[i].getPrefix() < objs[j].getPrefix() }
func (objs ByPrefix) Swap(i, j int)      { objs[i], objs[j] = objs[j], objs[i] }

// ObjectFromInterface returns an Object for the given input interface{}.
func ObjectFromInterface(prefix string, input interface{}) Object {
	switch vv := input.(type) {
	case map[string]interface{}:
		return NewMapObj(prefix, vv)
	case []interface{}:
		return NewSliceObj(prefix, vv)
	case string:
		return NewStringObj(prefix, vv)
	default:
		return nil
	}
}

type SliceObj struct {
	Prefix string
	Val    []Object
}

// NewSliceObj returns a SliceObj for the given input slice.
func NewSliceObj(prefix string, input []interface{}) *SliceObj {
	var vals []Object

	for _, v := range input {
		vals = append(vals, ObjectFromInterface(prefix, v))
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

func (o SliceObj) Parse() [][]string {
	var ret [][]string

	for _, item := range o.Val {
		parsed := item.Parse()
		if len(ret) == 0 {
			ret = append(ret, parsed[0])
		}

		ret = append(ret, parsed[1:]...)
	}

	return ret
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
		vals = append(vals, ObjectFromInterface(newPrefix, v))
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
func (o MapObj) Parse() [][]string {
	var ret [][]string

	for _, item := range o.Val {
		parsed := item.Parse()

		if ret == nil {
			ret = parsed
			continue
		}

		ret[0] = append(ret[0], parsed[0]...)

		if len(parsed) == 2 {
			for i := 1; i < len(ret); i++ {
				ret[i] = append(ret[i], parsed[1]...)
			}
		} else {
			lastRow := ret[len(ret)-1]
			for i := 1; i < len(parsed)-1; i++ {
				ret = append(ret, lastRow)
			}
			for i := 1; i < len(ret); i++ {
				ret[i] = append(ret[i], parsed[i]...)
			}
		}
	}

	return ret
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
func (o StringObj) Parse() [][]string {
	return [][]string{
		[]string{o.Prefix},
		[]string{o.Val},
	}
}
