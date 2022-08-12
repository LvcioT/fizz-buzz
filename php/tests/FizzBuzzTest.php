<?php

namespace Tests;

use App\FizzBuzz;
use PHPUnit\Framework\TestCase;

class FizzBuzzTest extends TestCase
{
    public function testStep() {
        $this->assertEquals(11, FizzBuzz::step(11));
        $this->assertEquals('Fizz', FizzBuzz::step(12));
        $this->assertEquals('Buzz', FizzBuzz::step(25));
        $this->assertEquals('FizzBuzz', FizzBuzz::step(30));
    }

    public function testSeries() {
        $this->assertEquals(21, count( FizzBuzz::series(-10,10)));
        $this->assertEquals(-10,  FizzBuzz::series(-10,0)[0]['step']);
        $this->assertEquals(10,  FizzBuzz::series(0,10)[10]['step']);
        
    }
}
