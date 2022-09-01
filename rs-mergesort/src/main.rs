mod lib;
use crate::lib::*;

fn main() {
    merge_sort(random_vec(100_000_000));
}
