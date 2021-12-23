pub fn run() -> usize {
    let input = include_str!("./input.txt");
    return solve(input, 80)
}

fn solve(input: &str, days: usize) -> usize {
    let mut lifecycle = input
        .split(",")
        .fold([0; 9], |mut init, v| {
            init[v.parse::<usize>().unwrap()] += 1;
            init
        });

    // Iterate over the days calculating the lifecycle
    (1..=days)
        .for_each(|_| {
            let new_fish = lifecycle[0];
            (1..=8).for_each(|i| lifecycle[i-1] = lifecycle[i]);

            lifecycle[8] = new_fish;
            lifecycle[6] += new_fish;
        });

    return lifecycle.iter().sum::<usize>();
}

#[cfg(test)]
mod tests {
    extern crate test;

    use test::Bencher;
    use super::*;

    #[test]
    fn test_example() {
        let input = "3,4,3,1,2";
        assert_eq!(solve(input, 80), 5934);
    }

    #[test]
    fn test_puzzle() {
        assert_eq!(run(), 350149);
    }

    #[bench]
    fn bench_run(b: &mut Bencher) {
        b.iter(|| run())
    }
}