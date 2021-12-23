pub fn run() -> usize {
    let input = include_str!("./input.txt");
    return solve(input)
}


fn solve(input: &str) -> usize {
    let hpos = input
        .split(',')
        .map(|s| s.parse().unwrap())
        .collect::<Vec<_>>();

    // Due to rounding around the mean, sample both bounds and return the smallest total
    let mean = hpos.iter().sum::<i32>() / hpos.len() as i32..;

    return mean
        .take(2)
        .map(|m| {
            hpos.iter()
                .map(|h| {
                    let dist = (h - m).abs();
                    dist * (dist + 1) / 2
                })
                .sum::<i32>()
        })
        .min()
        .unwrap() as usize;
}

#[cfg(test)]
mod tests {
    extern crate test;

    use test::Bencher;
    use super::*;

    #[test]
    fn test_example() {
        let input = "16,1,2,0,4,2,7,1,2,14";
        assert_eq!(solve(input), 168);
    }

    #[test]
    fn test_puzzle() {
        assert_eq!(run(), 86397080)
    }

    #[bench]
    fn bench_run(b: &mut Bencher) {
        b.iter(|| run())
    }
}