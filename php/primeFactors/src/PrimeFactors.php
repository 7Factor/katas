<?php

namespace App;

// 1. is it divisible by 2
// 2. a. if true, divide by 2
//     b. if false, increase the divisor
// 3. repeat

class PrimeFactors
{
    public function generate(int $number)
    {
        $factors = [];
        $divisor = 2;

        while ($number > 1)
        {
            while ($number % $divisor === 0)
            {
                $factors[] = $divisor;
                $number /= $divisor;
            }

            $divisor++;
        }
        

        return $factors;
    }
}