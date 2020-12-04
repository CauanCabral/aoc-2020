<?php
$pattern = array_map('trim', file('./input'));
$lines = count($pattern);
$columns = strlen($pattern[0]);

// Don't alloc memory to all position.
// Dynamically determine its value
function virtualMap(array $pattern, int $columns, int $x, int $y): string {
    $k = $x % $columns;
    return $pattern[$y][$k];
}

// Walk every line, $yStep by time
// Walk $xStep column by time
function sloopy($pattern, $lines, $columns, $xStep, $yStep): int {
    $trees = 0;

    for ($i = 0, $j = 0; $i < $lines; $i += $yStep, $j += $xStep) {
        $value = virtualMap($pattern, $columns, $j, $i);
        $trees += ($value === '#' ? 1 : 0);
    }

    return $trees;
}

echo "Trees: ", sloopy($pattern, $lines, $columns, 3, 1), "\n";