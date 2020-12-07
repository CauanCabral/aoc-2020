package main

type Group struct {
    yes map[string]int
    members int
}

// Default init
func (g *Group) Initialize() {
    if g.yes == nil {
        g.yes = make(map[string]int)
    }

    g.members = 0
}

// Parse a person answer
func (g *Group) AddLine(line string) {
    size := len(line)
    g.members += 1

    for i := 0; i < size; i++ {
        val := line[i:(i + 1)]
        _, exists := g.yes[val]

        if (exists) {
            g.yes[val] += 1

            continue
        }

        g.yes[val] = 1
    }
}

// Compute all yes answers of group
func (g Group) Sum() int {
    return len(g.yes)
}

// Compute everyone present symbol
func (g Group) AllSum() int {
    sum := 0

    for _, counts := range g.yes {
        if (counts == g.members) {
            sum += 1
        }
    }

    return sum
}