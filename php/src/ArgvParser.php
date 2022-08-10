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
        $results = static::$defaultParameters;

        $arguments = array_map(function ($item) {
            [$param, $value] = explode('=', $item);

            return [
                'param' => $param,
                'value' => $value
            ];
        }, array_slice($argv, 2));

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
