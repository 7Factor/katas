using System.Collections.Generic;

namespace PrimeFactors.Common
{
    public static class PrimeFactorGenerator
    {
        public static List<int> GenerateFactors(int number)
        {
            var retFactors = new List<int>();

            for (var divisor = 2; number > 1; divisor++)
            {
                while (number % divisor == 0)
                {
                    retFactors.Add(divisor);
                    number /= divisor;
                }
            }

            return retFactors;
        }
    }
}
