package parser

import (
	"fmt"

	"github.com/kr/pretty"
)

// Transform returns an Object for the given input map. It's meant to be called
// on the root map[string]interface{} that comes from Unmarshalling a JSON file
// into a map[string]interface{}.
func Transform(input map[string]interface{}) Object {
	o := ObjectFromInterface("", input)
	return o
}

func Parse(o Object) {
	parsed, err := o.Parse()
	if err != nil {
		fmt.Print(err)
	}
	pretty.Print(parsed)
}
