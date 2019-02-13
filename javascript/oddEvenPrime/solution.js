function isPrime(numberInput) {
  if (numberInput === 1) return false;

  for (let i = 2; i < numberInput; i++) {
    if (numberInput % i === 0) {
      return false;
    }
  }
  return true;
}

function isEven(numberInput) {
  return numberInput % 2 === 0;
}

function isOdd(numberInput) {
  return numberInput % 2 !== 0;
}

export default function isOddEvenOrPrime(numberInput) {

  if (isPrime(numberInput)) {
    return 'Prime';
  }

  if (isEven(numberInput)) {
    return 'Even';
  }

  if (isOdd(numberInput)) {
    return 'Odd';
  }
}
