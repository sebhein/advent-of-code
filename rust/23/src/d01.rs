use regex::Regex;

pub fn solve(_input: &String) {
    let sol_1 = _solve_p1(_input);
    println!("day 01, part 1, {sol_1}");
    let sol_2 = _solve_p2(_input);
    println!("day 01, part 2, {sol_2}");
}

fn _solve_p1(_input: &String) -> usize {
    let re = Regex::new(r"\d").unwrap();
    let mut sum = 0;
    for line in _input.lines() {
        let digits: Vec<_> = re.find_iter(line).map(|m| m.as_str()).collect();
        if digits.len() == 0 {
            continue;
        }
        let first = digits[0];
        let last = digits[digits.len() - 1];
        let digit = format!("{}{}", first, last).parse::<usize>().unwrap();
        sum += digit;
    }
    return sum;
}

fn _to_string_digit(digit: &str) -> &str {
    match digit {
        "one" => "1",
        "two" => "2",
        "three" => "3",
        "four" => "4",
        "five" => "5",
        "six" => "6",
        "seven" => "7",
        "eight" => "8",
        "nine" => "9",
        _ => digit,
    }
}

fn _solve_p2(_input: &String) -> usize {
    let find = vec![r"\d", r"one", r"two", r"three", r"four", r"five", r"six", r"seven", r"eight", r"nine"];
    let mut sum = 0;
    for line in _input.lines() {
        let mut min = 100000;
        let mut first = "";
        let mut max = 0;
        let mut last = "";
        for reg in find.iter() {
            let re = Regex::new(reg).unwrap();
            for cap in re.captures_iter(line) {
                let digit = cap.get(0).unwrap();
                if digit.start() < min {
                    min = digit.start();
                    first = _to_string_digit(digit.as_str());
                }
                if digit.start() >= max {
                    max = digit.start();
                    last = _to_string_digit(digit.as_str());
                }
            }

        }
        let digit = format!("{}{}", first, last).parse::<usize>().unwrap();
        sum += digit;
    }
    return sum;
}
