// Rust vectors return an Option when retrieving an element, unlike Go that will panic
// if accessing an out of bounds element. No need to perform if based lookups
const NEIGHBOURS: [(isize, isize); 4] = [(0, -1), (0, 1), (-1, 0), (1, 0)];

pub fn run() -> usize {
    let input = include_str!("./input.txt");
    return solve(input)
}

fn solve(input: &str) -> usize {
    let mut grid = input
        .as_bytes()
        .split(|&b| b == b'\n')
        .map(|l| l.to_vec())
        .collect::<Vec<_>>();

    let mut basins = vec![];
    for y in 0..grid.len() {
        for x in 0..grid[0].len() {
            if grid[y][x] < b'9' {
                basins.push(search_basin(&mut grid, x, y));
            }
        }
    }

    // Sort the slice and take the largest 3 values
    basins.sort_unstable();

    return basins.iter()
        .rev()
        .take(3)
        .product::<usize>()
}

fn search_basin(grid: &mut Vec<Vec<u8>>, x: usize, y: usize) -> usize {
    // Recursively search all neighbours, marking each visited point with a 9
    // to ensure it is not traversed again
    grid[y][x] = b'9';

    NEIGHBOURS.iter()
        .map(|(xx, yy)| ((x as isize + xx) as usize, (y as isize + yy) as usize))
        .fold(1, |acc, (x, y)| {
            match grid
                .get(y)
                .and_then(|l| l.get(x))
                .map(|&n| n < b'9') {
                Some(true) => acc + search_basin(grid, x, y),
                _ => acc,
            }
        })
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
        assert_eq!(solve(input), 1134);
    }

    #[test]
    fn test_puzzle() {
        assert_eq!(run(), 1045660);
    }

    #[bench]
    fn bench_run(b: &mut Bencher) {
        b.iter(|| run())
    }
}