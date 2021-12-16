use std::collections::HashMap;

const BOARD_DIMESION: usize = 5;
const BOARD_SIZE: usize = 25;

pub fn run() -> usize {
    let input = include_str!("./input.txt");
    return solve(input)
}


fn solve(input: &str) -> usize {
    let (nums, boards) = input.split_once("\n\n").unwrap();

    // Parse the boards into a structure containing both a map (for fast lookup) and a
    // vector for marking off each called number
    let mut boards: Vec<(HashMap<u8, usize>, Vec<char>)> = boards
        .split("\n\n")
        .map(|board| {
            (
                board.split_ascii_whitespace()
                .enumerate()
                .map(|(index, num)| {
                    (num.parse().unwrap(), index)
                })
                .collect(),
                vec!['0'; BOARD_SIZE],
            )
        })
        .collect();

    // Iterate through each called number, marking the boards until the first board wins
    let (board, num) = nums
        .split(",")
        .map(|num| num.parse().unwrap())
        .find_map(|num| {
            // Loop over the boards, marking as required
            boards
                .iter_mut()
                .find_map(|(board, marks)| {
                    // Check if the board has won
                    board
                        .clone()
                        .get(&num)
                        .map(|pos| {
                            marks[*pos] = '1';
                            board.remove(&num);
                            *pos
                        })
                        .filter(|p|
                            {
                                // Calculate positions within the board
                                let row = p / BOARD_DIMESION;
                                let row_pos = row * BOARD_DIMESION;

                                let col = p % BOARD_DIMESION;
                                let col_pos = col + BOARD_DIMESION;

                                let matched = marks[row_pos] == '1' &&
                                    marks[row_pos+1] == '1' &&
                                    marks[row_pos+2] == '1' &&
                                    marks[row_pos+3] == '1'
                                    && marks[row_pos+4] == '1';

                               return matched || marks[col] == '1' &&
                                   marks[col_pos] == '1' &&
                                   marks[col_pos*2] == '1' &&
                                   marks[col_pos*3] == '1' &&
                                   marks[col_pos*4] == '1';
                            })
                        .map(|_| (board.clone(), num))
                })
        })
        .unwrap();

    return board.into_iter()
        .map(|(k, _)| k as usize)
        .sum::<usize>() * num as usize;
}

#[cfg(test)]
mod tests {
    extern crate test;

    use test::Bencher;
    use super::*;

    #[test]
    fn test_solve() {
        let input = r"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7";
        assert_eq!(solve(input), 4512);
    }

    #[bench]
    fn bench_run(b: &mut Bencher) {
        b.iter(|| run())
    }
}