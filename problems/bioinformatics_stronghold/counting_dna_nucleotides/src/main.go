package main

import (
	"fmt"
	"os"
)

type DNA string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readDataFromFile(path string) DNA {
	data, err := os.ReadFile(path)

	check(err)

	return DNA(data)
}

func (data DNA) countNucleotides(symbol rune) int32 {
	count := 0

	for _, nucleotide := range data {
		if nucleotide == symbol {
			count++
		}
	}

	return int32(count)
}

func main() {
	data := readDataFromFile("../data/dataset.txt")

	numberOfA := data.countNucleotides('A')
	numberOfC := data.countNucleotides('C')
	numberOfG := data.countNucleotides('G')
	numberOfT := data.countNucleotides('T')

	fmt.Println("Order of symboles: A C G T")
	fmt.Printf("%d %d %d %d", numberOfA, numberOfC, numberOfG, numberOfT)
}
