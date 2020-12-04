<?php
$pattern = array_map('trim', file('./input'));
require('map.php');

// Walk every line, $yStep by time
// Walk $xStep column by time
function sloopy(Map $map, $xStep, $yStep): int {
    $trees = 0;

    // adjust virtual map
    $map->setWidth($map->getHeight() * $xStep);

    for ($i = 0, $j = 0; $i < $map->getHeight(); $i += $yStep, $j += $xStep) {
        $trees += ($map->get($j, $i) === '#' ? 1 : 0);
    }

    return $trees;
}

$map = new Map($pattern);

$tree11 = sloopy($map, 1, 1);
$tree31 = sloopy($map, 3, 1);
$tree51 = sloopy($map, 5, 1);
$tree71 = sloopy($map, 7, 1);
$tree12 = sloopy($map, 1, 2);

echo "Trees (1,1): ", $tree11, "\n";
echo "Trees (3,1): ", $tree31, "\n";
echo "Trees (5,1): ", $tree51, "\n";
echo "Trees (7,1): ", $tree71, "\n";
echo "Trees (1,2): ", $tree12, "\n";

echo "Multiplied = ", $tree11 * $tree31 * $tree51 * $tree71 * $tree12, "\n";