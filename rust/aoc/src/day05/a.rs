use atoi::atoi;
use nom::*;

pub fn run() -> usize {
    let input = include_str!("./input.txt");
    return solve(input, 1000)
}

fn solve(input: &str, dimension: usize) -> usize {
    let mut grid = vec![0u8; dimension * dimension];
    let mut overlaps = 0;

    // Define a function for plotting a point on the grid and incrementing overlap count
    let mut plot = |x: usize, y: usize| {
        let pos = x + y*dimension;

        if grid[pos] == 1 {
            overlaps += 1;
        }

        grid[pos] += 1;
    };

    // Parse as bytes to allow nom library to chomp away
    input
        .as_bytes()
        .split(|c| *c == b'\n' )
        .map(|l| {
            let ((x1,y1), (x2,y2)) = line(l).unwrap().1;

            // Simplify calculations by ordering all points to ensure steps are always positive.
            // We are ignoring diagonals here
            return if x1 == x2 && y1 > y2 {
                ((x1, y2), (x2, y1))
            } else if y1 == y2 && x1 > x2 {
                ((x2, y1), (x1, y2))
            } else {
                ((x1, y1), (x2, y2))
            }
        })
        .for_each(|((x1, y1), (x2, y2) )| {
            if x1 == x2 {
                (y1..=y2).for_each(|py| plot(x1, py));
            } else if y1 == y2 {
                (x1..=x2).for_each(|px| plot(px, y1));
            }
        });

    return overlaps;
}

// This is pretty amazing, love learning rust, https://docs.rs/nom/6.1.2/nom/macro.named.html
named!(coord<&[u8], usize>, map_opt!(nom::character::complete::digit1, atoi));
named!(coords<&[u8], (usize, usize)>, separated_pair!(coord, char!(','), coord));
named!(line<&[u8], ((usize, usize), (usize, usize))>, separated_pair!(coords, tag!(" -> "), coords));

#[cfg(test)]
mod tests {
    extern crate test;

    use test::Bencher;
    use super::*;

    #[test]
    fn test_solve() {
        let input = r"0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2";
        assert_eq!(solve(input, 10), 5);
    }

    #[bench]
    fn bench_run(b: &mut Bencher) {
        b.iter(|| run())
    }
}