// Rust vectors return an Option when retrieving an element, unlike Go that will panic
// if accessing an out of bounds element. No need to perform if based lookups
const NEIGHBOURS: [(isize, isize); 4] = [(0, -1), (0, 1), (-1, 0), (1, 0)];

pub fn run() -> usize {
    let input = include_str!("./input.txt");
    return solve(input)
}

fn solve(input: &str) -> usize {
    let grid = input
        .as_bytes()
        .split(|&b| b == b'\n')
        .collect::<Vec<_>>();

    let mut total = 0;
    for (y, line) in grid.iter().enumerate() {
        for (x, location) in line.iter().enumerate() {
            if NEIGHBOURS.iter().all(|&(xx, yy)| {
                grid
                    .get((y as isize + yy) as usize)
                    .and_then(|l| l.get((x as isize + xx) as usize))
                    .map(|n| location < n)
                    .unwrap_or(true)
                }) {
                total += (location - b'0') as usize + 1
            }
        }
    }

    return total
}

#[cfg(test)]
mod tests {
    extern crate test;

    use test::Bencher;
    use super::*;

    #[test]
    fn test_example() {
        let input = r"2199943210
3987894921
9856789892
8767896789
9899965678";
        assert_eq!(solve(input), 15);
    }

    #[test]
    fn test_puzzle() {
        assert_eq!(run(), 480);
    }

    #[bench]
    fn bench_run(b: &mut Bencher) {
        b.iter(|| run())
    }
}