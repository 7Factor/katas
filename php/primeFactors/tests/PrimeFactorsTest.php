<?php

use App\PrimeFactors;
use PHPUnit\Framework\TestCase;

class PrimeFactorsTest extends TestCase
{
    /** 
     * @dataProvider factors
     */
    function test_generate_factors($number, $expected)
    {
        $factors = new PrimeFactors;
        $this->assertEquals($expected, $factors->generate($number));
    }

    public function factors()
    {
        return [
            [1, []],
            [2, [2]],
            [3, [3]],
            [4, [2, 2]],
            [30, [2, 3, 5]], 
            [100, [2, 2, 5, 5]]
        ];
    }
    
}