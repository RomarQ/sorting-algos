function merge_sort(arr) {
    if (arr.length > 1) {
        const middle = arr.length / 2;

        const left = merge_sort(arr.slice(0, middle))
        const right = merge_sort(arr.slice(middle))

        return merge(left, right)
    }

    return arr;
}

function merge(left, right) {
    const sorted = [];

    while (left.length && right.length) {
        if (left[0] < right[0]) {
            sorted.push(left.shift())
        } else {
            sorted.push(right.shift())
        }
    }

    return sorted.concat([...left,...right])
}

const array = [2,1,6,4,5,3];

console.log("Original", array);
console.log("Sorted", merge_sort([...array]));
