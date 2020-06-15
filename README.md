# jcgo

![Go](https://github.com/ecshreve/jcgo/workflows/Go/badge.svg?branch=master)
![Go Report Card](https://goreportcard.com/badge/github.com/ecshreve/jcgo)
![Go Version](https://img.shields.io/github/go-mod/go-version/ecshreve/jcgo)
![Last Commit](https://img.shields.io/github/last-commit/ecshreve/jcgo)
![Codecov](https://img.shields.io/codecov/c/github/ecshreve/jcgo)

## Description

JSON to CSV converter in Golang.

**Note: at this time, this converter doesn't implement the full JSON spec.**

## Usage

`jcgo` reads JSON from the file specified as a command line argument, parses the JSON, and writes the data to a CSV file.

```{bash}
git clone https://github.com/ecshreve/jcgo.git
cd jcgo
make build
bin/jcgo --infile jsontestlocal.json
```

## reference

- [Effective Go](https://golang.org/doc/effective_go.html)
- [Tour of Go](https://tour.golang.org/list)
- [Go by Example](https://gobyexample.com/)
- [Using Go Modules](https://blog.golang.org/using-go-modules)
- [encoding/json](https://golang.org/pkg/encoding/json/)
- [encoding/csv](https://golang.org/pkg/encoding/csv/)
- [JSON specification](https://www.json.org/json-en.html)

## Googling

Here's an abbreviated list of the things I googled while making this (in no particular order).

- json to csv
- arbitrary json to csv
- parse json golang
- parse nested json
- json to csv golang github
- json spec
- graphql json results
- write to csv golang
- golang parse command line args
- golang flag package
- golang use flags in test
- longest common prefix
- golang module
- golang project organization
- golang remove dependencies from module
- go install vs go build
- godoc json
- godoc csv
- golang stringer
- golang interface inheritance
- golang init function
- golang sort custom type
- golang embedded struct
  ...
  and many, many more
