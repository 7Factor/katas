let sort = (arr) => {
    if (arr == null) {
        return [];
    }

    if (arr.length < 2) {
        return arr;
    }

    let separator = arr[0];
    let lessThan = [];
    let greaterThan = [];
    arr.forEach((int) => {
        if (int < separator) {
            lessThan.push(int);
        }
        if (int > separator) {
            greaterThan.push(int)
        }
    });

    return lessThan.concat(separator).concat(greaterThan);
};

module.exports = {
    sort
};
