# JSON

[JSON (JavaScript Object Notation)](https://www.json.org/json-en.html) is a lightweight data-interchange format.

## My Summary of the JSON Specification

JSON is built on two structures:

- `object`: collection of name/value pairs (in `go` this is a `map`)
- `array`: an ordered list of values (in `go` we use `slices` to access arrays)
- structures can be nested arbitrarily

## JSON and Golang

- [Go Blog - Json and Go](https://blog.golang.org/json)
- [encoding/json package](https://golang.org/pkg/encoding/json/)

### General

- encode and decode using the `Marshal` and `Unmarshal` functions
- JSON objects only support `string` keys
- complex types can't be encoded
- cyclic data structures can't be encoded
- pointers are encoded as the values they point to, or `null` if the pointer is `nil`

### Encoding / Decoding

- JSON can be decoded into and encoded from defined `go` structs
- `Unmarshal` will only decode fields it can match to the `go` type
- can use `json` tags in the struct definition to aid in decoding
- if the structure is unknown, can decode into the `interface{}` type
  - this lets us decode arbitrary json into a `go` structure

## Data Types

### Object

- unordered set of name/value pairs
- begins with `{` and ends with `}`
- each name is followed by a `:`
- name/value pairs are separated by `,`

### Array

- ordered collection of values
- begins with `[` and ends with `]`
- values are separated by `,`

### Value

values can be nested, and can be any of the following:

- a `string` in double quotes
- a `number`
- a `boolean` (true or false)
- `null`
- an `object`
- an `array`

### String

- zero or more unicode chars wrapped in double quotes
- uses backslash escapes

### Number

- can be positive or negative, indicated by a leading `-`
- can be a `float`, using either decimal or exponent notation
- octal and hex formats are not supported

### Whitespace

- spaces, tabs, linefeeds, cariage returns
- can be inserted between any pair of tokens
