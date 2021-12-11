pub fn run() -> usize {
    let input = include_str!("./input.txt");
    return solve(input)
}


fn solve(input: &str) -> usize {
    let (hpos, dep) =  input
        .lines()
        .map(|cmd| {
            let s = cmd.split_once(" ").unwrap();
            return (s.0, s.1.parse::<usize>().unwrap());
        })
        .fold((0, 0), |(hpos, dep), (cmd, v)| {
            match (cmd, v) {
                ("forward", v) => (hpos + v, dep),
                ("down", v) => (hpos, dep + v),
                _ => (hpos, dep - v),
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
    fn test_solve() {
        let input = r"forward 5
down 5
forward 8
up 3
down 8
forward 2";
        assert_eq!(solve(input), 150);
    }

    #[bench]
    fn bench_run(b: &mut Bencher) {
        b.iter(|| run())
    }
}