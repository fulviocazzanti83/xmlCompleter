package main

import "testing"

func TestFillFile(t *testing.T) {

	org := "./sampleData/exportFederalMugol.xml"
	dest := "./sampleData/dest.xml"

	fillFile(org, dest)

}
