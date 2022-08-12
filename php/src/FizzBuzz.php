<?php

namespace App;

class FizzBuzz
{
    static function series(int $from, int $to): array
    {
        $results = [];

        for ($i = $from; $i <= $to; $i++) {
            $results[] = [
                'step' => $i,
                'value' => static::step($i),
            ];
        }

        return $results;
    }

    static function step(int $num): string
    {
        if ($num % 3 == 0 && $num % 5 == 0) {
            return "FizzBuzz";
        }

        if ($num % 3 == 0) {
            return "Fizz";
        }

        if ($num % 5 == 0) {
            return "Buzz";
        }

        return $num;
    }

    static function printSeries(array $values): void
    {
        array_map(fn ($item) => printf(json_encode($item) . PHP_EOL), $values);
    }
}
