use atoi::atoi;
use nom::*;
use std::iter;

pub fn run() -> usize {
    let input = include_str!("./input.txt");
    return solve(input, 1000)
}

fn solve(input: &str, dimension: usize) -> usize {
    let (mut grid, mut overlaps) = (vec![0u8; dimension * dimension], 0);

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
            ((x1,y1),(x2,y2))
        })
        .for_each(|((x1, y1), (x2, y2))| {
            // By creating a custom iterator, we can easily range over both x and y dimensions.
            // Tried this in Go, but consumed too much memory and slowed execution time
            let next =
                |a: isize, b: isize| iter::successors(Some(a), move |pos| {
                    Some(pos + (b - a).signum())
                });

            // Plot the final coordinate, and then range over while plotting the rest
            plot(x2 as usize, y2 as usize);

            next(x1, x2)
                .zip(next(y1, y2))
                .take_while(|(x,y)| *x != x2 || *y != y2)
                .for_each(|(x,y)| {
                    plot(x as usize, y as usize);
                })
        });

    return overlaps;
}

named!(coord<&[u8], isize>, map_opt!(nom::character::complete::digit1, atoi));
named!(coords<&[u8], (isize, isize)>, separated_pair!(coord, char!(','), coord));
named!(line<&[u8], ((isize, isize), (isize, isize))>, separated_pair!(coords, tag!(" -> "), coords));

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
        assert_eq!(solve(input, 10), 12);
    }

    #[bench]
    fn bench_run(b: &mut Bencher) {
        b.iter(|| run())
    }
}