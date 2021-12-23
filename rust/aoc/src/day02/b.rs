pub fn run() -> usize {
    let input = include_str!("./input.txt");
    return solve(input);
}


pub fn solve(input: &str) -> usize {
    let (hpos, dep, _) =  input
        .lines()
        .map(|cmd| {
            let s = cmd.split_once(" ").unwrap();
            return (s.0, s.1.parse::<usize>().unwrap());
        })
        .fold((0, 0, 0), |(hpos, dep, aim), (cmd, v)| {
            match (cmd, v) {
                ("forward", v) => (hpos + v, dep + (aim * v), aim),
                ("down", v) => (hpos, dep, aim + v),
                _ => (hpos, dep, aim - v),
            }
        });

    return hpos * dep;
}

#[cfg(test)]
mod tests {
    extern crate test;

    use test::Bencher;
    use super::*;

    #[test]
    fn test_example() {
        let input = r"forward 5
down 5
forward 8
up 3
down 8
forward 2";
        assert_eq!(solve(input), 900);
    }

    #[test]
    fn test_puzzle() {
        assert_eq!(run(), 1975421260);
    }

    #[bench]
    fn bench_run(b: &mut Bencher) {
        b.iter(|| run())
    }
}