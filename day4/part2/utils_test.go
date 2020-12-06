package main

import (
    "testing"
)

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
func TestValidBirthYear(t *testing.T) {
    expected := true
    actual := validYear("1920", 1920, 2002)
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }

    expected = true
    actual = validYear("1980", 1920, 2002)
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }

    expected = true
    actual = validYear("2002", 1920, 2002)
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }

    expected = false
    actual = validYear("2003", 1920, 2002)
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }

    expected = false
    actual = validYear("2010", 1920, 2002)
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }

    expected = false
    actual = validYear("1919", 1920, 2002)
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }
}

// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func TestValidIssueYear(t *testing.T) {
    expected := true
    actual := validYear("2010", 2010, 2020)
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }

    expected = true
    actual = validYear("2011", 2010, 2020)
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }

    expected = true
    actual = validYear("2020", 2010, 2020)
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }

    expected = false
    actual = validYear("2009", 2010, 2020)
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }

    expected = false
    actual = validYear("1900", 2010, 2020)
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }

    expected = false
    actual = validYear("2021", 2010, 2020)
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }
}

// hgt (Height) - a number followed by either cm or in:
//     If cm, the number must be at least 150 and at most 193.
//     If in, the number must be at least 59 and at most 76.
func TestValidMeasurement(t *testing.T) {
    expected := true
    actual := validMeasurement("190cm")
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }

    expected = true
    actual = validMeasurement("60in")
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }

    expected = false
    actual = validMeasurement("70")
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }

    expected = false
    actual = validMeasurement("194cm")
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }

    expected = false
    actual = validMeasurement("58in")
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func TestValidHexColor(t *testing.T) {
    expected := true
    actual := validHexColor("#123abc")
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }

    expected = false
    actual = validHexColor("#123abz")
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }

    expected = false
    actual = validHexColor("123abc")
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }
}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func TestValidEyeColor(t *testing.T) {
    expected := true
    actual := validEyeColor("brn")
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }

    expected = true
    actual = validEyeColor("oth")
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }

    expected = false
    actual = validEyeColor("wat")
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }
}

func TestValidPassport(t *testing.T) {
    expected := true
    actual := validPassport("000000001")
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }

    expected = true
    actual = validPassport("100000000")
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }

    expected = false
    actual = validPassport("0123456789")
    if actual != expected {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", expected, actual)
    }
}