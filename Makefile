build:
	go build -o bin/jcgo github.com/ecshreve/jcgo/cmd/jcgo

install:
	go install -i github.com/ecshreve/jcgo/cmd/jcgo

run-only:
	bin/jcgo $(INFILE)

run: build run-only

test:
	go test github.com/ecshreve/jcgo/...

testv:
	go test -v github.com/ecshreve/jcgo/...

testc:
	go test -race -coverprofile=coverage.txt -covermode=atomic github.com/ecshreve/jcgo/...

clean:
	rm *.csv

