pub fn run() -> usize {
    let input = include_str!("./input.txt");
    return solve(input)
}


fn solve(input: &str) -> usize {
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
        assert_eq!(solve(input), 7);
    }

    #[test]
    fn test_puzzle() {
        assert_eq!(run(), 1298);
    }

    #[bench]
    fn bench_run(b: &mut Bencher) {
        b.iter(|| run())
    }
}