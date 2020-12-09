<?php
include 'bag.php';

function parseRule(string $rule, array &$bags): void {
    $parts = explode(' contain ', $rule);
    $actual = substr($parts[0], 0, -5);

    $bag = findOrCreate($actual, $bags);

    $contains = explode(', ', $parts[1]);
    foreach ($contains as $contain) {
        $contain = str_replace(['.', 'bags', 'bag'], '', $contain);
        $limit = intval(substr($contain, 0, 1));
        $internal = trim(substr($contain, 1));

        if ($limit == 0) {
            continue;
        }

        $obj = findOrCreate($internal, $bags);

        $bag->addChild(child: $obj, limit: $limit);
    }
}

function findOrCreate(string $name, array &$bags): Bag {
    if (isset($bags[$name])) {
        return $bags[$name];
    }

    return $bags[$name] = new Bag($name);
}

$bags = [];
$input = file('./input');

foreach ($input as $line) {
    parseRule($line, $bags);
}

$size = count($bags);
echo "There are {$size} bags\n";

// Part 1
$eventuallyParents = $bags['shiny gold']->eventuallyParents();
$parentsLen = count($eventuallyParents);
echo "Wich are {$parentsLen} eventually parents of shiny gold\n";

// Part 2
$mustContain = $bags['shiny gold']->mustContain();
echo "Wich must contain others {$mustContain} bags\n";