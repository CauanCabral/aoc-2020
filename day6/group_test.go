package main

import (
    "testing"
)

// Test AddLine
func TestAddLine(t *testing.T) {
    g := new(Group)
    g.Initialize()
    g.AddLine("abc")

    expected := 1
    actual, exists := g.yes["b"]
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }
    if exists != true {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", true, exists)
    }

    expected = 0
    actual, exists = g.yes["x"]
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }
    if exists != false {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", false, exists)
    }

    g.AddLine("axzb")

    expected = 1
    actual, exists = g.yes["x"]
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }
    if exists != true {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", true, exists)
    }

    expected = 2
    actual, exists = g.yes["b"]
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }
    if exists != true {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", true, exists)
    }
}

func TestSum(t *testing.T) {
    g := new(Group)
    g.Initialize()
    g.AddLine("abc")

    expected := 3
    actual := g.Sum()
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    g.AddLine("ab")

    expected = 3
    actual = g.Sum()
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    g.AddLine("xyz")

    expected = 6
    actual = g.Sum()
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }
}

func TestAllSum(t *testing.T) {
    g := new(Group)
    g.Initialize()
    g.AddLine("abc")

    expected := 3
    actual := g.AllSum()
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    g.AddLine("ab")

    expected = 2
    actual = g.AllSum()
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    g.AddLine("xyz")

    expected = 0
    actual = g.AllSum()
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }
}