package object

import (
	"sort"

	"github.com/samsarahq/go/oops"
)

// ArrayObj implements the Object interface for a JSON array.
type ArrayObj struct {
	*Prefix
	Val []Object
}

// NewArrayObj returns a ArrayObj for the given input slice.
func NewArrayObj(prefix string, input []interface{}) (*ArrayObj, error) {
	var vals []Object

	for _, v := range input {
		obj, err := FromInterface(prefix, v)
		if err != nil {
			return nil, oops.Wrapf(err, "unable to create ArrayObj")
		}

		vals = append(vals, obj)
	}

	// Sort the Objects in the ArrayObj by their Prefix and return.
	sort.Sort(ByPrefix(vals))
	return &ArrayObj{
		NewPrefix(prefix),
		vals,
	}, nil
}

// Parse returns the 2d slice of strings for the given ArrayObj.
func (o ArrayObj) Parse() ([][]string, error) {
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
