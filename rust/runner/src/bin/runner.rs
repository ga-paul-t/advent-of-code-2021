use took::Timer;

fn main() {
    println!("🎄 Advent of Code 2021");
    println!();

    runner::jobs()
        .iter()
        .for_each(|j| {
            print!("🧩 Puzzle {}: ", j.1);

            let now = Timer::new();
            j.0();
            println!(" [time taken: {}]", now.took().to_string())
        });
}