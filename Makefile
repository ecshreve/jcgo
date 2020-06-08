build:
	go build -o bin/jcgo github.com/ecshreve/jcgo

run-only:
	bin/jcgo $(INFILE)

run: build run-only
