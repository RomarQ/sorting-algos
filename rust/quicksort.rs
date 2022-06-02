
fn main() {
	let list = [5, 1, 4, 6, 3, 2].to_vec();

    eprintln!("Quicksort");
    eprintln!("Original {:?}", list);
    eprintln!("Sorted {:?}", quicksort(&mut list.clone()));
}

fn quicksort<T: Ord + Copy>(vector: &mut [T]) -> &[T] {
    if vector.len() > 1 {
        let p = partition(vector);

        quicksort(&mut vector[..p]);
        quicksort(&mut vector[p..]);
    }
    return vector;
}

fn partition<T: Ord + Copy>(vector: &mut [T]) -> usize {
    let (mut left, mut right) = (0, vector.len()-1);
    let pivot = vector[vector.len() / 2];

    loop {
        while vector[left] < pivot {
            left+=1;
        }

        while vector[right] > pivot {
            right-=1;
        }

        if left >= right {
            break;
        }

        vector.swap(left, right)
    }

    return right
}
