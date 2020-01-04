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
  });

  test('should not accept input longer than 200 characters', () => {
    let longInput = _randomString(201);
    let actual = songDecoder.decode(longInput);
    expect(actual).toEqual(new Error('inputs must be 200 characters or less'))
  });

  test('should remove a single instance of `WUB`', () => {
    let actual = songDecoder.decode('WUB');
    expect(actual).toEqual('');
  });

  test('should remove multiple instances of `WUB`', () => {
    let actual = songDecoder.decode('WUBWUB');
    expect(actual).toEqual('');
  });

  test('should handle an encoding of `WUB` before a letter', () => {
    let actual = songDecoder.decode('WUBT');
    expect(actual).toEqual('T')
  });

  test('should handle an encoding of `WUB` before one word', () => {
    let actual = songDecoder.decode('WUBTHE');
    expect(actual).toEqual('THE')
  });

  test('should handle an encoding of `WUB` before each word', () => {
    let actual = songDecoder.decode('WUBTHEWUBKOMBUCHA');
    expect(actual).toEqual('THE KOMBUCHA')
  });

  test('should handle multiple encodings of `WUB` in between words and letters', () => {
    let actual = songDecoder.decode('AWUBWUBB');
    expect(actual).toEqual('A B')
  })
});

let _randomString = (length) => {
  const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
  let result = '';
  for (let i = length; i > 0; --i) {
    result += chars[Math.floor(Math.random() * chars.length)];
  }
  return result;
};
