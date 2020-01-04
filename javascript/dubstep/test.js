const songDecoder = require('./solution');

describe('Dubstep', () => {
  test('should not accept a null input', () => {
    let actual = songDecoder.decode(null);
    expect(actual).toEqual('null is not allowed');
  });
});
