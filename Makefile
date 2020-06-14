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
	go test -cover github.com/ecshreve/jcgo/...

