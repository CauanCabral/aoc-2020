package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

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

    group := new(Group)
    group.Initialize()

    sums := 0
    allSums := 0
    s := bufio.NewScanner(f)
    for s.Scan() {
        var line = s.Text()

        // End a group
        if (len(line) == 0) {
            // Compute their values
            sums += group.Sum()
            allSums += group.AllSum()

            // Init new one
            group = new(Group)
            group.Initialize()

            continue
        }

        group.AddLine(line)
    }

    sums += group.Sum()
    allSums += group.AllSum()

    // Part 1
    fmt.Println("The sum of unique answers are", sums)

    // Part 2
    fmt.Println("The sum of unique and everyone answers are", allSums)

    err = s.Err()
    if err != nil {
        log.Fatal(err)
    }
}
