# jcgo

[![Go](https://github.com/ecshreve/jcgo/workflows/Go/badge.svg?branch=master)](https://github.com/ecshreve/jcgo/actions)
[![GoDoc](https://godoc.org/github.com/ecshreve/jcgo?status.svg)](https://godoc.org/github.com/ecshreve/jcgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/ecshreve/jcgo)](https://goreportcard.com/report/github.com/ecshreve/jcgo)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ecshreve/jcgo)](https://golang.org/doc/go1.14)
[![codecov](https://codecov.io/gh/ecshreve/jcgo/branch/master/graph/badge.svg)](https://codecov.io/gh/ecshreve/jcgo)

## Description

JSON to CSV converter in Golang.

**Note: at this time, this converter doesn't implement the full JSON spec.** (it doesn't currently handle arrays of scalar values, only arrays of `map[string]interface{}`)

## Usage

`jcgo` reads JSON from the file specified as a command line argument, parses the JSON, and writes the data to a CSV file.

```{bash}
> git clone https://github.com/ecshreve/jcgo.git
> cd jcgo
> make build
> bin/jcgo jsontestlocal.json jsontestlocal.output.csv
2020/07/24 18:39:22 generated csv file: jsontestlocal.output.csv

> cat jsontestlocal.output.csv | column -t -s ,
afterState_arrivedAt  afterState_departureTimeMs  afterState_destinationName  afterState_id  afterState_jobState  beforeState_arrivedAt  beforeState_departureTimeMs  beforeState_destinationName  beforeState_id  beforeState_jobState  changedAtMs    events_eventAt  events_eventType  id
0                     0                           CARI307                     4337769816     4                    0                      0                            CARI307                      4337769816      4                     1591056576414  1591056576414   0                 4333023554
1591034132029         1591034209011               JASON077                    4337769817     3                    1591034132029          1591034209011                JASON077                     4337769817      3                     1591056576414  1591056576414   0                 4333023554
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
