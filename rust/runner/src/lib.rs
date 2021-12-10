pub fn jobs() -> &'static [(fn(), &'static str)] {
    &[
        (day01a::main, "01a"),
        (day01b::main, "01b")
    ]
}