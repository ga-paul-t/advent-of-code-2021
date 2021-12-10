#![feature(test)]

pub fn main() {
    let input = include_str!("./input.txt");
    print!("{}", day01a(input));
}

pub fn day01a(input: &str) -> usize {
    return input
        .lines()
        .map(|n| n.parse().unwrap())
        .collect::<Vec<u16>>()
        .windows(2)
        .filter(|w| w[1] > w[0])
        .count();
}

#[cfg(test)]
mod tests {
    extern crate test;

    use test::Bencher;
    use super::*;

    #[test]
    fn test_day01a() {
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
        assert_eq!(day01a(input), 7);
    }

    #[bench]
    fn bench_main(b: &mut Bencher) {
        b.iter(|| main())
    }
}