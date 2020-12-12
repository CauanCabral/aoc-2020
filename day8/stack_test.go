package main

import (
    "testing"
)

// Test Init
func TestInit(t *testing.T) {
    s := new(Stack)
    s.Init()

    expected := 0
    actual := s.acc
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    if s.loop {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", false, s.loop)
    }

    expected = 0
    actual = len(s.instructions)
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }
}

func TestExecute(t *testing.T) {
    s := new(Stack)
    s.Init()

    inst1 := Instruction{0, "acc", 2}
    s.instructions[0] = inst1

    nextLine := s.Execute(inst1)

    expected := 1
    actual := nextLine
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    if s.loop {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", false, s.loop)
    }

    expected = 2
    actual = s.acc
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    expected = 1
    actual = len(s.instructions)
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    inst2 := Instruction{1, "acc", -1}
    s.instructions[1] = inst2
    nextLine = s.Execute(inst2)

    expected = 2
    actual = nextLine
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    if s.loop {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", false, s.loop)
    }

    expected = 1
    actual = s.acc
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    expected = 2
    actual = len(s.instructions)
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    inst3 := Instruction{2, "jmp", -2}
    s.instructions[2] = inst3
    nextLine = s.Execute(inst3)

    expected = 0
    actual = nextLine
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    if s.loop {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", false, s.loop)
    }

    expected = 1
    actual = s.acc
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    expected = 3
    actual = len(s.instructions)
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    // No executed, loop detected
    nextLine = s.Execute(inst1)

    expected = 0
    actual = nextLine
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    if !s.loop {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", true, s.loop)
    }

    expected = 1
    actual = s.acc
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }

    expected = 3
    actual = len(s.instructions)
    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
    }
}

func TestRun(t *testing.T) {
    s := new(Stack)
    s.Init()

    s.instructions[0] = Instruction{0, "acc", 2}
    s.instructions[1] = Instruction{1, "acc", -1}
    s.instructions[2] = Instruction{2, "jmp", -2}
    s.instructions[3] = Instruction{3, "acc", 3}

    actual := s.Run()
    expected := s.instructions[2]

    if actual != expected {
        t.Errorf("Test failed, expected: '%d', got:  '%d'", expected.line, actual.line)
    }

    if !s.loop {
        t.Errorf("Test failed, expected: '%t', got:  '%t'", true, s.loop)
    }
}