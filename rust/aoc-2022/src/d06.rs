pub fn solve_06(_input: &String) {
    let sol_1 = solve(_input, 4);
    println!("day 06, part 1, signal starts @ {sol_1}");
    let sol_2 = solve(_input, 14);
    println!("day 06, part 2, signal starts @ {sol_2}");
}

fn solve(_input: &String, marker_len: usize) -> usize {
    for (idx, window) in _input.chars().collect::<Vec<char>>().windows(marker_len).enumerate() {
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
