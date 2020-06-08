package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/kr/pretty"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		log.Fatalf("please provide a file")
	}

	path := args[1]
	ext := filepath.Ext(path)
	if ext != ".json" {
		log.Fatalf("please provide a .json file")
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("successfully opened file %s\n", path)
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	pretty.Print(result)
}
