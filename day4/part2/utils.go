package main

import (
    // "fmt"
    "strings"
    "strconv"
    "regexp"
)

// year must be int and between min and max
func validYear(value string, min int, max int) bool {
    year, err := strconv.Atoi(value)
    if (err != nil) {
        return false
    }

    return (year >= min && year <= max)
}

// hgt (Height) - a number followed by either cm or in:
//     If cm, the number must be at least 150 and at most 193.
//     If in, the number must be at least 59 and at most 76.
func validMeasurement(value string) bool {
    validLimits := map[string]map[string]int{
        "cm": {"min": 150, "max": 193},
        "in": {"min": 59, "max": 76},
    }

    valueLen := len(value)
    unit := value[valueLen - 2:]

    limits, exists := validLimits[unit]
    if !exists {
        return false
    }

    height, err := strconv.Atoi(value[: valueLen - 2])
    if (err != nil) {
        return false
    }

    return (height >= limits["min"] && height <= limits["max"])
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func validHexColor(value string) bool {
    hash := value[0:1]
    if hash != "#" {
        return false
    }

    color := value[1:]
    found, err := regexp.Match("[a-f0-9]{6}", []byte(color))
    if (err != nil) {
        return false
    }

    return found
}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func validEyeColor(value string) bool {
    validEyeColor := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}
    _, exists := validEyeColor[value]

    return exists
}

// pid (Passport ID) - a nine-digit number, including leading zeroes.
func validPassport(value string) bool {
    found, err := regexp.Match("^[0-9]{9}$", []byte(value))
    if (err != nil) {
        return false
    }

    return found
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