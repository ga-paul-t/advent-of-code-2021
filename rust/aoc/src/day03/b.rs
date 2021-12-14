pub fn run() -> usize {
    let input = include_str!("./input.txt");
    return solve(input)
}

fn solve(input: &str) -> usize {
    let readings = input
        .lines()
        .collect::<Vec<&str>>();

    let bits = readings[0].len();

    let oxygen_rating = (0..bits)
        .scan(readings.clone(), | oxygen, i| {
            let max = (oxygen.len() as f32 / 2.0).ceil() as usize;

            let mut sig = b'0';
            let count = oxygen.iter()
                .filter(|o| (**o).as_bytes()[i] == b'1')
                .count();
            if count >= max {
                sig = b'1';
            }

            oxygen.drain_filter(|o| (*o).as_bytes()[i] != sig);
            oxygen.first().copied()
        })
        .last()
        .unwrap();

    let co2_rating = (0..bits)
        .scan(readings, | co2, i| {
            let max = (co2.len() as f32 / 2.0).ceil() as usize;

            let mut sig = b'1';
            let count = co2.iter()
                .filter(|c| (**c).as_bytes()[i] == b'1')
                .count();
            if count >= max {
                sig = b'0';
            }

            co2.drain_filter(|c| (*c).as_bytes()[i] != sig);
            co2.first().copied()
        })
        .last()
        .unwrap();

    return usize::from_str_radix(oxygen_rating, 2).unwrap() *
        usize::from_str_radix(co2_rating, 2).unwrap();
}

#[cfg(test)]
mod tests {
    extern crate test;

    use test::Bencher;
    use super::*;

    #[test]
    fn test_solve() {
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
        assert_eq!(solve(input), 230);
    }

    #[bench]
    fn bench_run(b: &mut Bencher) {
        b.iter(|| run())
    }
}