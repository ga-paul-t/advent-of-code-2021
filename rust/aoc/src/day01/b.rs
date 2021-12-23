pub fn run() -> usize {
    let input = include_str!("./input.txt");
    return solve(input);
}


pub fn solve(input: &str) -> usize {
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
    fn test_example() {
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
        assert_eq!(solve(input), 5);
    }

    #[test]
    fn test_puzzle() {
        assert_eq!(run(), 1248);
    }

    #[bench]
    fn bench_run(b: &mut Bencher) {
        b.iter(|| run())
    }
}