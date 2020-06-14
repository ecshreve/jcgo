package parser

func Transform(input map[string]interface{}) Object {
	o := ObjectFromInterface("", input)
	return o
}

func Parse(o Object) [][]string {
	return o.Parse()
}
