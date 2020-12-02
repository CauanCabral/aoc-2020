<?php
$input = require('input.php');

sort($input);
$expenses = $x = $y = null;
for ($i = 0; $i < 199; $i++) {
    for ($j = 199; $j > 1; $j--) {
        $x = $input[$i];
        $y = $input[$j];

        if ($x + $y === 2020) {
            $expenses = $x * $y;
            break 2;
        }
    }
}

echo "Values: ", $x, " and ", $y, "\n";
echo "Expenses: ", $expenses, "\n";