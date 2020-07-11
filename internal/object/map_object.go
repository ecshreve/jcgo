package object

import (
	"fmt"
	"sort"

	"github.com/samsarahq/go/oops"
)

// MapObj implements the Object interface for a JSON map.
type MapObj struct {
	*Prefix
	SortedKeys []string
	Val        map[string]Object
}

// NewMapObj returns a MapObj for the given input map.
func NewMapObj(prefix string, input map[string]interface{}) (*MapObj, error) {
	var keys []string
	vals := make(map[string]Object)

	for k, v := range input {
		newPrefix := fmt.Sprintf("%s_%s", prefix, k)
		obj, err := FromInterface(newPrefix, v)
		if err != nil {
			return nil, oops.Wrapf(err, "unable to create MapObj")
		}
		vals[k] = obj
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return &MapObj{
		NewPrefix(prefix),
		keys,
		vals,
	}, nil
}

// Parse returns the 2d slice of strings for the given MapObj.
func (o MapObj) Parse() ([][]string, error) {
	var ret [][]string

	for _, key := range o.SortedKeys {
		item := o.Val[key]
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
