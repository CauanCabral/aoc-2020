package main

import(
    // "fmt"
)

type Instruction struct {
    line int
    op string
    value int
}

type Stack struct {
    acc int
    path map[int]int
    instructions map[int]Instruction
    loop bool
}

func (s *Stack) Init() {
    if s.path == nil {
        s.path = make(map[int]int)
    }

    if s.instructions == nil {
        s.instructions = make(map[int]Instruction)
    }

    s.acc = 0
    s.loop = false
}

func (s *Stack) Reset() {
    s.path = make(map[int]int)

    s.acc = 0
    s.loop = false
}

// Execute the instruction and return next line
func (s *Stack) Execute(i Instruction) int {
    _, repeated := s.path[i.line]
    if (repeated) {
        s.loop = true

        return i.line
    }

    s.path[i.line] = 1

    if i.op == "jmp" {
        return i.line + i.value
    }

    if i.op == "acc" {
        s.acc += i.value
    }

    return i.line + 1
}

// Execute all stack instructions, stopping when find a loop
// or if finish
func (s *Stack) Run() Instruction {
    line := 0
    size := len(s.instructions)
    last := s.instructions[line]

    for {
        inst := s.instructions[line]
        nextLine := s.Execute(inst)

        // Found a loop
        if (s.loop) {
            return last
        }

        // Finish execution
        if (nextLine >= size) {
            return last
        }

        last = s.instructions[line]
        line = nextLine
    }

    return last
}