<?php

include 'vendor/autoload.php';

use App\ArgvParser;
use App\FizzBuzz;

[
    'nFrom' => $nFrom,
    'nTo' => $nTo,
    'print' => $print
] = ArgvParser::parse();

$values = FizzBuzz::series($nFrom, $nTo);

if ($print == 'series') {
    FizzBuzz::printSeries($values);
} else {
    print($nTo . ' ' .  FizzBuzz::step($nTo) . PHP_EOL);
}

// print("{$nFrom} {$nTo} {$print}" . PHP_EOL);
