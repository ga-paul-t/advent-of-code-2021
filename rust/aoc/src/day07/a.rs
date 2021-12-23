pub fn run() -> usize {
    let input = include_str!("./input.txt");
    return solve(input)
}


fn solve(input: &str) -> usize {
    let mut hpos = input
        .split(',')
        .map(|s| s.parse::<i32>().unwrap())
        .collect::<Vec<_>>();

    // Determine the median from the set of horizontal positions
    let midpoint = hpos.len() / 2;
    let median = *hpos.select_nth_unstable(midpoint).1;

    return hpos.iter()
        .map(|h| {
            (h - median).abs()
        })
        .sum::<i32>() as usize;
}

#[cfg(test)]
mod tests {
    extern crate test;

    use test::Bencher;
    use super::*;

    #[test]
    fn test_example() {
        let input = "16,1,2,0,4,2,7,1,2,14";
        assert_eq!(solve(input), 37);
    }

    #[test]
    fn test_puzzle() {
        assert_eq!(run(), 329389);
    }

    #[bench]
    fn bench_run(b: &mut Bencher) {
        b.iter(|| run())
    }
}