use std::fs;
use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();

    let _input = &args[1];
    let sol_1 = solve(_input, 4);
    println!("part 1, signal starts @ {sol_1}");
    let sol_2 = solve(_input, 14);
    println!("part 2, signal starts @ {sol_2}");
}


fn solve(which_input: &String, marker_len: usize) -> usize {
    let mut filename: String = "../data/".to_owned();
    filename.push_str(which_input);
    let signal: String = fs::read_to_string(filename).unwrap();
    for (idx, window) in signal.chars().collect::<Vec<char>>().windows(marker_len).enumerate() {
        let mut dupe = false;
        for i in 0..marker_len {
            for j in 0..marker_len {
                if i == j { continue; }
                if window[i] == window[j] { dupe = true; break;}
            }
            if dupe { break; }
        }
        if !dupe {
            return idx + marker_len;
        }
    }
    return 0;
}
