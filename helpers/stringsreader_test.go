package helpers

import (
	"testing"
)

func TestParseStrings(t *testing.T) {
	testString := []uint8{192, 230, 237, 255, 255, 255, 255}
	stringParsed := ReadString(testString)
	if stringParsed != "Fry" {
		t.Fatalf("the string should be 'Fry' but is: %v", stringParsed)
	}
}
