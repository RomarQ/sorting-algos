fn main() {
	let list = [5, 1, 4, 6, 3, 2].to_vec();

    eprintln!("Merge Sort");
    eprintln!("Original {:?}", list);
    eprintln!("Sorted {:?}", merge_sort(list.clone()));
}

fn merge_sort<T: Ord + Copy>(list: Vec<T>) -> Vec<T> {
    if list.len() > 1 {
        let middle = list.len() / 2;

        let l = merge_sort(list[..middle].to_vec());
        let r = merge_sort(list[middle..].to_vec());

        return merge(l, r)
    }

    return list
}

fn merge<T: Ord + Copy>(mut l: Vec<T>, mut r: Vec<T>) -> Vec<T> {
    let mut new_vec = vec![];

    while l.len() > 0 && r.len() > 0 {
        if l[0] < r[0] {
            new_vec.push(l.remove(0));
        } else {
            new_vec.push(r.remove(0));
        }
    }

    new_vec.append(&mut l);
    new_vec.append(&mut r);

    return new_vec;
}
