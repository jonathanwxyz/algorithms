use rand::Rng;

pub fn random_vec(size: usize) -> Vec<isize> {
    let mut rng = rand::thread_rng();
    let mut arr = Vec::with_capacity(size);

    for _ in 0..size {
        arr.push(rng.gen());
    }
    arr
}

pub fn merge_sort<T>(arr: Vec<T>) -> Vec<T>
    where T: Ord + Copy + Send + Sync {
    if arr.len() <= 1 { return arr }
    let middle = arr.len() / 2;
    let (left, right) = rayon::join(|| merge_sort(arr[..middle].to_vec()),
                                    || merge_sort(arr[middle..].to_vec()));
    merge(&left, &right)
}

pub fn seq_merge_sort<T: Ord + Copy>(arr: Vec<T>) -> Vec<T> {
    if arr.len() <= 1 { return arr }
    let middle = arr.len() / 2;
    merge(&seq_merge_sort(arr[..middle].to_vec()),
            &seq_merge_sort(arr[middle..].to_vec()))
}

pub fn con_merge_sort<T>(arr: Vec<T>) -> Vec<T> 
    where T: 'static + Ord + Copy + Send + Sync {
    if arr.len() < 512 { return merge_sort(arr) }
    let num_cpus = num_cpus::get();
    parallel_merge_sort(arr, num_cpus)
}

fn parallel_merge_sort<T>(arr: Vec<T>, n: usize) -> Vec<T>
    where T: 'static + Ord + Copy + Send + Sync {
    if arr.len() <= 1 { return arr }

    let middle = arr.len() / 2;
    if n <= 2 {return merge(&merge_sort(arr[..middle].to_vec()),
                &merge_sort(arr[middle..].to_vec()))}

    let half1= arr[..middle].to_vec();
    let half2 = arr[middle..].to_vec();
    let handle1 = std::thread::spawn(move || {
        parallel_merge_sort(half1, n - 2)});
    let handle2 = std::thread::spawn(move || {
        parallel_merge_sort(half2, n - 2)});

    let ms1 = handle1.join().unwrap();
    let ms2 = handle2.join().unwrap();

    return merge(&ms1, &ms2)
}

fn merge<T: Ord + Copy> (a: &[T], b: &[T]) -> Vec<T> {
    let mut sorted: Vec<T> = Vec::with_capacity(a.len() + b.len());
    let mut i = 0;
    let mut j = 0;
    while i != a.len() || j != b.len() {
        if i == a.len() {
            sorted.extend_from_slice(&b[j..]);
            break;
        } else if j == b.len() {
            sorted.extend_from_slice(&a[i..]);
            break;
        } else if a[i] < b[j] {
            sorted.push(a[i]);
            i += 1;
        } else {
            sorted.push(b[j]);
            j += 1;
        }
    }
    sorted
}

