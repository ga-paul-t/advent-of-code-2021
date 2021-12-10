#![feature(test)]

pub fn main() {
    let input = include_str!("./input.txt");
    println!("{}", day01b(input));
}

pub fn day01b(input: &str) -> usize {
    // Only need to compare the first and last numbers, as each sliding window
    // will share 2 common values:
    //
    // [199, 200, 208, 210] -> A: [199, 299, 208] B: [200, 208, 210]
    // 210 > 199 ++(increased)
    return input
        .lines()
        .map(|n| n.parse().unwrap())
        .collect::<Vec<u16>>()
        .windows(4)
        .filter(|w| w[3] > w[0])
        .count();
}

#[cfg(test)]
mod tests {
    extern crate test;

    use test::Bencher;
    use super::*;

    #[test]
    fn test_day01b() {
        let input = r"199
200
208
210
200
207
240
269
260
263";
        assert_eq!(day01b(input), 5);
    }

    #[bench]
    fn bench_main(b: &mut Bencher) {
        b.iter(|| main())
    }
}