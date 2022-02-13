package main

import (
	"fmt"
	"os"
	"strings"
)

type DNA string
type RNA string

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadDataFromFile(path string) DNA {
	data, err := os.ReadFile(path)

	Check(err)

	return DNA(data)
}

func (dna DNA) Transcribe() RNA {
	return RNA(strings.Replace(string(dna), "T", "U", -1))
}

func main() {
	data := ReadDataFromFile("../data/dataset.txt")
	rna := data.Transcribe()

	fmt.Println("Result:", rna)
}
