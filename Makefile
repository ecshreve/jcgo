build:
	go build -o bin/jcgo github.com/ecshreve/jcgo

run-only:
	bin/jcgo $(INFILE)

run: build run-only

test:
	go test github.com/ecshreve/jcgo/...

testv:
	go test -v github.com/ecshreve/jcgo/...
