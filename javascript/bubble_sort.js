function bubble_sort(arr) {
    arr.forEach(() => {
        arr.slice(1).forEach((el, idx) => {
            if (arr[idx] > arr[idx + 1]) {
                arr[idx + 1] = arr[idx];
                arr[idx] = el;
            }
        })
    });
    return arr;
}

const array = [2,1,6,4,5,3];

console.log("Original", array);
console.log("Sorted", bubble_sort([...array]));
