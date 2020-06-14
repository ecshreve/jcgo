package parser

import (
	"fmt"
	"sort"
)

type Object interface {
	GetPrefix() string
	Parse() [][]string
}

type ByPrefix []Object

func (objs ByPrefix) Len() int           { return len(objs) }
func (objs ByPrefix) Less(i, j int) bool { return objs[i].GetPrefix() < objs[j].GetPrefix() }
func (objs ByPrefix) Swap(i, j int)      { objs[i], objs[j] = objs[j], objs[i] }

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

type MapObj struct {
	Prefix string
	Val    []Object
}

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

func (o MapObj) GetPrefix() string {
	return o.Prefix
}

func (o MapObj) Parse() [][]string {
	return nil
}

type StringObj struct {
	Prefix string
	Val    string
}

func NewStringObj(prefix string, input string) *StringObj {
	return &StringObj{
		Prefix: prefix,
		Val:    input,
	}
}

func (o StringObj) GetPrefix() string {
	return o.Prefix
}

func (o StringObj) Parse() [][]string {
	return [][]string{[]string{o.Prefix}, []string{o.Val}}
}
