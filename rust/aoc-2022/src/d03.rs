use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::env;
use std::collections::HashSet;
use std::iter::FromIterator;


fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}


fn main() {
    let args: Vec<String> = env::args().collect();

    let _input = &args[1];
    part_1(_input);
    part_2(_input);
}


fn priority_as_num(priority: char) -> u32 {
    if priority.is_lowercase() {
        return (priority as u32) - 96;
    }
    return (priority as u32) - 38;
}


fn sift_rucksack(rucksack: String) -> Option<char> {
    // something like set intersection may be more performant?
    // does rust have something like python sets? HashSet? See Part 2
    let len: usize = rucksack.chars().count();
    for (i, c) in rucksack.chars().enumerate() {
        if i >= len / 2 { break; }
        for rev_i in len / 2 .. len {
            if c == rucksack.chars().nth(rev_i).unwrap() {
                return Some(c);
            }
        }
    }
    return None;
}


fn part_1(which_input: &String) {
    let mut filename: String = "../data/".to_owned();
    filename.push_str(which_input);
    let mut sum_of_priorities: u32 = 0;
    if let Ok(lines) = read_lines(filename) {
        for line in lines {
            if let Ok(rucksack) = line {
                let priority = sift_rucksack(rucksack);
                match priority {
                    Some(x) => {
                        let num_rep = priority_as_num(x);
                        //println!("Priority {} {:?}", x, num_rep);
                        sum_of_priorities += num_rep;
                    }
                    None => panic!("Couldn't find a priority"),
                }
            }
        }
    }
    println!("sum of priorities: {sum_of_priorities}");
}


fn part_2(which_input: &String) {
    let mut filename: String = "../data/".to_owned();
    filename.push_str(which_input);
    let mut sum_of_priorities: u32 = 0;
    if let Ok(mut lines) = read_lines(filename) {
        while let (Some(first), Some(second), Some(third)) = (lines.next(), lines.next(), lines.next()) {
            let first_set: HashSet<char> = HashSet::from_iter(first.expect("Expected String").chars().collect::<Vec<char>>());
            let second_set: HashSet<char> = HashSet::from_iter(second.expect("Expected String").chars().collect::<Vec<char>>());
            let third_set: HashSet<char> = HashSet::from_iter(third.expect("Expected String").chars().collect::<Vec<char>>());

            let first_second_common = first_set.intersection(&second_set);
            let first_second_set: HashSet<char> = HashSet::from_iter(first_second_common.copied().collect::<Vec<char>>());
            let mut badge = first_second_set.intersection(&third_set);
            match badge.nth(0) {
                Some(x) => {
                    sum_of_priorities += priority_as_num(*x);
                }
                None => panic!("Couldn't find a badge"),
            }
        }
    }
    println!("sum of priorities: {sum_of_priorities}");
}
