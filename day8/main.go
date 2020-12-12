package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func ParseLine(num int, value string) Instruction {
	name := value[0:3]
	val, _ := strconv.Atoi(value[4:len(value)])

	return Instruction{num, name, val}
}

func CopyInstructions(m map[int]Instruction) map[int]Instruction {
    cp := make(map[int]Instruction)
    for k, v := range m {
        cp[k] = v
    }

    return cp
}

func main() {
    f, err := os.Open("./input")
    if err != nil {
        log.Fatal(err)
    }

    defer func() {
        if err = f.Close(); err != nil {
            log.Fatal(err)
        }
    }()

    stack := new(Stack)
    stack.Init()

    s := bufio.NewScanner(f)
    number := 0
    for s.Scan() {
        line := s.Text()
        stack.instructions[number] = ParseLine(number, line)
        number += 1
    }

    fmt.Println("There are", len(stack.instructions), "instructions")

    stack.Run()

    // Part 1
    if (stack.loop) {
    	fmt.Println("The acc before loop is", stack.acc)
    } else {
    	fmt.Println("The acc after all execution is", stack.acc)
    }

    // Instructions copy
    backup := CopyInstructions(stack.instructions);

    // Point of failure
    failure := -1

    counter := 0

    // Part 2, try fix
    for {
        last := stack.Run()
        counter += 1

        if (!stack.loop) {
            break
        }

        stack.instructions = CopyInstructions(backup)

        if (failure == 0) {
            fmt.Println("Can't fix the instructions")
            break
        }

        if (failure == -1) {
            failure = last.line
        }

        // changeable instruction
        instruction := stack.instructions[failure]

        // changes
        if (instruction.op == "jmp") {
            instruction.op = "nop"
        } else if instruction.op == "nop" {
            instruction.op = "jmp"
        }

        stack.instructions[failure] = instruction

        stack.Reset()
        failure -= 1
    }

    fmt.Println("The acc on finished execution is", stack.acc, "Tried", counter)

    err = s.Err()
    if err != nil {
        log.Fatal(err)
    }
}