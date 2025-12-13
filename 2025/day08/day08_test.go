package day08

import "testing"

var TEST_PATH = "./test-input.txt"

func TestPart1(t *testing.T) {
	got := part1(TEST_PATH, 10)
	expected := 40

	if got != expected {
		t.Errorf("Got %d, expected %d", got, expected)
	}
}

func TestPart2(t *testing.T) {
	got := part2(TEST_PATH)
	expected := 25272

	if got != expected {
		t.Errorf("Got %d, expected %d", got, expected)
	}
}
