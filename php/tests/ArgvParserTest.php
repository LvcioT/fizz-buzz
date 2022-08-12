<?php

namespace Tests;

use App\ArgvParser;
use PHPUnit\Framework\TestCase;

class ArgvParserTest extends TestCase
{
    protected array $defaultValues = [
        'nFrom' => 1,
        'nTo' => 100,
        'print' => 'series'
    ];

    public function testDefault()
    {
        global $argv;
    
        // no values (fist one is always script name) 
        $argv = ['dummy1'];
        $this->assertEquals($this->defaultValues, ArgvParser::parse());

        // extra arguments
        $argv = ['dummy1', 'dummy=argument'];
        $this->assertEquals($this->defaultValues, ArgvParser::parse());
    }

    public function testArgs()
    {
        global $argv;

        $argv = ['dummy1', 'nFrom=-5'];
        $this->assertEquals(['nFrom' => -5] + $this->defaultValues, ArgvParser::parse());

        $argv = ['dummy1', 'print=end'];
        $this->assertEquals(['print' => 'end'] + $this->defaultValues, ArgvParser::parse());
    }

    public function testArgsValue()
    {
        global $argv;
     
        // actual arguments with useless value
        $argv = ['dummy1', 'print=everything'];
        $this->assertEquals($this->defaultValues, ArgvParser::parse());
    }
}
