#![feature(test)]

pub mod day01;

pub fn jobs() -> &'static [(fn() -> usize, &'static str)] {
    &[
        (day01::a::run, "01a"),
        (day01::b::run, "01b"),
    ]
}