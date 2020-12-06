package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "sort"
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

    var ids []int
    highestId := 0
    s := bufio.NewScanner(f)
    for s.Scan() {
        var line = s.Text()
        var s = Seat{line}
        var id = s.Id()

        ids = append(ids, id)

        if (id > highestId) {
            highestId = id
        }
    }

    // Part 1
    fmt.Println("The highest Seat ID is ", highestId)

    sort.Ints(ids)

    // Part 2
    var size int = len(ids)
    var smallest int = ids[0]
    for i := smallest; i < highestId; i++ {
        var pos int = sort.SearchInts(ids, i)

        if (pos >= size || ids[pos] != i) {
            fmt.Println("Your Seat ID is ", i)
        }
    }

    err = s.Err()
    if err != nil {
        log.Fatal(err)
    }
}
