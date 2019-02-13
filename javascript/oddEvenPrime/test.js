import isOddEvenOrPrime from './solution';

describe('Odd, Even, Prime', () => {

  test('Should determine that 4 is even', () => {
    expect(isOddEvenOrPrime(4)).toEqual('Even');
  });

  test('Should determine that 9 is odd', () => {
    expect(isOddEvenOrPrime(9)).toEqual('Odd');
  });

  test('Should determine that 6 is even', () => {
    expect(isOddEvenOrPrime(6)).toEqual('Even');
  });

  test('Should determine that 15 is odd', () => {
    expect(isOddEvenOrPrime(15)).toEqual('Odd');
  });

  test('Should determine that 3 is prime', () => {
    expect(isOddEvenOrPrime(3)).toEqual('Prime');
  });

  test('Should determine that 67 is prime', () => {
    expect(isOddEvenOrPrime(67)).toEqual('Prime');
  });

  test('Should determine that 1 is odd', () => {
    expect(isOddEvenOrPrime(1)).toEqual('Odd');
  });

  test('Should determine that 2 is prime', () => {
    expect(isOddEvenOrPrime(2)).toEqual('Prime');
  });
});
