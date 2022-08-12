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
    print([$nto, FizzBuzz::step($nTo)]);
}
