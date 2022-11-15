use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::env;
use std::collections::HashSet;
use std::iter::FromIterator;


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
    let mut overlaps: i32 = 0;
    if let Ok(lines) = read_lines(filename) {
        for line in lines {
            if let Ok(assignments) = line {
                let mut pair = assignments.split(",");
                let mut first = pair.next().expect("string").split("-");
                let mut second = pair.next().expect("string").split("-");
                let first_min: i32 = first.next().expect("string").parse().unwrap();
                let first_max: i32 = first.next().expect("string").parse().unwrap();
                let second_min: i32 = second.next().expect("string").parse().unwrap();
                let second_max: i32 = second.next().expect("string").parse().unwrap();
                if first_min <= second_min && first_max >= second_max {
                    overlaps += 1;
                    continue;
                }
                if second_min <= first_min && second_max >= first_max {
                    overlaps += 1;
                    continue;
                }
            }
        }
    }
    println!("Part 1 contains {overlaps} overlaps");
}


fn part_2(which_input: &String) {
    let mut filename: String = "../data/".to_owned();
    filename.push_str(which_input);
    let mut overlaps: i32 = 0;
    if let Ok(lines) = read_lines(filename) {
        for line in lines {
            if let Ok(assignments) = line {
                // so much boiler just to parse the string
                // is there a better way?
                let mut pair = assignments.split(",");
                let mut first = pair.next().expect("string").split("-");
                let mut second = pair.next().expect("string").split("-");
                let first_min: i32 = first.next().expect("string").parse().unwrap();
                let first_max: i32 = first.next().expect("string").parse().unwrap();
                let second_min: i32 = second.next().expect("string").parse().unwrap();
                let second_max: i32 = second.next().expect("string").parse().unwrap();
                let first_set: HashSet<i32> = HashSet::from_iter((first_min..first_max + 1).collect::<Vec<i32>>());
                let second_set: HashSet<i32> = HashSet::from_iter((second_min..second_max + 1).collect::<Vec<i32>>());
                if first_set.intersection(&second_set).count() > 0 {
                    overlaps += 1;
                }
            }
        }
    }
    println!("Part 2 contains {overlaps} partial overlaps");
}
