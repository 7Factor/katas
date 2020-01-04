let decode = (song) => {
  let err = _catchErrors(song);
  if (err != null) {
    return err;
  }

  let wubRe = new RegExp(/(?<name>WUB)+/g);
  return song.replace(wubRe, ' ').trim();
};

let _catchErrors = (input) => {
  let lowerCaseRe = new RegExp('[a-z]');
  let spaceRe = new RegExp('[\\s]');

  if (input === null) {
    return new Error('null is not allowed')
  }

  if (input.match(lowerCaseRe)) {
    return new Error('lower case letters are not allowed')
  }

  if (input.match(spaceRe)) {
    return new Error('no whitespace characters allowed')
  }

  if (input.length > 200) {
    return new Error('inputs must be 200 characters or less');
  }
};

module.exports = {
  decode
};