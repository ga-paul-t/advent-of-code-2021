use nom::InputIter;

pub fn run() -> usize {
    let input = include_str!("./input.txt");
    return solve(input)
}


fn solve(input: &str) -> usize {
    let total = input
        .lines()
        .map(|line| {
            let parts = line.split_once("|").unwrap();
            let mut sigs = parts.0.split_ascii_whitespace();
            let digits = parts.1.split_ascii_whitespace();

            // Digits 1, 4, 7 and 8 have a unique number of segments. All other digits have either
            // 5 [2, 3 and 5] or 6 [0, 6, and 9] segments. All digits based on the illustration contain
            // parts of either digit 1 or 4. By generating a lookup table, it should be possibly to decode
            // the other digits
            let one = sigs.clone().find(|s| s.len() == 2).unwrap_or_default();
            let four = sigs.find(|s| s.len() == 4).unwrap_or_default();

            digits.map(|d| {
                match d.len() {
                    2 => 1,
                    3 => 7,
                    4 => 4,
                    7 => 8,
                    // Mappings:
                    // 3 :> 4: [3 parts] - 1: [2 parts] - Len [5]
                    // 5 :> 4: [3 parts] - 1: [*]       - Len [5]
                    // 2 :> 4: [*]       - 1: [*]       - Len [5]
                    // 9 :> 4: [4 parts] - 1: [*]       - Len [6]
                    // 0 :> 4: [3 parts] - 1: [2 parts] - Len [6]
                    // 6 :> 4: [3 parts] - 1: [*]       - Len [6]
                    len=> match (len,
                    d.iter_elements().filter(|&c| one.contains(c)).count(),
                    d.iter_elements().filter(|&c| four.contains(c)).count()) {
                        (5, 2, 3) => 3,
                        (5, _, 3) => 5,
                        (5, _, _) => 2,
                        (6, _, 4) => 9,
                        (6, 2, 3) => 0,
                        (6, _, 3) => 6,
                        _ => unreachable!(),
                    },
                }
            })
                .enumerate()
                .fold(0, |acc, (i, n)| acc + n * 10_u32.pow(3 - i as u32))
        })
        .sum::<u32>();

    return total as usize
}

#[cfg(test)]
mod tests {
    extern crate test;

    use test::Bencher;
    use super::*;

    #[test]
    fn test_example() {
        let input = r"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce";

        assert_eq!(solve(input), 61229);
    }

    #[test]
    fn test_puzzle() {
        assert_eq!(run(), 994266);
    }

    #[bench]
    fn bench_run(b: &mut Bencher) {
        b.iter(|| run())
    }
}