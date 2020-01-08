let sort = (arr) => {
    if (arr == null) {
        return [];
    }

    if (arr.length < 2) {
        return arr;
    }

    return _buildNewArray(arr);
};

let _buildNewArray = (arr) => {
    let first = arr[0];
    let second = arr[1];
    if (first<second) {
        return [first, second];
    } else {
        return [second, first];
    }
};

module.exports = {
    sort
};
