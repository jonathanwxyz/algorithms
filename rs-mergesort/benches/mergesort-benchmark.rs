use rs_mergesort::{merge_sort, random_vec};

use criterion::{
    black_box,
    criterion_group,
    criterion_main,
    Criterion
};


fn merge_sort_benchmark(c: &mut Criterion) {
    let arr = black_box(
        random_vec(10000)  
    );
    c.bench_function(
        "merge sort",
        |b| b.iter(|| merge_sort(arr.clone()))
    );

}

criterion_group!(benches, merge_sort_benchmark);
criterion_main!(benches);
