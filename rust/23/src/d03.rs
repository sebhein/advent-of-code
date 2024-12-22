use regex::Regex;

pub fn solve(_input: &String) {
    let sol_1 = _solve_p1(_input);
    println!("day 03, part 1, {sol_1}");
    //let sol_2 = _solve_p2(_input);
    //println!("day 03, part 2, {sol_2}");
}

fn _solve_p1(_input: &String) -> usize {
    let re = Regex::new(r"\*|#|\+|\$|\/|=|%|-").unwrap();
    for line in _input.lines() {
        for cap in re.captures_iter(line) {
            let symbol = cap.get(0).unwrap();
            println!("{}", symbol.as_str());
        }
    }
    return 0;
}
