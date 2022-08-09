<?php

namespace Tests;

use PHPUnit\Framework\TestCase;

class BaseTest extends TestCase
{
    public function testBase()
    {
        $this->assertEquals(true, 1 == 1);
    }
}
