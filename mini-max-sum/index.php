<?php

miniMaxSum([1, 2, 3, 4, 5]);
miniMaxSum([1, 3, 5, 7, 9]);

function miniMaxSum(array $arr): void
{
    $min = 0;
    $max = 0;

    foreach ($arr as $key => $value) {
        if ($value < $min || $min == 0) {
            $min = $value;
        }

        if ($value > $max) {
            $max = $value;
        }
    }

    $minSum = array_sum(array_diff($arr, [$min]));
    $maxSum = array_sum(array_diff($arr, [$max]));

    if ($minSum > $maxSum) {
        echo "$maxSum  $minSum\n";
    } else {
        echo "$minSum  $maxSum\n";
    }
}