let sort = (arr) => {
    let cleaned = _clean(arr);

    if (cleaned.length < 2) {
        return cleaned;
    }

    return _quickSort(arr);
};

let _quickSort = (arr) => {
    let separator = arr[0];
    let lessThan = [];
    let greaterThan = [];

    arr.forEach((int) => {
        lessThan = _buildLessThan(lessThan, separator, int);
        greaterThan = _buildGreaterThan(greaterThan, separator, int);
    });

    return sort(lessThan).concat(separator).concat(sort(greaterThan));
};

let _buildLessThan = (arr, separator, int) => {
    if (int < separator) {
        arr.push(int);
    }
    return arr;
};

let _buildGreaterThan = (arr, separator, int) => {
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
