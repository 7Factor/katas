using System;
using PrimeFactors.Common;

namespace PrimeFactors
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine(string.Join(",", PrimeFactorGenerator.GenerateFactors(5123523)));
        }
    }
}
