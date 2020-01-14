const sorter = require('./solution');

describe('The Integer Sorting Kata', () => {
  test('should return empty array with null input', () => {
    _runSort(null, []);
  });

  test('should return empty array with empty array input', () => {
    _runSort([], [])
  });

  test('should return the array with single item array input', () => {
    _runSort([1], [1]);
  });

  test('should return the array with a sorted, two item array input', () => {
    _runSort([1,2], [1,2]);
  });

  test('should return a sorted array with an unsorted, two item array input', () => {
    _runSort([2,1], [1,2]);
  });

  test('should return the array with a sorted, three item array input', () => {
    _runSort([1,2,3], [1,2,3]);
  });

  test('should return the array with the first element of a three item array out of order', () => {
    _runSort([2,1,3], [1,2,3]);
  });

  test('should return the array with the second element of a three item array out of order', () => {
    _runSort([1,3,2], [1,2,3])
  })
});

let _runSort = (input, expected) => {
  let actual = sorter.sort(input);
  expect(actual).toEqual(expected);
};
