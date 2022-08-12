<?php

namespace App;

use JetBrains\PhpStorm\ArrayShape;

class ArgvParser
{
    static protected array $defaultParameters = [
        'nFrom' => 1,
        'nTo' => 100,
        'print' => 'series'
    ];

    #[ArrayShape([
        'nFrom' => 'int',
        'nTo' => 'int',
        'print' => 'string'
    ])]
    static function parse(): array
    {
        global $argv;

        // print(json_encode($argv));

        $results = static::$defaultParameters;

        $arguments = array_map(function ($item) {
            $parts = explode('=', $item);

            return [
                'param' => $parts[0],
                'value' => $parts[1] ?? null
            ];
        }, array_slice($argv, 1));

        foreach ($arguments as $argument) {
            if (isset($results[$argument['param']])) {
                $results[$argument['param']] = $argument['value'];
            }
        }

        // sanitize

        return [
            'nFrom' =>  (int)$results['nFrom'],
            'nTo' => (int)$results['nTo'],
            'print' => in_array($results['print'], ['series', 'end']) ? $results['print'] : 'series'
        ];
    }
}
