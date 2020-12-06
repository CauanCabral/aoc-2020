package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

// Validation rules
//
// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in:
//     If cm, the number must be at least 150 and at most 193.
//     If in, the number must be at least 59 and at most 76.
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.
// cid (Country ID) - ignored, missing or not.
func isValid(doc map[string]string) bool {
    for _, field := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
        value, exists := doc[field]
        if exists == false {
            return false
        }

        if len(value) == 0 {
            return false
        }

        if (field == "byr" && !validYear(value, 1920, 2002)) {
            return false
        }

        if (field == "iyr" && !validYear(value, 2010, 2020)) {
            return false
        }

        if (field == "eyr" && !validYear(value, 2020, 2030)) {
            return false
        }

        if (field == "hgt" && !validMeasurement(value)) {
            return false
        }

        if (field == "hcl" && !validHexColor(value)) {
            return false
        }

        if (field == "ecl" && !validEyeColor(value)) {
            return false
        }

        if (field == "pid" && !validPassport(value)) {
            return false
        }
    }

    return true
}

func main() {
    f, err := os.Open("../input")
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
