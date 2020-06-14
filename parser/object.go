package parser

import (
	"fmt"
	"sort"
)

// Object is representation of a JSON object.
type Object interface {
	getPrefix() string
	parse() [][]string
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
	case string:
		return NewStringObj(prefix, vv)
	default:
		return nil
	}
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

func (o MapObj) parse() [][]string {
	return nil
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

func (o StringObj) parse() [][]string {
	return [][]string{[]string{o.Prefix}, []string{o.Val}}
}
