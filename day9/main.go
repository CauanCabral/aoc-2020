package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "sort"
    "strconv"
)

func isValid(numberList []int, number int, start int, end int) bool {
    for i := start; i < end; i += 1 {
        x := numberList[i]
        for j := i + 1; j < end; j += 1 {
            y := numberList[j]

            if (x + y == number) {
                return true
            }
        }
    }

    return false
}

func searchWeakness(numberList []int, target int) []int {
    var sequence []int
    limit := len(numberList)
    for i := 0; i < limit - 2; i++ {
        sum := 0
        sequence = nil

        for j := i + 1; j < limit; j++ {
            sum += numberList[j]
            sequence = append(sequence, numberList[j])

            if (sum > target) {
                break
            }

            if (sum == target && j - i > 1) {
                return sequence
            }
        }
    }

    return sequence
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

    var numberList []int
    counter := 0
    size := 25

    s := bufio.NewScanner(f)
    for s.Scan() {
        number, _ := strconv.Atoi(s.Text())
        numberList = append(numberList, number)
        counter += 1
    }

    // Part 1
    var invalid int = -1
    for i := size; i < counter; i++ {
        number := numberList[i]
        if !isValid(numberList, number, i - size, i) {
            invalid = number
            break;
        }
    }

    if (invalid != -1) {
        fmt.Println("Invalid number is", invalid)
    } else {
        log.Fatal("Invalid number not found")
    }

    // Part 2
    weak := searchWeakness(numberList, invalid)

    if (len(weak) == 0) {
        log.Fatal("Cannot find weakness sequence")
    }

    sort.Ints(weak)

    smallest := weak[0]
    largest := weak[len(weak) - 1]
    fmt.Println("Sum of smallest and largest is", smallest + largest)

    err = s.Err()
    if err != nil {
        log.Fatal(err)
    }
}
