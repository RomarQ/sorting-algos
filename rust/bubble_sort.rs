fn main() {
	let list = [5, 1, 4, 6, 3, 2].to_vec();

    eprintln!("Bubble Sort");
    eprintln!("Original {:?}", list);
    eprintln!("Sorted {:?}", bubble_sort(list.clone()));
}

fn bubble_sort<T: Ord + Copy>(mut list: Vec<T>) -> Vec<T> {
    for i in 0 .. list.len() {
        for j in i .. list.len() {
            if list[i] > list[j] {
                list.swap(i, j);
            }
        }
    }

    return list;
}
