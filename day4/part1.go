package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

// Document have these fields
// byr (Birth Year)
// iyr (Issue Year)
// eyr (Expiration Year)
// hgt (Height)
// hcl (Hair Color)
// ecl (Eye Color)
// pid (Passport ID)
// cid (Country ID)

/**
 * A valid document should had a value for every fields, except for 'cid', which is
 * optional
 */
func isValid(doc map[string]string) bool {
    for _, field := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
        value, exists := doc[field]
        if exists == false {
            return false
        }

        if len(value) == 0 {
            return false
        }
    }

    return true
}

/**
 * A line contain one or more field:value definition separeted by space
 */
func parseLine(line string, doc map[string]string) map[string]string {
    var splits = strings.Split(line, " ")
    for i := 0; i < len(splits); i++ {
        var kv = strings.Split(splits[i], ":")
        doc[kv[0]] = kv[1]
    }

    return doc;
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

    var doc = make(map[string]string)
    valids := 0
    docs := 0
    s := bufio.NewScanner(f)
    for s.Scan() {
        var line = s.Text()

        // Begin a document
        if (line == "") {
            docs++

            // test previous one
            if (isValid(doc)) {
                valids++
            }

            // reset document
            doc = make(map[string]string)

            continue
        }

        doc = parseLine(line, doc)
    }

    // last doc
    docs++
    if (isValid(doc)) {
        valids++
    }

    fmt.Println("There are ", docs, " documents. Valids one are ", valids)

    err = s.Err()
    if err != nil {
        log.Fatal(err)
    }
}