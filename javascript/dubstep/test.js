const songDecoder = require('./solution');

describe('Dubstep', () => {
  test('should not accept a null input', () => {
    let actual = songDecoder.decode(null);
    expect(actual).toEqual(new Error('null is not allowed'))
  });

  test('should not accept lowercase letters in input', () => {
    let actual = songDecoder.decode('somelowercaseletters');
    expect(actual).toEqual(new Error('lower case letters are not allowed'))
  });

  test('should not accept whitespace characters in input', () => {
    let actual = songDecoder.decode('  WHITESPACE   ');
    expect(actual).toEqual(new Error('no whitespace characters allowed'))
  })
});
