package main

import (
    "testing"
)

// Test row decode
func TestSeatRow(t *testing.T) {
    expected := 70
    actual := Seat{"BFFFBBFRRR"}.Row()
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    expected = 14
    actual = Seat{"FFFBBBFRRR"}.Row()
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    expected = 102
    actual = Seat{"BBFFBBFRLL"}.Row()
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }
}

// Test column decode
func TestSeatColumn(t *testing.T) {
    expected := 7
    actual := Seat{"BFFFBBFRRR"}.Column()
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    expected = 7
    actual = Seat{"FFFBBBFRRR"}.Column()
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    expected = 4
    actual = Seat{"BBFFBBFRLL"}.Column()
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }
}

// Test seat ID calc
func TestSeatId(t *testing.T) {
	expected := 567
    actual := Seat{"BFFFBBFRRR"}.Id()
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    expected = 119
    actual = Seat{"FFFBBBFRRR"}.Id()
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    expected = 820
    actual = Seat{"BBFFBBFRLL"}.Id()
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }
}