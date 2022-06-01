function quicksort(arr, partitionFn) {
    if (arr.length > 1) {
        const p = partitionFn(arr);
        return [...quicksort(arr.slice(0, p), partitionFn), ...quicksort(arr.slice(p), partitionFn)];
    }
    return arr;
}

function partitionLeft(arr) {
    const pivotIdx = 0;
    const pivot = arr[pivotIdx];

    let i = arr.length - 1;
    for (let j = i; j > 0; j--) {
        if (arr[j] > pivot) {
            const aux = arr[j];
            arr[j] = arr[i];
            arr[i] = aux;
            i--;
        }
    }

    const aux = arr[i];
    arr[i] = pivot;
    arr[pivotIdx] = aux;

    return i + 1;
}

function partitionRight(arr) {
    const pivotIdx = arr.length - 1;
    const pivot = arr[pivotIdx];

    let i = 0;
    arr.forEach((el, idx) => {
        if (el < pivot) {
            arr[idx] = arr[i];
            arr[i] = el;
            i++;
        }
    });

    const aux = arr[i];
    arr[i] = pivot;
    arr[pivotIdx] = aux;

    return i;
}

function partition(arr) {
    let [left, right] = [0, arr.length - 1];
    const pivot = arr[Math.floor(arr.length / 2)];

    while (1) {
        while (arr[left] < pivot) {
            left++;
        }
        while (arr[right] > pivot) {
            right--;
        }

        if (left >= right) {
            break;
        }

        const aux = arr[left];
        arr[left] = arr[right];
        arr[right] = aux;
    }

    return right;
}

const array = [2,1,6,4,5,3];
console.log("Original", array);
console.log("Sorted (Left)", quicksort([...array], partitionLeft));
console.log("Sorted (Right)", quicksort([...array], partitionRight));
console.log("Sorted (Middle)", quicksort([...array], partition));
