use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::env;
use std::collections::VecDeque;

const MAGIC: usize = 9;

fn main() {
    let args: Vec<String> = env::args().collect();

    let _input = &args[1];
    part_1(_input);
    part_2(_input);
}


fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn part_1(which_input: &String) {
    let mut filename: String = "../data/".to_owned();
    filename.push_str(which_input);
    let mut yard: Vec<VecDeque<char>> = std::iter::repeat(VecDeque::new()).take(MAGIC).collect::<Vec<_>>();
    if let Ok(mut lines) = read_lines(filename) {
        while let Some(line) = lines.next() {
            let _crate: String = line.unwrap();
            if _crate.len() == 0 { break; }
            let row_chars: Vec<char> = _crate.chars().collect();
            if row_chars[1] == '1' { break; }
            for i in 0..MAGIC {
                let c = row_chars[4 * i + 1];
                if c == ' ' { continue; }
                yard[i].push_front(c);
            }
        }
        while let Some(line) = lines.next() {
            let instruction: String = line.unwrap();
            if instruction.len() == 0 { continue; }
            let split_int: Vec<&str> = instruction.split_whitespace().collect::<Vec<&str>>();
            let iterations: usize = split_int[1].parse().unwrap();
            let source: usize = split_int[3].parse().unwrap();
            let destination: usize = split_int[5].parse().unwrap();
            let mut idx = 0;
            while idx < iterations {
                let _crate = yard[source -  1].pop_back().unwrap();
                yard[destination - 1].push_back(_crate);
                idx += 1;
            }
        }
    }
    let mut top = String::from("");
    for i in 0..MAGIC {
        top.push(yard[i].pop_back().unwrap());
    }
    println!("Top crates for part 1: {top}");
}


fn part_2(which_input: &String) {
    let mut filename: String = "../data/".to_owned();
    filename.push_str(which_input);
    let mut yard: Vec<VecDeque<char>> = std::iter::repeat(VecDeque::new()).take(MAGIC).collect::<Vec<_>>();
    if let Ok(mut lines) = read_lines(filename) {
        while let Some(line) = lines.next() {
            let _crate: String = line.unwrap();
            if _crate.len() == 0 { break; }
            let row_chars: Vec<char> = _crate.chars().collect();
            if row_chars[1] == '1' { break; }
            for i in 0..MAGIC {
                let c = row_chars[4 * i + 1];
                if c == ' ' { continue; }
                yard[i].push_front(c);
            }
        }
        while let Some(line) = lines.next() {
            let instruction: String = line.unwrap();
            if instruction.len() == 0 { continue; }
            let split_int: Vec<&str> = instruction.split_whitespace().collect::<Vec<&str>>();
            let iterations: usize = split_int[1].parse().unwrap();
            let source: usize = split_int[3].parse().unwrap();
            let destination: usize = split_int[5].parse().unwrap();
            let mut idx = 0;
            let stack_height = yard[destination - 1].len();
            while idx < iterations {
                let _crate = yard[source -  1].pop_back().unwrap();
                yard[destination - 1].insert(stack_height, _crate);
                idx += 1;
            }
        }
    }
    let mut top = String::from("");
    for i in 0..MAGIC {
        top.push(yard[i].pop_back().unwrap());
    }
    println!("Top crates for part 2: {top}");
}
