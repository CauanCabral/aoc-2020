<?php
$input = require('input.php');

sort($input);
$expenses = $x = $y = $z = null;
for ($i = 0; $i < 198; $i++) {
    for ($k = $i; $k < 199; $k++) {
        for ($j = 199; $j > 1; $j--) {
            $x = $input[$i];
            $y = $input[$j];
            $z = $input[$k];

            if ($x + $y + $z === 2020) {
                $expenses = $x * $y * $z;
                break 3;
            }
        }
    }
}

echo "Values: ", $x, ", ", $y, " and ", $z, "\n";
echo "Expenses: ", $expenses, "\n";