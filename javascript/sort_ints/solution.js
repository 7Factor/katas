let sort = (arr) => {

    let cleaned = _clean(arr);
    if (cleaned.length < 2) {
        return cleaned;
    }

    let separator = arr[0];
    let lessThan = [];
    let greaterThan = [];
    cleaned.forEach((int) => {
        lessThan = buildLessThan(lessThan, separator, int);
        greaterThan = buildGreaterThan(greaterThan, separator, int);
    });

    return lessThan.concat(separator).concat(greaterThan);
};

let buildLessThan = (arr, separator, int) => {
    if (int < separator) {
        arr.push(int);
    }
    return arr;
};

let buildGreaterThan = (arr, separator, int) => {
    if (int > separator) {
        arr.push(int)
    }
    return arr;
};

let _clean = (arr) => {
    if (arr == null) {
        return [];
    }

    return arr;
};


module.exports = {
    sort
};
