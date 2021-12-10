use took::Timer;

fn main() {
    println!("🎄 Advent of Code 2021");
    println!();

    aoc::jobs()
        .iter()
        .for_each(|job| {
            let now = Timer::new();
            println!("🧩 Puzzle {}: {} [time taken: {}]", job.1, job.0(), now.took().to_string())
        });
}