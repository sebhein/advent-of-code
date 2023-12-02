use regex::Regex;

pub fn solve(_input: &String) {
    let sol_1 = _solve(_input);
    println!("day 01, part 1, {sol_1}");
    //let sol_2 = _solve(_input);
    //println!("day 01, part 2, {sol_2}");
}

fn _solve(_input: &String) -> usize {
    let re = Regex::new(r"\d").unwrap();
    let mut sum = 0;
    for line in _input.lines() {
        let digits: Vec<_> = re.find_iter(line).map(|m| m.as_str()).collect();
        let first = digits[0];
        let last = digits[digits.len() - 1];
        let digit = format!("{}{}", first, last).parse::<usize>().unwrap();
        sum += digit;
    }
    return sum;
}
