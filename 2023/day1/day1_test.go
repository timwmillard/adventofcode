package main

import (
	"strings"
	"testing"
)

const file = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

func TestCalibration(t *testing.T) {

	input := strings.NewReader(file)

	want := 142
	got, err := Calibration(input)
	if err != nil {
		t.Fatal("Calibration error", err)
	}

	if want != got {
		t.Errorf("Calibration() got %d but want %d", got, want)
	}
}
