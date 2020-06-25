using System;
using System.Collections.Generic;
using System.Linq;

namespace WordWrap.Common
{
    public static class WordWrapper
    {
        public static string WrapWords(string input, int rowLength)
        {
            if (input.Length <= rowLength || rowLength < 1)
            {
                return input;
            }

            var space = input.Substring(0, rowLength).LastIndexOf(' ');

            if (space != -1)
            {
                return $"{input.Substring(0, space)}\n{WrapWords(input.Substring(space + 1), rowLength)}";
            }

            if (input.ElementAt(rowLength) == ' ')
            {
                return $"{input.Substring(0, rowLength)}\n{WrapWords(input.Substring(rowLength + 1), rowLength)}";
            }

            return $"{input.Substring(0, rowLength)}\n{WrapWords(input.Substring(rowLength), rowLength)}";
        }
    }
}
