pub fn run() -> usize {
    let input = include_str!("./input.txt");
    return solve(input)
}

fn solve(input: &str) -> usize {
    let readings = input
        .lines()
        .collect::<Vec<&str>>();

    let max = (readings.len() as f32 / 2.0).ceil() as usize;
    let bits = readings[0].len();

    // Bit shift gamma as required
    let mut gamma = 0;

    readings
        .into_iter()
        .fold(vec![0; bits], |count, reading| {
            // Count all occurrences of 1 at each bit position within a reading
            count
                .into_iter()
                .enumerate()
                .map(|(i, n)| {
                    return if reading.as_bytes()[i] == b'1' {
                        n + 1
                    } else {
                        n
                    }
                })
                .collect()
        })
        .into_iter()
        .enumerate()
        .for_each(|(i, b)| {
            if b >= max {
                gamma += 1 << (bits - (i + 1))
            }
        });

    let epsilon = !gamma & ((1 << bits) - 1);
    return gamma * epsilon;
}

#[cfg(test)]
mod tests {
    extern crate test;

    use test::Bencher;
    use super::*;

    #[test]
    fn test_example() {
        let input = r"00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010";
        assert_eq!(solve(input), 198);
    }

    #[test]
    fn test_puzzle() {
        assert_eq!(run(), 3882564);
    }

    #[bench]
    fn bench_run(b: &mut Bencher) {
        b.iter(|| run())
    }
}