# jcgo

JSON to CSV converter in Golang.

## Usage

`jcgo` reads JSON from the file specified as a command line argument, parses the JSON, and writes the data to a CSV file.

```{bash}
jcgo jsontestlocal.json
```

```{bash}
cat jsontestlocal.output.csv | column -t -s,
>
afterState_arrivedAt  afterState_departureTimeMs  afterState_destinationName  afterState_id  afterState_jobState  beforeState_arrivedAt  beforeState_departureTimeMs  beforeState_destinationName  beforeState_id  beforeState_jobState  changedAtMs    events_eventAt  events_eventType  id
0                     0                           CARI307                     4337769816     4                    0                      0                            CARI307                      4337769816      4                     1591056576414  1591056576414   0                 4333023554
1591034132029         1591034209011               JASON077                    4337769817     3                    1591034132029          1591034209011                JASON077                     4337769817      3                     1591056576414  1591056576414   0                 4333023554
```

## Reference Links

- [Effective Go](https://golang.org/doc/effective_go.html)
- [Using Go Modules](https://blog.golang.org/using-go-modules)
- [encoding/json](https://golang.org/pkg/encoding/json/)
