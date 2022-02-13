package main

import "testing"

func TestTranscribing(t *testing.T) {
	dna := ReadDataFromFile("../data/sample_dataset.txt")
	expectedResult := ReadDataFromFile("../data/sample_result.txt")

	rna := dna.Transcribe()

	if string(expectedResult) != string(rna) {
		t.Errorf("Expected %s, got %s", expectedResult, rna)
	}
}
